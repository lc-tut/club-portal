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
				personalGroup.GET("/favs/:clubuuid", r.middleware.SetClubUUIDKey(), h.GetIsFavoriteClub())
				personalGroup.POST("/favs", h.CreateFavoriteClub())
				personalGroup.POST("/unfav", h.UnFavoriteClub())
			}
		}
		clubGroup := v1Group.Group("/clubs")
		{
			clubGroup.GET("/", h.GetAllClub())
			clubGroup.POST("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.CreateClub())
			clubGroup.PUT("/", r.middleware.CheckSession(), r.middleware.GeneralOnly(), h.UpdateClub())
			clubGroup.GET("/search", h.GetSearch())
			clubGroup.GET("/slug/:clubslug", r.middleware.SetClubSlugKey(), r.middleware.SetIsRestrictedSession(), h.GetClubFromSlug())
			personalClubGroup := clubGroup.Group("/uuid/:clubuuid", r.middleware.SetClubUUIDKey())
			{
				personalClubGroup.GET("/", r.middleware.SetIsRestrictedSession(), h.GetClubFromUUID())
				personalClubGroup.DELETE("/", r.middleware.CheckSession(), r.middleware.AdminOnly(), h.DeleteClub())
				personalClubGroup.GET("/achievement", h.GetClubAchievement())
				personalClubGroup.PUT("/achievement", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubAchievement())
				personalClubGroup.GET("/activity_detail", h.GetClubActivityDetails())
				personalClubGroup.PUT("/activity_detail", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubActivityDetails())
				personalClubGroup.GET("/content", h.GetClubContent())
				personalClubGroup.PUT("/content", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubContent())
				personalClubGroup.GET("/description", h.GetClubDescription())
				personalClubGroup.PUT("/description", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubDescription())
				personalClubGroup.GET("/image", h.GetClubImages())
				personalClubGroup.PUT("/image", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubImages())
				personalClubGroup.GET("/link", r.middleware.CheckSession(), h.GetClubLinks())
				personalClubGroup.PUT("/link", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubLinks())
				personalClubGroup.GET("/schedule", r.middleware.CheckSession(), h.GetClubSchedule())
				personalClubGroup.PUT("/schedule", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubSchedule())
				//personalClubGroup.GET("/tpremark")
				//personalClubGroup.PUT("/tpremark", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName))
				personalClubGroup.GET("/video", h.GetClubVideo())
				personalClubGroup.PUT("/video", r.middleware.CheckSession(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubVideo())
			}
		}
		uploadGroup := v1Group.Group("/upload")
		{
			imageGroup := uploadGroup.Group("/images")
			{
				imageGroup.GET("/", r.middleware.CheckSession(), h.GetImages())
				imageGroup.POST("/", r.middleware.CheckSession(), h.UploadImage())
				imageGroup.GET("/:imageid", r.middleware.SetImageIDKey(), h.GetSpecificImage())
				imageGroup.DELETE("/:imageid", r.middleware.CheckSession(), r.middleware.SetImageIDKey(), h.DeleteImage()) // 他ユーザが消せる可能性があるかも
			}
			thumbnailGroup := uploadGroup.Group("/thumbnail")
			{
				thumbnailClubGroup := thumbnailGroup.Group("/clubs")
				{
					thumbnailClubGroup.PUT("/:clubuuid", r.middleware.CheckSession(), r.middleware.SetClubUUIDKey(), r.middleware.GeneralOnly(), r.middleware.IdentifyUUID(consts.ClubUUIDKeyName), h.UpdateClubThumbnail())
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
			adminUserGroup := adminGroup.Group("/users")
			{
				adminUserGroup.GET("/", adminH.GetAllUserFromAdmin())
				adminUserGroup.GET("/:useruuid", r.middleware.SetUserUUIDKey(), adminH.GetUserFromAdmin())
				adminUserGroup.PUT("/:useruuid", r.middleware.SetUserUUIDKey(), adminH.UpdateUserFromAdmin())
			}
			adminClubGroup := adminGroup.Group("/clubs")
			{
				adminClubGroup.PUT("/:clubuuid", r.middleware.SetClubUUIDKey(), adminH.UpdateClubWithVisible())
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
