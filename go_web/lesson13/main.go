package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
		} else {
			dst := fmt.Sprintf("./%s", file.Filename)
			_ = context.SaveUploadedFile(file, dst)

			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	_ = r.Run(":8080")
}
