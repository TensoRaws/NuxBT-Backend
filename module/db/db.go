package db

import (
	"sync"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	dbType, dsn, err := config.GenerateDSN()
	if err != nil {
		log.Logger.Error(err)
		return
	}

	cfg := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	DB = ConnectDB(dbType, dsn, &cfg)

	err = DB.AutoMigrate(
		model.User{},
		model.UserRole{},
		model.Torrent{},
	)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	query.SetDefault(DB)
	log.Logger.Debugf("Set query default database")
}

func ConnectDB(dbType, dsn string, config *gorm.Config) (db *gorm.DB) {
	var err error

	log.Logger.Debugf("DBType: %v", dbType)
	log.Logger.Debugf("DSN: %v", dsn)

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), config)
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), config)
	default:
		db, err = gorm.Open(mysql.Open(dsn), config)
	}

	if err != nil {
		log.Logger.Fatalf("connect db fail: %v", err)
	}

	return db
}
