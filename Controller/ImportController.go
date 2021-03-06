package Controller

import (
	"fmt"
	"path/filepath"

	"github.com/chapzin/GoSped/Utilidades"
)

// ImportController : Responsavel por separar e executar os comandos nos arquivos txt, csv e xml
type ImportController struct {
	arquivo string
}

// Importar : metodo responsavel pela separacao dos arquivos
func (i *ImportController) Importar(path string) {
	arquivos, _ := Utilidades.ListFiles(path)
	for _, arq := range arquivos {
		extensao := filepath.Ext(arq)
		if extensao == ".txt" || extensao == ".TXT" {
			// var spedcontroller SpedController
			// spedcontroller.addMongo(arquivo)
			// spedcontroller.validacoes(arquivo)
		}
		if extensao == ".csv" || extensao == ".CSV" {
			// TODO : faz verificacao se é um arquivo tipo siget e importa
		}
		if extensao == ".xml" || extensao == ".XML" {
			// TODO : faz verificacao se é um arquivo xml nfe, cte ou evento e importa
			fmt.Println(arq)
			ListaXmls(arq)
		}
	}
}
