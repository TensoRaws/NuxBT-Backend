package config

import (
	"fmt"
)

// GenerateDSN 根据数据库类型生成相应的 DSN 字符串, 并返回数据库类型和 DSN 字符串
func GenerateDSN() (string, string, error) {
	var dbConfig DB
	err := config.UnmarshalKey("db", &dbConfig)
	if err != nil {
		return "", "", err
	}

	var dsn string

	switch dbConfig.Type {
	case "mysql":
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)

	default: // 默认使用 mysql
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	}

	return dbConfig.Type, dsn, nil
}
