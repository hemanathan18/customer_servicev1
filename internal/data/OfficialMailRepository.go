package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type OfficialMailRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewOfficialMailRepository(d *Data, logger log.Logger) *OfficialMailRepository {
	return &OfficialMailRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBOffMail(d *Data) error {

	err := d.db.AutoMigrate(&model.OffMail{})
	if err != nil {
		return err
	}
	return nil
}

func (r *OfficialMailRepository) CreateOffMail(email string) error {
	offmail := model.OffMail{Email: email}
	err := r.db.Create(&offmail).Error
	if err != nil {
		r.log.Errorf("failed to create official mail: %v", err)
		return errors.New("failed to create official mail")
	}
	return nil
}

func (r *OfficialMailRepository) GetByIDOffMail(id uint64) (*model.OffMail, error) {
	var offmail model.OffMail
	err := r.db.First(&offmail, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("official mail not found")
		}
		r.log.Errorf("failed to get official mail by ID: %v", err)
		return nil, errors.New("failed to get offcial mail by ID")
	}
	return &offmail, nil
}

func (r *OfficialMailRepository) UpdateOffMail(id uint64, email string) error {
	offmail := model.OffMail{}
	r.db.First(&offmail, id)
	offmail.Email = email
	err := r.db.Save(&offmail).Error
	if err != nil {
		r.log.Errorf("failed to update official mail: %v", err)
		return errors.New("failed to update official mail")
	}
	return nil
}

func (r *OfficialMailRepository) DeleteOffMail(id uint64) error {
	err := r.db.Delete(&model.OffMail{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete official mail: %v", err)
		return errors.New("failed to delete official mail")
	}
	return nil
}

func (r *OfficialMailRepository) GetAllOffMail() (*[]model.OffMail, error) {
	var offmails []model.OffMail
	err := r.db.Find(&offmails).Error
	if err != nil {
		r.log.Errorf("failed to retrieve official mails: %v", err)
		return &offmails, errors.New("failed to get official mail list")
	}
	return &offmails, nil
}
