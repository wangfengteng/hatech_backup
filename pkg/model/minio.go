package model

import "time"

type BackupConfig struct {
	Backup     bool
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	Region     string
	EndpointCA string
	Folder     string
}

const (
	BackupBaseDir         = "/backup"
	DefaultBackupRetries  = 4
	ClusterStateExtension = "rkestate"
	CompressedExtension   = "zip"
	ContentType           = "application/zip"
	K8sBaseDir            = "/etc/kubernetes"
	DefaultS3Retries      = 3
	ServerPort            = "2379"
	S3Endpoint            = "s3.amazonaws.com"
	TmpStateFilePath      = "/tmp/cluster.rkestate"
	FailureInterval       = 15 * time.Second
)
