/**
2 * @Author: lzq
3 * @Date: 2021/10/22 10:10
4 */

package main

import (
	"github.com/gin-gonic/gin"
	"lzq/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/intercept", controller.InterceptVideo)
	return r
}
