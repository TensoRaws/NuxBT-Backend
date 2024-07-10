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
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=allow",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)

	default: // 默认使用 mysql
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	}

	return dbConfig.Type, dsn, nil
}

// GenerateOSSPrefix 生成 OSS 对象存储的前缀
func GenerateOSSPrefix() string {
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
	useSSL := config.GetString("oss.ssl")
	var protocol string
	if useSSL == "true" {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	switch ossType {
	case "minio":
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	case "cos":
		OSS_PREFIX = fmt.Sprintf("https://%v.%v/", OSSConfig.Bucket, OSSConfig.Endpoint)
	default:
		// 默认使用 minio
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	}

	return OSS_PREFIX
}
