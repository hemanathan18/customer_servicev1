package data

import (
	"errors"
	"time"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewCustomerRepository(d *Data, logger log.Logger) *CustomerRepository {
	return &CustomerRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDB(d *Data) error {

	err := d.db.AutoMigrate(&model.Customer{})
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) Create(name string, gender string, dob time.Time) error {
	customer := model.Customer{Name: name, Gender: gender, DOB: dob}
	err := r.db.Create(&customer).Error
	if err != nil {
		r.log.Errorf("failed to create customer: %v", err)
		return errors.New("failed to create customer")
	}
	return nil
}

func (r *CustomerRepository) GetByID(id uint64) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		r.log.Errorf("failed to get customer by ID: %v", err)
		return nil, errors.New("failed to get customer by ID")
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(id uint64, name string, gender string, dob time.Time) error {
	customer := model.Customer{}
	r.db.First(&customer, id)
	customer.Name = name
	customer.Gender = gender
	customer.DOB = dob
	err := r.db.Save(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		r.log.Errorf("failed to update customer: %v", err)
		return errors.New("failed to update customer")
	}
	return nil
}

func (r *CustomerRepository) Delete(id uint64) error {
	err := r.db.Delete(&model.Customer{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		r.log.Errorf("failed to delete customer: %v", err)
		return errors.New("failed to delete customer")
	}
	return nil
}

func (r *CustomerRepository) GetAll() (*[]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	if err != nil {
		r.log.Errorf("failed to retrieve customers: %v", err)
		return &customers, errors.New("failed to get customers list")
	}
	return &customers, nil
}
