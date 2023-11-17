package di

type FileUploadManager interface {
	UploadFile(content []byte, filename string) (*string, error)
}
