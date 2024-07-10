package config

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	once   sync.Once
)

var (
	OSSConfig  OSS
	OSS_PREFIX string
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
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 配置文件未找到错误
			fmt.Println("config file not found use default config")
			config.SetDefault("server", map[string]interface{}{
				"port": 8080,
				"mode": "prod",
			})

			config.SetDefault("jwt", map[string]interface{}{
				"timeout": 60,
				"key":     "nuxbt",
			})

			config.SetDefault("log", map[string]interface{}{
				"level": "debug",
				"mode":  []string{"console", "file"},
				"path":  "./log",
			})

			config.SetDefault("db", map[string]interface{}{
				"type":     "mysql",
				"host":     "127.0.0.1",
				"port":     5432,
				"username": "root",
				"password": "123456",
				"database": "nuxbt",
				"ssl":      false,
			})

			config.SetDefault("redis", map[string]interface{}{
				"host":     "127.0.0.1",
				"port":     6379,
				"password": "123456",
				"poolSize": 10,
			})
		}
	}

	err := config.UnmarshalKey("server", &Server{})
	if err != nil {
		log.Fatalf("unable to decode into server struct, %v", err)
	}
	err = config.UnmarshalKey("jwt", &Jwt{})
	if err != nil {
		log.Fatalf("unable to decode into jwt struct, %v", err)
	}
	err = config.UnmarshalKey("log", &Log{})
	if err != nil {
		log.Fatalf("unable to decode into log struct, %v", err)
	}
	err = config.UnmarshalKey("db", &DB{})
	if err != nil {
		log.Fatalf("unable to decode into db struct, %v", err)
	}
	err = config.UnmarshalKey("redis", &Redis{})
	if err != nil {
		log.Fatalf("unable to decode into redis struct, %v", err)
	}
	err = config.UnmarshalKey("oss", &OSSConfig)
	if err != nil {
		log.Fatalf("unable to decode into oss struct, %v", err)
	}

	OSS_PREFIX = GenerateOSSPrefix()
	fmt.Printf("OSS TYPE: %v", config.GetString("oss.type"))
	fmt.Printf(" OSS PREFIX: %v\n", OSS_PREFIX)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}
