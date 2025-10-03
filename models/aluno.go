package models

type Aluno struct {
	Nome   string `json:"Nome"`
	CPF    string `json:"CPF"`
	Cidade string `json:"Cidade"`
}

var Alunos []Aluno
