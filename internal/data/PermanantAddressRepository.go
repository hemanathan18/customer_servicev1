package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type PermAddressRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewPermAdressRepository(d *Data, logger log.Logger) *PermAddressRepository {
	return &PermAddressRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBPermAddress(d *Data) error {

	err := d.db.AutoMigrate(&model.PermAddress{})
	if err != nil {
		return err
	}
	return nil
}

func (r *PermAddressRepository) CreatePermAddress(doorno int32, street string, city string, zipcode int64, state string, country string) error {
	permaddress := model.PermAddress{
		Door_no: doorno,
		Street:  street,
		City:    city,
		Zipcode: zipcode,
		State:   state,
		Country: country,
	}
	err := r.db.Create(&permaddress).Error
	if err != nil {
		r.log.Errorf("failed to create perm address: %v", err)
		return errors.New("failed to create perm address")
	}
	return nil
}

func (r *PermAddressRepository) GetByIDPermAdress(id uint64) (*model.PermAddress, error) {
	var permaddress model.PermAddress
	err := r.db.First(&permaddress, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("perm address not found")
		}
		r.log.Errorf("failed to get perm address by ID: %v", err)
		return nil, errors.New("failed to get perm address by ID")
	}
	return &permaddress, nil
}

func (r *PermAddressRepository) UpdatePermAddress(id uint64, doorno int32, street string, city string, zipcode int64, state string, country string) error {
	permaddress := model.PermAddress{}
	r.db.First(&permaddress, id)

	permaddress.Door_no = doorno
	permaddress.City = city
	permaddress.Street = street
	permaddress.Zipcode = zipcode
	permaddress.State = state
	permaddress.Country = country

	err := r.db.Save(&permaddress).Error
	if err != nil {
		r.log.Errorf("failed to update perm address: %v", err)
		return errors.New("failed to update perm address")
	}
	return nil
}

func (r *PermAddressRepository) DeletePermAddress(id uint64) error {
	err := r.db.Delete(&model.PermAddress{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete perm address: %v", err)
		return errors.New("failed to delete perm address")
	}
	return nil
}

func (r *PermAddressRepository) GetAllPermAddress() (*[]model.PermAddress, error) {
	var permaddresses []model.PermAddress
	err := r.db.Find(&permaddresses).Error
	if err != nil {
		r.log.Errorf("failed to retrieve perm addresss: %v", err)
		return &permaddresses, errors.New("failed to get perm addresss list")
	}
	return &permaddresses, nil
}
