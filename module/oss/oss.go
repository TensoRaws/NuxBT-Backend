package oss

import (
	"context"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"io"
	"sync"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/eleven26/goss/v4"
)

var (
	oss  *goss.Goss
	err  error
	once sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	var cfg = &goss.Config{
		Endpoint:          config.OSSConfig.Endpoint,
		AccessKey:         config.OSSConfig.AccessKey,
		SecretKey:         config.OSSConfig.SecretKey,
		Region:            config.OSSConfig.Region,
		Bucket:            config.OSSConfig.Bucket,
		UseSsl:            &config.OSSConfig.UseSSL,
		HostnameImmutable: &config.OSSConfig.HostnameImmutable,
	}
	oss, err = goss.New(goss.WithConfig(cfg))
	if err != nil {
		log.Logger.Errorf("init goss faild: %v", err)
	}
}

// Put saves the content read from r to the key of oss.
func Put(key string, r io.Reader) error {
	return oss.Put(context.TODO(), key, r)
}

// PutFromFile saves the file pointed to by the `localPath` to the oss key.
func PutFromFile(key string, localPath string) error {
	return oss.PutFromFile(context.TODO(), key, localPath)
}

// Get gets the file pointed to by key.
func Get(key string) (io.ReadCloser, error) {
	return oss.Get(context.TODO(), key)
}

// GetString gets the file pointed to by key and returns a string.
func GetString(key string) (string, error) {
	return oss.GetString(context.TODO(), key)
}

// GetBytes gets the file pointed to by key and returns a byte array.
func GetBytes(key string) ([]byte, error) {
	return oss.GetBytes(context.TODO(), key)
}

// GetToFile saves the file pointed to by key to the localPath.
func GetToFile(key string, localPath string) error {
	return oss.GetToFile(context.TODO(), key, localPath)
}

// Delete the file pointed to by key.
func Delete(key string) error {
	return oss.Delete(context.TODO(), key)
}

// Exists determines whether the file exists.
func Exists(key string) (bool, error) {
	return oss.Exists(context.TODO(), key)
}

// Size fet the file size.
func Size(key string) (int64, error) {
	return oss.Size(context.TODO(), key)
}
