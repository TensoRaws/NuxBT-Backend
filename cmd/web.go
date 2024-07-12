package cmd

import (
	"fmt"

	"github.com/TensoRaws/NuxBT-Backend/internal/router"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/db"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/oss"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{
	Name:        "server",
	Usage:       "Start NuxBT api server",
	Description: `Star NuxBT api server`,
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(ctx *cli.Context) error {
	defer func() {
		for k := range cache.Clients {
			err := cache.Clients[k].C.Close()
			if err != nil {
				fmt.Printf("close redis: %v", err)
				return
			}
		}
	}()
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	cache.Init()
	router.Init()
	return nil
}
