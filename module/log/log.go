package log

import (
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/charmbracelet/log"
	"os"
	"sync"
	"time"
)

var (
	Logger *log.Logger
	once   sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339Nano,
		Prefix:          "NuxBT-Backend",
	})

	level := config.GetString("log.level")
	switch level {
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "info":
		Logger.SetLevel(log.InfoLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	case "error":
		Logger.SetLevel(log.ErrorLevel)
	default:
		Logger.SetLevel(log.DebugLevel)
	}
	Logger.Debugf("log level: %v", level)
}
