package service

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/mariarobertap/invoice-pdf-generator/interfaces"
)

type fileService struct {
	File *os.File
}

func NewFileService(fileDst *os.File) interfaces.FileService {
	return &fileService{
		File: fileDst,
	}
}

func (f *fileService) LoadHtmlFile(htmlFileDir string, dataToLoad any) error {

	content, err := ioutil.ReadFile(htmlFileDir)

	var htmlLoader = template.Must(template.New("htmlLoader").Parse(string(content)))

	if err != nil {
		log.Fatal(err)
	}
	if err := htmlLoader.Execute(f.File, dataToLoad); err != nil {
		log.Fatal(err)
	}

	return err
}

func (f *fileService) CreateHtmlFile() {
	fin, err := os.Open("./template/header.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create("invoice_2022.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}
