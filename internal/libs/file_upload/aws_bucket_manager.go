package fileupload

import (
	"os"

	di_file_manager "khrix/egommerce/internal/libs/file_manager/di"

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
	fileManager di_file_manager.FileManager
	s3          *s3.S3
}

func NewAwsBuckerManager(
	fileManager di_file_manager.FileManager, s3Data S3Data,
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

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	return s3.New(sess, aws.NewConfig().WithRegion(m.s3Data.BuckerRegion))
}

func (m AwsBuckerManager) UploadFile(content []byte, filename string) (*string, error) {
	tempPath, err := m.fileManager.CreateFile(content, filename)
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

	putResponse, err := m.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(m.s3Data.Bucket),
		Key:    aws.String(m.s3Data.Key),
		Body:   tempFile,
	})
	if err != nil {
		return nil, err
	}

	imageUrl := putResponse.GoString()

	return &imageUrl, nil
}
