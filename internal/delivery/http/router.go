package http

import (
	"coins/configs"
	"coins/configs/types/logger"
	"coins/internal/delivery/http/middlewares"
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter(config configs.Config) {
	if config.App.Environment.IsProduction() {
		switch config.Log.Level {
		case logger.Debug:
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.Use(middlewares.Auth)

	d.coins(router.Group("/coins"))

	d.router = router
}

func (d *Delivery) coins(router *gin.RouterGroup) {
	router.POST("/", d.Handlers.Coin.Create)
	router.GET("/", d.Handlers.Coin.List)

	d.coin(router.Group("/:id"))
}

func (d *Delivery) coin(router *gin.RouterGroup) {
	router.PUT("/", d.Handlers.Coin.Update)
	router.DELETE("/", d.Handlers.Coin.Delete)

	d.urls(router.Group("/urls"))
}

func (d *Delivery) urls(router *gin.RouterGroup) {
	router.GET("/", d.Handlers.Url.List)
	router.POST("/", d.Handlers.Url.Create)

	d.url(router.Group("/:urlId"))
}

func (d *Delivery) url(router *gin.RouterGroup) {
	router.PUT("/", d.Handlers.Url.Update)
	router.DELETE("/", d.Handlers.Url.Delete)
}
