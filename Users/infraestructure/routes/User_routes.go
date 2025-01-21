package routes


import (
	"github.com/gin-gonic/gin"
	"ejemplo/practica/Users/infraestructure/controllers/postUser"
)

func RoutesUser(){
	r:=gin.Default()
	r.GET("/users",func(ctx *gin.Context) {
		postuser.Post()
		ctx.JSON(200,gin.H{
			"message":"todo ok",
		})
	})
	r.Run()
}