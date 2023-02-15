package service

import (
	"context"
	"fmt"
	client2 "github.com/hatech/backup/pkg/client"
	"github.com/hatech/backup/pkg/model"
	"github.com/minio/minio-go/v7"
	"k8s.io/klog/v2"
	"path/filepath"
)

var s3Retries uint = model.DefaultS3Retries

type MinioBackupService struct {
}

var defaultMinioBackupService *MinioBackupService

func GetMinioBackupService() *MinioBackupService {
	if defaultMinioBackupService == nil {
		defaultMinioBackupService = newMinioBackupService()
	}
	return defaultMinioBackupService
}

func newMinioBackupService() *MinioBackupService {
	return &MinioBackupService{}
}

func (s *MinioBackupService) CreateS3Backup(backupName, compressedFilePath string, bc *model.BackupConfig) error {
	// If the minio client doesn't work now, it won't after retrying
	client := client2.NewMinioClient(bc)

	compressedFile := filepath.Base(compressedFilePath)
	// If folder is specified, prefix the file with the folder
	if len(bc.Folder) != 0 {
		compressedFile = fmt.Sprintf("%s/%s", bc.Folder, compressedFile)
	}
	// check if it exists already in the bucket, and if versioning is disabled on the bucket. If an error is detected,
	// assume we aren't privy to that information and do multiple uploads anyway.
	info, _ := client.StatObject(context.TODO(), bc.BucketName, compressedFile, minio.StatObjectOptions{})
	if info.Size != 0 {
		versioning, _ := client.GetBucketVersioning(context.TODO(), bc.BucketName)
		if !versioning.Enabled() {
			klog.Info("Skipping upload to s3 because snapshot already exists and versioning is not enabled for the bucket")
			return nil
		}
	}

	err := uploadBackupFile(client, bc.BucketName, compressedFile, compressedFilePath, s3Retries)
	if err != nil {
		return err
	}
	return nil
}

func uploadBackupFile(svc *minio.Client, bucketName, fileName, filePath string, s3Retries uint) error {
	var info minio.UploadInfo
	var err error
	// Upload the zip file with FPutObject
	klog.Infof("invoking uploading backup file [%s] to s3", fileName)
	for i := uint(0); i <= s3Retries; i++ {
		info, err = svc.FPutObject(context.TODO(), bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: model.ContentType})
		if err == nil {
			klog.Infof("Successfully uploaded [%s] of size [%d]", fileName, info.Size)
			return nil
		}
		klog.Infof("failed to upload etcd snapshot file: %v, retried %d times", err, i)
	}
	return fmt.Errorf("failed to upload etcd snapshot file: %v", err)
}
