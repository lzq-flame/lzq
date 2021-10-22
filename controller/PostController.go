/**
2 * @Author: lzq
3 * @Date: 2021/10/22 9:51
4 */

package controller

import (
	"github.com/gin-gonic/gin"
	"lzq/responese"
	"lzq/utils"
)

func InterceptVideo(ctx *gin.Context) {
	url := ctx.PostForm("url")
	//startTime := ctx.PostForm("startTime")
	//endTIme := ctx.PostForm("endTime")
	fromUrl, err := utils.DownloadFromUrl(url)
	if err != nil {
		responese.Fail(ctx, nil, "服务器处理错误")
		return
	}
	responese.Success(ctx, gin.H{"path": fromUrl}, "ok")
}
