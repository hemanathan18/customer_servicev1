package model

type PersPhonenumber struct {
	Pers_ph_id   uint64 `gorm:"primaryKey"`
	Country_Code int32
	Number       int64
}
