package model

type CustomerRelation struct {
	Cr_Id       uint64 `gorm:"primaryKey"`
	Customer_ID uint64
	Off_em_id   uint64
	Pers_em_id  uint64
	Off_ph_id   uint64
	Pers_ph_id  uint64
	Temp_add_id uint64
	Perm_add_id uint64
}
