package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TempAddressRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewTempAdressRepository(d *Data, logger log.Logger) *TempAddressRepository {
	return &TempAddressRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBTempAddress(d *Data) error {

	err := d.db.AutoMigrate(&model.TempAddress{})
	if err != nil {
		return err
	}
	return nil
}

func (r *TempAddressRepository) CreateTempAddress(doorno int32, street string, city string, zipcode int64, state string, country string) error {
	tempaddress := model.TempAddress{
		Door_no: doorno,
		Street:  street,
		City:    city,
		Zipcode: zipcode,
		State:   state,
		Country: country,
	}
	err := r.db.Create(&tempaddress).Error
	if err != nil {
		r.log.Errorf("failed to create temp address: %v", err)
		return errors.New("failed to create temp address")
	}
	return nil
}

func (r *TempAddressRepository) GetByIDTempAdress(id uint64) (*model.TempAddress, error) {
	var tempaddress model.TempAddress
	err := r.db.First(&tempaddress, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("temp address not found")
		}
		r.log.Errorf("failed to get temp address by ID: %v", err)
		return nil, errors.New("failed to get temp address by ID")
	}
	return &tempaddress, nil
}

func (r *TempAddressRepository) UpdateTempAddress(id uint64, doorno int32, street string, city string, zipcode int64, state string, country string) error {
	tempaddress := model.TempAddress{}
	r.db.First(&tempaddress, id)

	tempaddress.Door_no = doorno
	tempaddress.City = city
	tempaddress.Street = street
	tempaddress.Zipcode = zipcode
	tempaddress.State = state
	tempaddress.Country = country

	err := r.db.Save(&tempaddress).Error
	if err != nil {
		r.log.Errorf("failed to update temp address: %v", err)
		return errors.New("failed to update temp address")
	}
	return nil
}

func (r *TempAddressRepository) DeleteTempAddress(id uint64) error {
	err := r.db.Delete(&model.TempAddress{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete temp address: %v", err)
		return errors.New("failed to delete temp address")
	}
	return nil
}

func (r *TempAddressRepository) GetAllTempAddress() (*[]model.TempAddress, error) {
	var tempaddresses []model.TempAddress
	err := r.db.Find(&tempaddresses).Error
	if err != nil {
		r.log.Errorf("failed to retrieve temp addresss: %v", err)
		return &tempaddresses, errors.New("failed to get temp addresss list")
	}
	return &tempaddresses, nil
}
