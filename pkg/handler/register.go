package handler

import "github.com/gin-gonic/gin"

func Register(rg *gin.RouterGroup) {
	base := "/backups"
	uid := "/:uid"
	backupHandler := NewBackupHandler()
	rg.POST(base, backupHandler.CreateBackup)
	rg.GET(base, backupHandler.ListBackups)
	rg.GET(base+uid, backupHandler.GetBackup)
	rg.PUT(base+uid, backupHandler.UpdateBackup)
	rg.DELETE(base+uid, backupHandler.DeleteBackup)
	rg.POST(base+"/delete", backupHandler.DeleteBackups)
}
