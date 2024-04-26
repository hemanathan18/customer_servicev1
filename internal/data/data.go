package data

import (
	"project/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCustomerRepository, NewOfficialMailRepository, NewOfficialNumberRepository,
	NewPermAdressRepository, NewPersonalMailRepository, NewPersonalNumberRepository,
	NewTempAdressRepository, NewCustomerRelationRepository, NewDB)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}
	return db
}

func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	d := &Data{
		db:  db,
		log: log.NewHelper(logger),
	}

	return d, cleanup, nil
}
