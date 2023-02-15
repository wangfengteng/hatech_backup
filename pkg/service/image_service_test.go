package service

import "testing"

func TestSave(t *testing.T) {
	GetImageBackupService().Save()
}
