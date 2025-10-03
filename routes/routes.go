package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequestes() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)     // Exibe todos os alunos
	r.GET("/:nome", controllers.Saudacao)         //
	r.POST("/alunos", controllers.CriarNovoAluno) //Criar novos alunos
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.Run()
}

//	r.HandleFunc("/api/personalidades/{id}", controller.DeletaUmaPersonalidade).Methods("Delete")
