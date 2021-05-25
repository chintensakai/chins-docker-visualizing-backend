package routers

import (
	"github.com/chins/go-docker-visualizing/middleware"
	"github.com/chins/go-docker-visualizing/pkg"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(pkg.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/")
	{
		//获取容器列表
		api.GET("/containers", GetContainers)
		//获取镜像列表
		api.GET("/images", GetImages)
	}

	return r
}
