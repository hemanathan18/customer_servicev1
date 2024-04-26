package model

type OffPhonenumber struct {
	Off_ph_id    uint64 `gorm:"primaryKey"`
	Country_Code int64
	Number       int64
}
