package minio

import (
	"github.com/amorindev/go-tmpl/pkg/features/app/leader/port"
	"github.com/minio/minio-go/v7"
)

var _ port.LeaderFileStg = &FileStorage{}

type FileStorage struct {
	MinioClient *minio.Client
	BucketName  string
	BaseUrl     string
}

func NewLeaderFileStg(client *minio.Client, bucketName string, baseUrl string) *FileStorage {
	return &FileStorage{
		MinioClient: client,
		BucketName:  bucketName,
		BaseUrl:     baseUrl,
	}
}
