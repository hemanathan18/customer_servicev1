package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type PersonalMailRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewPersonalMailRepository(d *Data, logger log.Logger) *PersonalMailRepository {
	return &PersonalMailRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBPersMail(d *Data) error {

	err := d.db.AutoMigrate(&model.PersMail{})
	if err != nil {
		return err
	}
	return nil
}

func (r *PersonalMailRepository) CreatePersMail(email string) error {
	persmail := model.PersMail{Email: email}
	err := r.db.Create(&persmail).Error
	if err != nil {
		r.log.Errorf("failed to create personal mail: %v", err)
		return errors.New("failed to create personal mail")
	}
	return nil
}

func (r *PersonalMailRepository) GetByIDPersMail(id uint64) (*model.PersMail, error) {
	var persmail model.PersMail
	err := r.db.First(&persmail, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("personal mail not found")
		}
		r.log.Errorf("failed to get personal mail by ID: %v", err)
		return nil, errors.New("failed to get personal mail by ID")
	}
	return &persmail, nil
}

func (r *PersonalMailRepository) UpdatePersMail(id uint64, email string) error {
	persmail := model.PersMail{}
	r.db.First(&persmail, id)
	persmail.Email = email
	err := r.db.Save(&persmail).Error
	if err != nil {
		r.log.Errorf("failed to update personal mail: %v", err)
		return errors.New("failed to update personal mail")
	}
	return nil
}

func (r *PersonalMailRepository) DeletePersMail(id uint64) error {
	err := r.db.Delete(&model.PersMail{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete personal mail: %v", err)
		return errors.New("failed to delete personal mail")
	}
	return nil
}

func (r *PersonalMailRepository) GetAllPersMail() (*[]model.PersMail, error) {
	var persmails []model.PersMail
	err := r.db.Find(&persmails).Error
	if err != nil {
		r.log.Errorf("failed to retrieve personal mails: %v", err)
		return &persmails, errors.New("failed to get personal mail list")
	}
	return &persmails, nil
}
