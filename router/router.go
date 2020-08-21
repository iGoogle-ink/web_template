package router

import (
	"gotil/web"
	"web_template/config"
	"web_template/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

func Init(c *config.Config, s *service.Service) {
	svc = s
	g := web.InitServer(c.Web)
	initRoute(g.Gin)
	g.Start()
}

func initRoute(g *gin.Engine) {
	g.GET("/app/ping", ping)
	app := g.Group("/app", svc.VerifyAppKey)
	{
		auth := app.Group("/auth")
		{
			auth.POST("/login")
		}
	}
}

func ping(c *gin.Context) { web.JSON(c, "OK", nil) }
