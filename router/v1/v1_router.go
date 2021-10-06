package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
	"github.com/lc-tut/club-portal/router/v1/admins"
	"go.uber.org/zap"
)

type Handler struct {
	config *config.V1Config
	logger *zap.Logger
	repo   repos.IRepository
}

func newHandler(config *config.V1Config, logger *zap.Logger, repo repos.IRepository) *Handler {
	return &Handler{
		config: config,
		logger: logger,
		repo:   repo,
	}
}

type Router struct {
	rg         *gin.RouterGroup
	config     *config.V1Config
	logger     *zap.Logger
	repo       repos.IRepository
	middleware *middlewares.Middleware
}

func (r *Router) AddRouter() {
	r.logger.Debug("initializing v1 router")
	h := newHandler(r.config, r.logger, r.repo)
	adminH := admins.NewAdminHandler(r.config, r.logger, r.repo)

	v1Group := r.rg.Group("/v1")
	{
		userGroup := v1Group.Group("/users", r.middleware.CheckSession())
		{
			userGroup.GET("/", h.GetUser())
			userGroup.POST("/", r.middleware.AdminOnly(), h.CreateGeneralUser())
			personalGroup := userGroup.Group("/:useruuid", r.middleware.SetUserUUIDKey(), r.middleware.IdentifyUUID(consts.UserUUIDKeyName))
			{
				personalGroup.GET("/", h.GetUserUUID())
				personalGroup.PUT("/", h.UpdateUser())
				personalGroup.GET("/favs", h.GetFavoriteClubs())
				personalGroup.POST("/favs", h.CreateFavoriteClub())
				personalGroup.POST("/unfav", h.UnFavoriteClub())
			}
		}
		clubGroup := v1Group.Group("/clubs")
		{
			clubGroup.GET("/", h.GetAllClub())
			clubGroup.POST("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.CreateClub())
			clubGroup.PUT("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.UpdateClub())
			clubGroup.DELETE("/:clubuuid", r.middleware.CheckSession(), r.middleware.SetClubUUIDKey(), r.middleware.AdminOnly(), h.DeleteClub())
			clubGroup.GET("/:clubslug", r.middleware.SetClubSlugKey(), h.GetClub())
		}
		uploadGroup := v1Group.Group("/upload")
		{
			imageGroup := uploadGroup.Group("/images")
			{
				imageGroup.GET("/", h.GetImages())
				imageGroup.POST("/", r.middleware.CheckSession(), h.UploadImage())
				imageGroup.GET("/:imageid", r.middleware.SetImageIDKey(), h.GetSpecificImage())
				imageGroup.DELETE("/:imageid", r.middleware.CheckSession(), r.middleware.SetImageIDKey(), h.DeleteImage())
			}
			thumbnailGroup := uploadGroup.Group("/thumbnail")
			{
				thumbnailClubGroup := thumbnailGroup.Group("/clubs")
				{
					thumbnailClubGroup.POST("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.UploadClubThumbnail())
					thumbnailClubGroup.PUT("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.UpdateClubThumbnail())
					thumbnailClubGroup.GET("/:clubuuid", r.middleware.SetClubUUIDKey(), h.GetClubThumbnail())
					thumbnailClubGroup.DELETE("/:clubuuid", r.middleware.CheckSession(), r.middleware.SetClubUUIDKey(), r.middleware.GeneralOnly(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.DeleteClubThumbnail())
				}
				thumbnailIDGroup := thumbnailGroup.Group("/id")
				{
					thumbnailIDGroup.GET("/:thumbnailid", r.middleware.SetThumbnailIDKey(), h.GetThumbnail())
				}
			}
		}
		adminGroup := v1Group.Group("/admin", r.middleware.CheckSession(), r.middleware.AdminOnly())
		{
			userGroup := adminGroup.Group("/users/:useruuid", r.middleware.SetUserUUIDKey())
			{
				userGroup.GET("/", adminH.GetUserFromAdmin())
			}
			clubGroup := adminGroup.Group("/clubs")
			{
				clubGroup.PUT("/", adminH.UpdateClubUserFromAdmin())
			}
		}
	}
}

func NewV1Router(rg *gin.RouterGroup, config *config.V1Config, logger *zap.Logger, repo repos.IRepository, middleware *middlewares.Middleware) *Router {
	r := &Router{
		rg:         rg,
		config:     config,
		logger:     logger,
		repo:       repo,
		middleware: middleware,
	}
	return r
}
