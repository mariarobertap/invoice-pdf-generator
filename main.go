package main

import (
	"log"

	"github.com/mariarobertap/invoice-pdf-generator/models"
	"github.com/mariarobertap/invoice-pdf-generator/service"
	"github.com/mariarobertap/invoice-pdf-generator/utils"
)

var (
	invoice = models.InvoiceInfo{
		StatusPedido: "checked",
		Due:          "22 de abril, 2022",
		CreatedAt:    "22 de abril, 2001",
		NotaID:       "1232334",
		ClienteName:  "Maria Roberta",
		Endereco:     "Imaginario",
		Bairro:       "Imaginario",
		Cidade:       "Imaginario",
		Telefone:     "(55)44 901829038",
	}
	payment = models.Payment{PaymentMethod: "Cartão de Débito"}

	result = []models.Item{
		{
			Name:       "Espelho",
			ValueTotal: 300,
			ValueUnity: "300",
		},
		{
			Name:       "Vidro",
			ValueTotal: 400,
			ValueUnity: "300",
		},
	}

	final = models.ItemsResult{
		Items:      result,
		ItemsTotal: 500,
	}
)

func main() {

	utils.CreateHtmlFile()
	file, err := utils.GetFile()
	defer file.Close()

	if err != nil {
		log.Panic("Erro", err.Error())
	}

	fileService := service.NewFileService(file)

	fileService.LoadHtmlFile("./template/clientBody.html", invoice)
	fileService.LoadHtmlFile("./template/paymentBody.html", payment)
	fileService.LoadHtmlFile("./template/itemsBody.html", final)

}
