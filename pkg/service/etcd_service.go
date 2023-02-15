package service

type EtcdBackupService struct {
}

var defaultEtcdBackupService *EtcdBackupService

func GetEtcdBackupService() *EtcdBackupService {
	if defaultEtcdBackupService == nil {
		defaultEtcdBackupService = newEtcdBackupService()
	}
	return defaultEtcdBackupService
}

func newEtcdBackupService() *EtcdBackupService {
	return &EtcdBackupService{}
}
