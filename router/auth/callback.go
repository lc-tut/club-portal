package auth

import (
	"encoding/base64"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/data"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
	"strings"
)

func (h *Handler) Callback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer utils.DeleteCookie(ctx, consts.AuthCSRFCookieName)

		email, err := h.checkValidState(ctx)

		if err != nil || !h.config.WhitelistUsers.IsUser(email) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		newUUID, err := uuid.NewRandom()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		sessionData := data.NewSessionData(newUUID, email)

		b, err := json.Marshal(sessionData)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		sess := sessions.Default(ctx)

		sess.Set(consts.SessionKey, b)

		if err := sess.Save(); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}

func (h *Handler) checkValidState(ctx *gin.Context) (string, error) {
	queries := ctx.Request.URL.Query()

	queryState, ok := queries["state"]

	if !ok {
		return "", errors.New("invalid query")
	}

	cookieState, err := ctx.Cookie(consts.AuthCSRFCookieName)

	if err != nil {
		return "", err
	}

	if queryState[0] != cookieState {
		return "", errors.New("invalid state")
	}

	code, ok := queries["code"]

	if !ok {
		return "", errors.New("invalid query")
	}

	token, err := h.config.GoogleOAuthConfig.Exchange(ctx, code[0])

	if err != nil {
		return "", err
	}

	idToken := token.Extra("id_token").(string)

	email, err := parseJWT(idToken)

	if err != nil {
		return "", err
	}

	return email, nil
}

type jwtData struct {
	Email string `json:"email"`
}

func parseJWT(token string) (string, error) {
	jwt := strings.Split(token, ".")
	payload := strings.TrimSuffix(jwt[1], "=")
	b, err := base64.RawURLEncoding.DecodeString(payload)

	if err != nil {
		return "", err
	}

	jd := &jwtData{}

	if err := json.Unmarshal(b, jd); err != nil {
		return "", err
	}

	return jd.Email, nil
}
