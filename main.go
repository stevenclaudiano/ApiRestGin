package main

import (
	"go-api/models"
	"go-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Steven", CPF: "484.087.698-49", Cidade: "Itapui"},
		{Nome: "Bia", CPF: "484.434.5656-69", Cidade: "Bauru"},
	}
	routes.HandleRequestes()
}
