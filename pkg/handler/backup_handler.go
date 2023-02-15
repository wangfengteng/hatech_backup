package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hatech/backup/pkg/service"
	"net/http"
)

type BackupHandler struct {
	imageService *service.ImageBackupService
	etcdService  *service.EtcdBackupService
}

func NewBackupHandler() *BackupHandler {
	return &BackupHandler{
		imageService: service.GetImageBackupService(),
		etcdService:  service.GetEtcdBackupService(),
	}
}

// backup godoc
// @Summary hatech kubernetes list backups
// @Description list backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups [get]
func (h *BackupHandler) ListBackups(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "request not found"})
}

// backup godoc
// @Summary hatech kubernetes create backup
// @Description list backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups [post]
func (h *BackupHandler) CreateBackup(c *gin.Context) {

}

// backup godoc
// @Summary hatech kubernetes get backup
// @Description list backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups [get]
func (h *BackupHandler) GetBackup(c *gin.Context) {

}

// backup godoc
// @Summary hatech kubernetes backup list
// @Description list backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups/:uid [get]
//func (h *BackupHandler) PauseBackup(c *gin.Context) {
//
//}

// backup godoc
// @Summary hatech kubernetes backup list
// @Description list backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups/:uid [get]
//func (h *BackupHandler) ResumeBackup(c *gin.Context) {
//
//}

// backup godoc
// @Summary hatech kubernetes update backup
// @Description update backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.Backup
// @Failure 500 {object} model.R
// @Router /backups/:uid [put]
func (h *BackupHandler) UpdateBackup(c *gin.Context) {

}

// backup godoc
// @Summary hatech kubernetes delete backup
// @Description delete backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.R
// @Failure 500 {object} model.R
// @Router /backups [delete]
func (h *BackupHandler) DeleteBackup(c *gin.Context) {

}

// backup godoc
// @Summary hatech kubernetes multi delete backup
// @Description backup delete backups
// @Tags backup
// @Accept json
// @Produce json
// @Success 200 {object} model.R
// @Failure 500 {object} model.R
// @Router /backups/delete [post]
func (h *BackupHandler) DeleteBackups(c *gin.Context) {

}
