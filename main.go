package main

import (
	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Controller"
)

var cofing2 = ConfigTom.ConfigSped{}

// var regC170Dao Dao.RegC170Dao
// var produtosNfe Dao.ProdutosNfeDao

//var regC425Dao Dao.RegC425Dao

// var regC100Dao Dao.RegC100Dao

func init() {
	cofing2.Read()
	// regC100Dao.Server = cofing2.Server
	// regC100Dao.Database = cofing2.Database
	// regC100Dao.Connect()
	// regC170Dao.Server = cofing2.Server
	// regC170Dao.Database = cofing2.Database
	// regC170Dao.Connect()
	// regC170Dao.Server = cofing2.Server
	// regC170Dao.Database = cofing2.Database
	// regC170Dao.Connect()

	// produtosNfe.Server = cofing2.Server
	// produtosNfe.Database = cofing2.DataBaseNfe
	// produtosNfe.Connect()

	// regC425Dao.Server = cofing2.Server
	// regC425Dao.Database = cofing2.Database
	// regC425Dao.Connect()
}

func main() {
	// ----------- Lista todos os itens dos cupons fiscais
	// itens, err := regC425Dao.FindAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, item := range itens {
	// 	fmt.Println("" + item.DtVenda + ";" + item.CodItem + ";" + item.Qtd + ";" + item.Unid)
	// }

	// ----------- Lista todos os produtos baixados pelo scrapy da auditoria -----------
	// produtos, err := produtosNfe.FindAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, produto := range produtos {
	// 	fmt.Println("" + produto.Chave + ";" + produto.Codigo + ";" + produto.Descricao + ";" + produto.Un + ";" + produto.Cfop + ";" + produto.Qtd + ";" + produto.VlUnitComercial + ";")
	// }

	// --------------- Lista todas as chaves onde a situacao seja regular ----------------
	// regc100s, err := regC100Dao.FindByCnpj("000000000")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, regc100 := range regc100s {
	// 	if regc100.CodSit == "00" {
	// 		if regc100.ChvNfe != "" {
	// 			if regc100.IndEmit == "0" {
	// 				fmt.Println(regc100.ChvNfe)
	// 			}
	// 		}
	// 	}
	// }

	// ---------------- Lista todos produtos onde a qtd de digitos seja diferente de 10 -----------------
	// regc170s, err := regC170Dao.FindByCnpj("01990447000118")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, regc170 := range regc170s {
	// 	if len(regc170.CodItem) != 10 {
	// 		fmt.Println(regc170.CodItem)
	// 	}

	// }
	// ----------------- Importa todos arquivos da pasta que fica no config.toml ------------------------
	var importar Controller.ImportController
	importar.Importar(cofing2.PathImport)

}
