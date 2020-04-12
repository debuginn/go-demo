package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/a", func(context *gin.Context) {
		var user user
		if err := context.ShouldBind(&user); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	_ = r.Run(":9090")
}
