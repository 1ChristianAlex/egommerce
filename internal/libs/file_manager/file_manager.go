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

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (f FileManager) VerifyOrCreateDir(subFolder string) (*string, error) {
	folderPath := pathresolver.GetCurrentPath(subFolder)

	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		return &folderPath, nil
	}

	err := os.Mkdir(folderPath, 0o755)
	if err != nil {
		return nil, err
	}

	return &folderPath, nil
}

func (f FileManager) CreateFile(content []byte, filename string) (*string, error) {
	fileExt := strings.Split(filename, ".")[1]

	pathParts := []string{"assets", "temp"}
	var completPath string

	for index := range pathParts {
		customPath := pathParts[0:index]
		completePathLoop, err := f.VerifyOrCreateDir(strings.Join(customPath, "/"))
		if err != nil {
			return nil, err
		}

		completPath = *completePathLoop
	}

	newFileName := fmt.Sprintf("%s.%s", uuid.NewString(), fileExt)
	newFilePath := filepath.Join(completPath, newFileName)

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

func (f FileManager) DeleteFile(filepath string) {
	os.Remove(filepath)
}
