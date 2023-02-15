package model

type Backup struct {
	//Uid    string       `json:"uid"`
	Spec   BackupSpec   `json:"spec"`
	Status BackupStatus `json:"status"`
}

type BackupSpec struct {
	BackupType  BackupType         `json:"backupType"`
	Cron        string             `json:"cron"`
	BackupImage BachupImagesDetail `json:"backupImage"`
	BackupEtcd  BachupEtcdDetail   `json:"backupEtcd"`
}

type BachupImagesDetail struct {
	//Registry
	Src []string `json:"src"`
	// hub.hatech.local
	// hub.hatech.local/project/
	// hub.hatech.local/project/aaa
	Dst []string `json:"dst"`
}

type BachupEtcdDetail struct {
	Storage string `json:"storage"`
}

type BackupType string

const (
	BackupType_Image BackupType = "image"
	BackupType_Etcd  BackupType = "etcd"
)

type BackupStatus struct {
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Phase      string `json:"phase"`
	Reason     string `json:"reason"`
}
