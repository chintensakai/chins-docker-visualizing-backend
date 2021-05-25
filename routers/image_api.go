package routers

import (
	"net/http"

	"github.com/chins/go-docker-visualizing/pkg"
	"github.com/chins/go-docker-visualizing/sdk"
	"github.com/gin-gonic/gin"
)

// 镜像列表
func GetImages(c *gin.Context) {
	images := sdk.GetImages()
	c.JSON(http.StatusOK, gin.H{
		"code": pkg.SUCCESS,
		"msg":  pkg.GetMsg(pkg.SUCCESS),
		"data": images,
	})
}
