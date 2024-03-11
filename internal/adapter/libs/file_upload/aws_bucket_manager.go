package fileupload

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"khrix/egommerce/internal/port/libs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Data struct {
	Bucket       string
	Key          string
	BuckerRegion string
}

type AwsBuckerManager struct {
	s3Data      S3Data
	fileManager libs.FileManager
	s3          *s3.S3
}

func NewAwsBuckerManager(
	fileManager libs.FileManager, s3Data S3Data,
) *AwsBuckerManager {
	instance := AwsBuckerManager{
		fileManager: fileManager,
		s3Data:      s3Data,
	}

	instance.s3 = instance.createSession()

	return &instance
}

func (m AwsBuckerManager) createSession() *s3.S3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return s3.New(sess, aws.NewConfig().WithRegion(*aws.String(m.s3Data.BuckerRegion)))
}

func (m AwsBuckerManager) generateFileLink(fileName string) *string {
	url := "https://%s.s3.%s.amazonaws.com/%s"
	url = fmt.Sprintf(url, m.s3Data.Bucket, m.s3Data.BuckerRegion, fileName)

	return &url
}

func (m AwsBuckerManager) UploadFile(content []byte, filename string) (*string, error) {
	tempPath, err := m.fileManager.CreateFile(content, filename)
	cType := http.DetectContentType(content)

	if err != nil {
		return nil, err
	}

	tempFile, err := os.Open(*tempPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		tempFile.Close()
		m.fileManager.DeleteFile(*tempPath)
	}()

	fileNameKey := fmt.Sprintf("%s-%s", m.s3Data.Key, filepath.Base(tempFile.Name()))
	publicTag := "public=yes"

	putResponse, err := m.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(m.s3Data.Bucket),
		Key:         &fileNameKey,
		Body:        tempFile,
		Tagging:     &publicTag,
		ContentType: &cType,
	})
	if err != nil {
		return nil, err
	}

	uploadUrl := m.generateFileLink(fileNameKey)

	fmt.Printf("File Upload - %s - %s", *putResponse.ETag, *uploadUrl)

	return uploadUrl, nil
}
