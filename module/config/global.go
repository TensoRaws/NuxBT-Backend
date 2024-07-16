package config

import (
	"log"
)

var (
	ServerConfig   Server
	RegisterConfig Register
	JwtConfig      Jwt
	LogConfig      Log
	DBConfig       DB
	RedisConfig    Redis
	OSSConfig      OSS
	OSS_PREFIX     string
)

func setConfig() {
	err := config.UnmarshalKey("server", &ServerConfig)
	if err != nil {
		log.Fatalf("unable to decode into server struct, %v", err)
	}

	err = config.UnmarshalKey("register", &RegisterConfig)
	if err != nil {
		log.Fatalf("unable to decode into register struct, %v", err)
	}

	err = config.UnmarshalKey("jwt", &JwtConfig)
	if err != nil {
		log.Fatalf("unable to decode into jwt struct, %v", err)
	}

	err = config.UnmarshalKey("log", &LogConfig)
	if err != nil {
		log.Fatalf("unable to decode into log struct, %v", err)
	}

	err = config.UnmarshalKey("db", &DBConfig)
	if err != nil {
		log.Fatalf("unable to decode into db struct, %v", err)
	}

	err = config.UnmarshalKey("redis", &RedisConfig)
	if err != nil {
		log.Fatalf("unable to decode into redis struct, %v", err)
	}

	err = config.UnmarshalKey("oss", &OSSConfig)
	if err != nil {
		log.Fatalf("unable to decode into oss struct, %v", err)
	}

	OSS_PREFIX = GenerateOSSPrefix()
}
