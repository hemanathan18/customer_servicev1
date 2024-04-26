package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type OfficialNumberRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewOfficialNumberRepository(d *Data, logger log.Logger) *OfficialNumberRepository {
	return &OfficialNumberRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBOffPhone(d *Data) error {

	err := d.db.AutoMigrate(&model.OffPhonenumber{})
	if err != nil {
		return err
	}
	return nil
}

func (r *OfficialNumberRepository) CreateOffPhone(country_code int64, number int64) error {
	offphone := model.OffPhonenumber{Country_Code: country_code, Number: number}
	err := r.db.Create(&offphone).Error
	if err != nil {
		r.log.Errorf("failed to create official number: %v", err)
		return errors.New("failed to create official number")
	}
	return nil
}

func (r *OfficialNumberRepository) GetByIDOffPhone(id uint64) (*model.OffPhonenumber, error) {
	var offphone model.OffPhonenumber
	err := r.db.First(&offphone, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("official number not found")
		}
		r.log.Errorf("failed to get official number by ID: %v", err)
		return nil, errors.New("failed to get official number by ID")
	}
	return &offphone, nil
}

func (r *OfficialNumberRepository) UpdateOffPhone(id uint64, country_code int64, number int64) error {
	offphone := model.OffPhonenumber{}
	r.db.First(&offphone, id)
	offphone.Country_Code = country_code
	offphone.Number = number
	err := r.db.Save(&offphone).Error
	if err != nil {
		r.log.Errorf("failed to update official number: %v", err)
		return errors.New("failed to update official number")
	}
	return nil
}

func (r *OfficialNumberRepository) DeleteOffPhone(id uint64) error {
	err := r.db.Delete(&model.OffPhonenumber{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete official number: %v", err)
		return errors.New("failed to delete official number")
	}
	return nil
}

func (r *OfficialNumberRepository) GetAllOffPhone() (*[]model.OffPhonenumber, error) {
	var offphones []model.OffPhonenumber
	err := r.db.Find(&offphones).Error
	if err != nil {
		r.log.Errorf("failed to retrieve official numbers: %v", err)
		return &offphones, errors.New("failed to get official number list")
	}
	return &offphones, nil
}
