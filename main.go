package main

import (
	"github.com/pasteldiventu/go-rest-api/models"
	"github.com/pasteldiventu/go-rest-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Guilherme", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana Clara", CPF: "11111111111", RG: "4800000000"},
	}
	routes.HandleRequests()
}
