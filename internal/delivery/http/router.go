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

	d.coin(router.Group("/coins"))
	d.url(router.Group("/urls"))

	d.router = router
}

func (d *Delivery) coin(router *gin.RouterGroup) {
	router.POST("/", d.Handlers.Coin.Create)
	router.PUT("/:id", d.Handlers.Coin.Update)
	router.DELETE("/:id", d.Handlers.Coin.Delete)
}

func (d *Delivery) url(router *gin.RouterGroup) {

}
