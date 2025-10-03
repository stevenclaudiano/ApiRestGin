package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequestes() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
