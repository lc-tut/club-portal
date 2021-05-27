package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
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

	v1Group := r.rg.Group("/v1")
	{
		userGroup := v1Group.Group("/users", r.middleware.CheckSession())
		{
			userGroup.GET("/", h.GetUser())
			userGroup.POST("/", r.middleware.AdminOnly(), h.CreateGeneralUser())
			personalGroup := userGroup.Group("/:useruuid", r.middleware.SetUserUUIDKey())
			{
				personalGroup.GET("/", r.middleware.PersonalOrAdminOnly(), h.GetUserUUID())
				personalGroup.PUT("/", r.middleware.OverGeneralOnly(), h.UpdateUser())
				personalGroup.GET("/favs", r.middleware.PersonalOrAdminOnly(), h.GetFavoriteClubs())
				personalGroup.POST("/favs", r.middleware.PersonalOrAdminOnly(), h.CreateFavoriteClub())
				personalGroup.POST("/unfav", r.middleware.PersonalOrAdminOnly(), h.UnFavoriteClub())
			}
		}
		clubGroup := v1Group.Group("/clubs")
		{
			clubGroup.GET("/", h.GetAllClub())
			clubGroup.POST("/", r.middleware.CheckSession(), r.middleware.OverGeneralOnly(), h.CreateClub())
			clubGroup.PUT("/:clubuuid", r.middleware.CheckSession(), r.middleware.SetClubUUIDKey(), r.middleware.IdentifyClubUUID(), r.middleware.OverGeneralOnly(), h.UpdateClub())
			clubGroup.DELETE("/:clubuuid", r.middleware.CheckSession(), r.middleware.SetClubUUIDKey(), r.middleware.AdminOnly(), h.DeleteClub())
			clubGroup.GET("/:clubslug", r.middleware.SetClubSlugKey(), h.GetClub())
		}
		uploadGroup := v1Group.Group("/upload", r.middleware.CheckSession())
		{
			imageGroup := uploadGroup.Group("/images")
			{
				imageGroup.GET("/", h.GetImages())
				imageGroup.POST("/", h.UploadImage())
				imageGroup.GET("/:imageid", r.middleware.SetImageIDKey(), h.GetSpecificImage())
				imageGroup.DELETE("/:imageid", r.middleware.SetImageIDKey(), h.DeleteImage())
			}
			thumbnailGroup := uploadGroup.Group("/thumbnail")
			{
				thumbnailClubGroup := thumbnailGroup.Group("/clubs")
				{
					thumbnailClubGroup.POST("/", r.middleware.GeneralOnly(), h.UploadClubThumbnail())
					thumbnailClubGroup.GET("/:clubuuid", r.middleware.SetClubUUIDKey(), h.GetClubThumbnail())
					thumbnailClubGroup.PUT("/:clubuuid", r.middleware.SetClubUUIDKey(), r.middleware.OverGeneralOnly(), r.middleware.IdentifyClubUUID(), h.UpdateClubThumbnail())
					thumbnailClubGroup.DELETE("/:clubuuid", r.middleware.SetClubUUIDKey(), r.middleware.OverGeneralOnly(), r.middleware.IdentifyClubUUID(), h.DeleteClubThumbnail())
				}
				thumbnailIDGroup := thumbnailGroup.Group("/ids")
				{
					thumbnailIDGroup.GET("/:thumbnailid", r.middleware.SetThumbnailIDKey(), h.GetThumbnail())
				}
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
