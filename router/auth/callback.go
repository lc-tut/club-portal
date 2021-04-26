package auth

import (
	"encoding/base64"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func (h *Handler) Callback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			h.logger.Info("deleted cookie", zap.String("cookie_name", consts.AuthCSRFCookieName))
			utils.DeleteCookie(ctx, consts.AuthCSRFCookieName)
		}()

		data, err := h.checkValidState(ctx)

		if err != nil || !h.config.WhitelistUsers.IsUser(data.Email) {
			if err != nil {
				h.logger.Error(err.Error())
			} else {
				h.logger.Warn("invalid user", zap.String("email", data.Email))
			}
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.createSession(ctx, data); err != nil {
			h.logger.Warn("failed to create session")
			ctx.Status(http.StatusInternalServerError)
		} else {
			h.logger.Info("deleted cookie", zap.String("cookie_name", consts.AuthCSRFCookieName))
			utils.DeleteCookie(ctx, consts.AuthCSRFCookieName) // defer だと redirect 時に Cookie が削除されない
			ctx.Redirect(http.StatusFound, "/")
		}
	}
}

type jwtData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *Handler) checkValidState(ctx *gin.Context) (*jwtData, error) {
	queries := ctx.Request.URL.Query()

	queryState, ok := queries["state"]

	if !ok {
		return nil, errors.New("invalid query")
	}

	cookieState, err := ctx.Cookie(consts.AuthCSRFCookieName)

	if err != nil {
		return nil, err
	}

	if queryState[0] != cookieState {
		return nil, errors.New("invalid state")
	}

	code, ok := queries["code"]

	if !ok {
		return nil, errors.New("invalid query")
	}

	token, err := h.config.GoogleOAuthConfig.Exchange(ctx, code[0])

	if err != nil {
		return nil, err
	}

	idToken := token.Extra("id_token").(string)

	h.logger.Debug("get id_token from Google OAuth", zap.String("id_token", idToken))

	data, err := parseJWT(idToken)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseJWT(token string) (*jwtData, error) {
	jwt := strings.Split(token, ".")
	payload := strings.TrimSuffix(jwt[1], "=")
	b, err := base64.RawURLEncoding.DecodeString(payload)

	if err != nil {
		return nil, err
	}

	jd := &jwtData{}

	if err := json.Unmarshal(b, jd); err != nil {
		return nil, err
	}

	return jd, nil
}

func (h *Handler) createSession(ctx *gin.Context, data *jwtData) error {
	user, err := h.getUserOrCreate(data)

	if err != nil {
		return err
	}

	sessionUUID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	sessionData := utils.NewSessionData(sessionUUID.String(), user.GetUserID(), user.GetEmail(), user.GetName(), user.GetRole())

	b, err := json.Marshal(sessionData)

	if err != nil {
		return err
	}

	sess := sessions.Default(ctx)

	sess.Set(consts.SessionKey, b)

	if err := sess.Save(); err != nil {
		return err
	}

	h.logger.Info("created session data",
		zap.String("session_uuid", sessionUUID.String()),
		zap.String("user_uuid", user.GetUserID()),
		zap.String("email", user.GetEmail()),
		zap.String("name", user.GetName()),
		zap.String("role", user.GetRole()),
	)

	return nil
}

func (h *Handler) getUserOrCreate(data *jwtData) (users.UserInfo, error) {
	var user users.UserInfo
	var err error

	email := data.Email

	if h.config.WhitelistUsers.IsAdminUser(email) {
		h.logger.Info("logged in as admin user", zap.String("email", email))
		user, err = h.repo.GetAdminUserByEmail(email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("create admin user because of no data in db")
			newUserUUID, _err := uuid.NewRandom()

			if _err != nil {
				return nil, _err
			}

			user, err = h.repo.CreateAdminUser(newUserUUID.String(), email, data.Name)
		}
	} else if h.config.WhitelistUsers.IsGeneralUser(email) {
		h.logger.Info("logged in as general user", zap.String("email", email))
		user, err = h.repo.GetGeneralUserByEmail(email)
	} else {
		h.logger.Info("logged in as domain user", zap.String("email", email))
		user, err = h.repo.GetDomainUserByEmail(email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("create domain user because of no data in db")
			newUserUUID, _err := uuid.NewRandom() // err のオーバライドを回避するために _err とする

			if _err != nil {
				return nil, _err
			}

			user, err = h.repo.CreateDomainUser(newUserUUID.String(), email, data.Name)
		}
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
