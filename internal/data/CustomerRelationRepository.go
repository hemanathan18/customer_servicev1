package data

import (
	"errors"

	"project/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CustomerRelationRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewCustomerRelationRepository(d *Data, logger log.Logger) *CustomerRelationRepository {
	return &CustomerRelationRepository{
		db:  d.db,
		log: d.log,
	}
}

func MigrateDBCustomerRelation(d *Data) error {

	err := d.db.AutoMigrate(&model.CustomerRelation{})
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRelationRepository) CreateRelation(cusId uint64, offeId uint64, perseId uint64,
	offpId uint64, perspId uint64, tAddId uint64, pAddId uint64) error {
	customerrelation := model.CustomerRelation{
		Customer_ID: cusId,
		Off_em_id:   offeId,
		Pers_em_id:  perseId,
		Off_ph_id:   offpId,
		Pers_ph_id:  perspId,
		Temp_add_id: tAddId,
		Perm_add_id: pAddId,
	}

	err := r.db.Create(&customerrelation).Error
	if err != nil {
		r.log.Errorf("failed to create customer relation: %v", err)
		return errors.New("failed to create customer relation")
	}
	return nil
}

func (r *CustomerRelationRepository) GetByIDRelation(id uint64) (*model.CustomerRelation, error) {
	var customerrelation model.CustomerRelation
	err := r.db.First(&customerrelation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		r.log.Errorf("failed to get customer relation by ID: %v", err)
		return nil, errors.New("failed to get customer relation by ID")
	}
	return &customerrelation, nil
}

func (r *CustomerRelationRepository) UpdateRelation(id uint64, cusId uint64, offeId uint64, perseId uint64,
	offpId uint64, perspId uint64, tAddId uint64, pAddId uint64) error {
	customerrelation := model.CustomerRelation{}
	r.db.First(&customerrelation, id)

	customerrelation.Customer_ID = cusId
	customerrelation.Off_em_id = offeId
	customerrelation.Pers_em_id = perseId
	customerrelation.Off_ph_id = offpId
	customerrelation.Pers_ph_id = perspId
	customerrelation.Temp_add_id = tAddId
	customerrelation.Perm_add_id = pAddId

	err := r.db.Save(&customerrelation).Error
	if err != nil {
		r.log.Errorf("failed to update customer relation: %v", err)
		return errors.New("failed to update customer relation")
	}
	return nil
}

func (r *CustomerRelationRepository) DeleteRelation(id uint64) error {
	err := r.db.Delete(&model.CustomerRelation{}, id).Error
	if err != nil {
		r.log.Errorf("failed to delete customer relation: %v", err)
		return errors.New("failed to delete customer relation")
	}
	return nil
}

func (r *CustomerRelationRepository) GetAllRelation() (*[]model.CustomerRelation, error) {
	var customerrelations []model.CustomerRelation
	err := r.db.Find(&customerrelations).Error
	if err != nil {
		r.log.Errorf("failed to retrieve customer relations: %v", err)
		return &customerrelations, errors.New("failed to get customer relations list")
	}
	return &customerrelations, nil
}
