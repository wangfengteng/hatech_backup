package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/hatech/backup/pkg/client"
	"k8s.io/klog/v2"
	"time"
)

type ImageBackupService struct {
}

var defaultImageBackupService *ImageBackupService

func GetImageBackupService() *ImageBackupService {
	if defaultImageBackupService == nil {
		defaultImageBackupService = newImageBackupService()
	}
	return defaultImageBackupService
}

func newImageBackupService() *ImageBackupService {
	return &ImageBackupService{}
}

func (s *ImageBackupService) Save() {
	var authConfig = types.AuthConfig{
		Username:      "hatech1",
		Password:      "p@ssw0rD!",
		ServerAddress: "https://index.docker.io/v1/",
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	tag := "hatech1" + "/hello"
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	rd, err := client.GetDockerClient().ImagePush(ctx, tag, opts)
	if err != nil {
		klog.Errorf("push docker image : %s ,error :%s", tag, err.Error())
	}
	defer rd.Close()
	klog.Infof("%v", rd)
}
