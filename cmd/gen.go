package cmd

import (
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/db"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/urfave/cli/v2"
	"gorm.io/gen"
)

// CmdGen 子命令
var CmdGen = &cli.Command{ //nolint:typecheck
	Name:        "gen",
	Usage:       "gen gorm code",
	Description: `GEN: Friendly & Safer GORM powered by Code Generation.`,
	Action:      runGen,
} //nolint:typecheck

func Init() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		ModelPkgPath: "dal/model",

		WithUnitTest: true,

		FieldNullable:     false,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db.DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}

func runGen(ctx *cli.Context) error { //nolint:typecheck
	config.Init()
	log.Init()
	db.Init()
	Init()
	return nil
}
