package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"lzq/common"
	"os"
)

var inputPath = "D:/video/lzqys.mov"
var outputPath = "D:/video/lzqys.mov"

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run()

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
