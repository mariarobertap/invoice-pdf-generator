package interfaces

type FileService interface {
	LoadHtmlFile(htmlFileDir string, dataToLoad any) error
	CreateHtmlFile()
}
type PDFGenetorInterface interface {
	Create(htmlFile string) (string, error)
}
