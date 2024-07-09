package db

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	DB = ConnectDB(config.MySQLDSN())
	err := DB.AutoMigrate(
		model.User{},
	)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	query.SetDefault(DB)
	log.Logger.Debugf("Set query default database")
}

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	log.Logger.Debugf("MySQL DSN: %v", config.MySQLDSN())

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Logger.Fatalf("connect db fail: %v", err)
	}

	return db
}
