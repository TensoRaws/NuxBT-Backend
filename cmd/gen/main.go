package main

import (
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func ConnectDB(dbType, dsn string) (db *gorm.DB) {
	var err error

	log.Logger.Debugf("DSN: %v", dsn)

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Logger.Fatalf("connect db fail: %v", err)
	}

	return db
}

func Init() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		ModelPkgPath: "model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,

		WithUnitTest: false,

		FieldNullable:     false,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	dbType, dsn, err := config.GenerateDSN()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	DB := ConnectDB(dbType, dsn)

	g.UseDB(DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}

func main() {
	config.Init()
	log.Init()
	Init()
}
