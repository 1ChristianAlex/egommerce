package libs

type FileUploadManager interface {
	UploadFile(content []byte, filename string) (*string, error)
}
