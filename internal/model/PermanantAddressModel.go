package model

type PermAddress struct {
	Perm_add_id uint64 `gorm:"primaryKey"`
	Door_no     int32
	Street      string
	City        string
	Zipcode     int64
	State       string
	Country     string
}
