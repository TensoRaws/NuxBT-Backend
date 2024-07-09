package config

type Server struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Jwt struct {
	Key string `yaml:"key"`
}

type Log struct {
	Level string   `yaml:"level"`
	Mode  []string `yaml:"mode"`
	Path  string   `yaml:"path"`
}

type DB struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolSize"`
	Database int    `yaml:"database"`
}

type OSS struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`

	// 如果是使用 minio，并且没有使用 https，需要设置为 false
	UseSsl *bool `yaml:"useSsl"`
	// 如果是使用 minio，需要设置为 true
	HostnameImmutable *bool `yaml:"hostnameImmutable"`
}
