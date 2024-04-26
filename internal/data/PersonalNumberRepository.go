package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type PersonalNumberRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewPersonalNumberRepository(d *Data, logger log.Logger) *PersonalNumberRepository {
	return &PersonalNumberRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBPersPhone(d *Data) error {

	err := d.db.AutoMigrate(&model.PersPhonenumber{})
	if err != nil {
		return err
	}
	return nil
}

func (r *PersonalNumberRepository) CreatePersPhone(country_code int32, number int64) error {
	persphone := model.PersPhonenumber{Country_Code: country_code, Number: number}
	err := r.db.Create(&persphone).Error
	if err != nil {
		r.log.Errorf("failed to create personal number: %v", err)
		return errors.New("failed to create personal number")
	}
	return nil
}

func (r *PersonalNumberRepository) GetByIDPersPhone(id uint64) (*model.PersPhonenumber, error) {
	persphone := model.PersPhonenumber{}
	err := r.db.First(&persphone, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("personal number not found")
		}
		r.log.Errorf("failed to get personal number by ID: %v", err)
		return nil, errors.New("failed to get personal number by ID")
	}
	return &persphone, nil
}

func (r *PersonalNumberRepository) UpdatePersPhone(id uint64, country_code int32, number int64) error {
	persphone := model.PersPhonenumber{}
	r.db.First(&persphone, id)
	persphone.Country_Code = country_code
	persphone.Number = number
	err := r.db.Save(&persphone).Error
	if err != nil {
		r.log.Errorf("failed to update personal number: %v", err)
		return errors.New("failed to update personal number")
	}
	return nil
}

func (r *PersonalNumberRepository) DeletePersPhone(id uint64) error {
	err := r.db.Delete(&model.PersPhonenumber{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete personal number: %v", err)
		return errors.New("failed to delete personal number")
	}
	return nil
}

func (r *PersonalNumberRepository) GetAllPersPhone() (*[]model.PersPhonenumber, error) {
	var persphones []model.PersPhonenumber
	err := r.db.Find(&persphones).Error
	if err != nil {
		r.log.Errorf("failed to retrieve personal numbers: %v", err)
		return &persphones, errors.New("failed to get personal number list")
	}
	return &persphones, nil
}
