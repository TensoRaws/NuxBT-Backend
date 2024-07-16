package config

func configSetDefault() {
	config.SetDefault("server", map[string]interface{}{
		"port":                      8080,
		"mode":                      "prod",
		"allowRegister":             true,
		"useInvitationCode":         false,
		"inviteCodeEligibilityTime": 30,
		"requestLimit":              50,
		"cros":                      []string{},
	})

	config.SetDefault("jwt", map[string]interface{}{
		"timeout": 60,
		"key":     "nuxbt",
	})

	config.SetDefault("log", map[string]interface{}{
		"level": "debug",
		"mode":  []string{"console", "file"},
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