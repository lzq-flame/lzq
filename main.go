package main

import "github.com/gin-gonic/gin"

var inputPath = "D:/video/lzqys.mov"
var outputPath = "D:/video/lzqys.mov"

func main() {
	r := gin.Default()
	r = CollectRoute(r)
	r.Run()
}
