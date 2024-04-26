package model

type PersMail struct {
	Pers_em_id uint64 `gorm:"primaryKey"`
	Email      string
}
