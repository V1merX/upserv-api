package rest

import (
	"fmt"
	"net/http"

	"github.com/V1merX/upserv-api/internal/config"
	"github.com/V1merX/upserv-api/internal/service"
	v1 "github.com/V1merX/upserv-api/internal/transport/rest/v1"
	auth "github.com/V1merX/upserv-api/pkg/auth/jwt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	if cfg.Environment != config.Prod {
		// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
