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
				"key": "nuxbt",
			})

			config.SetDefault("log", map[string]interface{}{
				"level": "debug",
				"mode":  []string{"console", "file"},
				"path":  "./log",
			})

			config.SetDefault("db", map[string]interface{}{
				"type":     "postgres",
				"host":     "127.0.0.1",
				"port":     3306,
				"username": "root",
				"password": "123456",
				"database": "nuxbt",
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
	err = config.UnmarshalKey("oss", &OSS{})
	if err != nil {
		log.Fatalf("unable to decode into oss struct, %v", err)
	}

	// use env var to set oss config when some field is nil
	if OSSConfig.Endpoint == "" {
		OSSConfig.Endpoint = config.GetString("oss.endpoint")
	}
	if OSSConfig.AccessKey == "" {
		OSSConfig.AccessKey = config.GetString("oss.accesskey")
	}
	if OSSConfig.SecretKey == "" {
		OSSConfig.SecretKey = config.GetString("oss.secretkey")
	}
	if OSSConfig.Region == "" {
		OSSConfig.Region = config.GetString("oss.region")
	}
	if OSSConfig.Bucket == "" {
		OSSConfig.Bucket = config.GetString("oss.bucket")
	}

	ossType := config.GetString("oss.type")
	useSsl := config.GetString("oss.useSsl")
	var protocol string
	if useSsl == "true" {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	fmt.Printf("oss type: %v\n", ossType)
	switch ossType {
	case "minio":
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	case "cos":
		OSS_PREFIX = fmt.Sprintf("https://%v.%v/", OSSConfig.Bucket, OSSConfig.Endpoint)
	default:
		// 默认使用 minio
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	}
	fmt.Printf("OSS PREFIX: %v\n", OSS_PREFIX)
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}
