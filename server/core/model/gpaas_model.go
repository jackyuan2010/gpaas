package model

type GPaasModel struct {
	ID          string `gorm:"primaryKey"`
	TenantId    string
	DataStatus  int
	DataVersion int
	CreatedAt   int64 `gorm:"autoCreateTime:nano"`
	CreatedBy   string
	UpdatedAt   int64 `gorm:"autoUpdateTime:nano"`
	UpdatedBy   string
}
