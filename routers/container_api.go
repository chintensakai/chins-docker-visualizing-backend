package routers

import (
	"net/http"

	"github.com/chins/go-docker-visualizing/pkg"
	"github.com/chins/go-docker-visualizing/sdk"
	"github.com/gin-gonic/gin"
)

// 容器列表
func GetContainers(c *gin.Context) {
	containers := sdk.GetContainers()
	c.JSON(http.StatusOK, gin.H{
		"code": pkg.SUCCESS,
		"msg":  pkg.GetMsg(pkg.SUCCESS),
		"data": containers,
	})
}

// 启动容器
func StartContainer(c *gin.Context) {
	containerId := c.Param("id")

	sdk.StartContainer(containerId)
	c.JSON(http.StatusOK, gin.H{
		"code": pkg.SUCCESS,
		"msg":  pkg.GetMsg(pkg.SUCCESS),
	})
}

// 停止容器
func StopContainer(c *gin.Context) {
	containerId := c.Param("id")

	sdk.StopContainer(containerId)
	c.JSON(http.StatusOK, gin.H{
		"code": pkg.SUCCESS,
		"msg":  pkg.GetMsg(pkg.SUCCESS),
	})
}
