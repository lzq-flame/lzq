/**
2 * @Author: lzq
3 * @Date: 2021/10/22 9:51
4 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lzq/common"
	"lzq/model"
	"lzq/responese"
	"lzq/utils"
)

func InterceptVideo(ctx *gin.Context) {
	url := ctx.PostForm("url")
	startTime := ctx.PostForm("startTime")
	endTIme := ctx.PostForm("endTime")
	fromUrl, err := utils.DownloadFromUrl(url)
	if err != nil {
		responese.Fail(ctx, nil, "服务器处理错误")
		return
	}
	err = utils.Intercept(fromUrl, startTime, endTIme)

	//将截取视频的信息保存数据库
	var object model.Intercept
	object.GetUrl = url
	object.PostUrl = "http://127.0.0.1/after/first.mp4"
	object.StartTime = startTime
	object.EndTime = endTIme
	db := NewCreateDB()
	db.Create(&object)

	if err != nil {
		responese.Fail(ctx, nil, "服务器处理错误")
		return
	}
	responese.Success(ctx, gin.H{"url": "http://127.0.0.1/after/first.mp4"}, "ok")
}
func NewCreateDB() *gorm.DB {
	db := common.GetDB()
	db.AutoMigrate(model.Intercept{})
	return db
}
