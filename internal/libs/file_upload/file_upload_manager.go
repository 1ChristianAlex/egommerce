package fileupload

import "khrix/egommerce/internal/libs/file_manager/di"

type FileUploadManager struct {
	fileManager di.FileManager
}

func NewFileUploadManager(fileManager di.FileManager) *FileUploadManager {
	return &FileUploadManager{fileManager: fileManager}
}

func (fm FileUploadManager) UploadFile(content []byte, filename string) (*string, error) {
	return fm.fileManager.CreateFile(content, filename)
}
