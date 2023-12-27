package di

type FileManager interface {
	VerifyOrCreateDir(subFolder string) (*string, error)
	CreateFile(content []byte, filename string) (*string, error)
	DeleteFile(filepath string)
}
