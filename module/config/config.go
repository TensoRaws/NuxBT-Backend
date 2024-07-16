package config

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	once   sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	config = viper.New()

	config.SetConfigName("nuxbt")
	config.AddConfigPath("./conf/")
	config.AddConfigPath("./")
	config.AddConfigPath("$HOME/.nuxbt/")
	config.AddConfigPath("/etc/nuxbt/")
	config.SetConfigType("yml")

	config.AutomaticEnv()
	config.SetEnvPrefix("NUXBT")
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后，重新初始化配置
		setConfig()
		fmt.Println("Config file changed:", e.Name)
	})

	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("config file not found use default config")
			configSetDefault()
		}
	}

	// 初始化配置
	setConfig()

	fmt.Printf("OSS TYPE: %v", config.GetString("oss.type"))
	fmt.Printf(" OSS PREFIX: %v\n", OSS_PREFIX)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}
