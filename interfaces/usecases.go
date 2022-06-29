package interfaces

type FileService interface {
	LoadHtmlFile(htmlFileDir string, dataToLoad any) error
	CreateHtmlFile()
}
