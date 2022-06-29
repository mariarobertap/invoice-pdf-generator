package utils

import (
	"io"
	"log"
	"os"
)

func GetFile() (*os.File, error) {
	file, err := os.OpenFile("./file/invoice_2022.html", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return file, err
	}
	return file, nil
}

func CreateHtmlFile() {
	fin, err := os.Open("./template/header.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create("./file/invoice_2022.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}
