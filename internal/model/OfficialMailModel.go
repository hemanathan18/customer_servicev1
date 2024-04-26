package model

type OffMail struct {
	Off_em_id uint64 `gorm:"primaryKey"`
	Email     string
}
