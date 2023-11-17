package fileupload

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pathresolver "khrix/egommerce/internal/libs/path_resolver"

	"github.com/google/uuid"
)

type FileUploadManager struct{}

func NewFileUploadManager() *FileUploadManager {
	return &FileUploadManager{}
}

func (manager FileUploadManager) verifyOrCreateDir() (*string, error) {
	folderPath := pathresolver.GetCurrentPath("asset")

	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		return &folderPath, nil
	}

	err := os.Mkdir(folderPath, 0o755)
	if err != nil {
		return nil, err
	}

	return &folderPath, nil
}

func (manager FileUploadManager) UploadFile(content []byte, filename string) (*string, error) {
	fileExt := strings.Split(filename, ".")[1]

	folderPath, err := manager.verifyOrCreateDir()
	if err != nil {
		return nil, err
	}

	newFileName := fmt.Sprintf("%s.%s", uuid.NewString(), fileExt)
	newFilePath := filepath.Join(*folderPath, newFileName)

	file, err := os.Create(newFilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	defer bufferedWriter.Flush()

	_, err = bufferedWriter.Write(content)

	if err != nil {
		return nil, err
	}

	err = bufferedWriter.Flush()
	if err != nil {
		return nil, err
	}

	return &newFilePath, nil
}
