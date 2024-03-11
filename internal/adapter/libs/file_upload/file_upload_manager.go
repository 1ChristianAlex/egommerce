package fileupload

import "khrix/egommerce/internal/port/libs"

type FileUploadManager struct {
	fileManager libs.FileManager
}

func NewFileUploadManager(fileManager libs.FileManager) *FileUploadManager {
	return &FileUploadManager{fileManager: fileManager}
}

func (fm FileUploadManager) UploadFile(content []byte, filename string) (*string, error) {
	return fm.fileManager.CreateFile(content, filename)
}
