package db

import (
	"gorm.io/gorm/schema"
	"sync"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	//var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	//	// int mapping
	//	"int": func(columnType gorm.ColumnType) (dataType string) {
	//		if n, ok := columnType.Nullable(); ok && n {
	//			return "*int64"
	//		}
	//		return "int64"
	//	},
	//}

	var cfg = gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	DB = ConnectDB(dbType, dsn, &cfg)

	err = DB.AutoMigrate(
		model.User{},
		model.UserRole{},
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
