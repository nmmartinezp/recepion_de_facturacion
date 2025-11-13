package configs

func VarConfig() Configuration {
	config := Configuration{
		App: AppConfig{
			Port: getEnv("PORT", "3002"),
			Env:  getEnv("ENV", "DEV"),
		},
		DB: map[string]DBConfig{
			"DB_POSTGRES": {
				Database: getEnv("DB_POSTGRES_NAME", "test"),
				Port:     getEnv("DB_POSTGRES_PORT", "3306"),
				User:     getEnv("DB_POSTGRES_USER", "root"),
				Password: getEnv("DB_POSTGRES_PASSWORD", ""),
				Host:     getEnv("DB_POSTGRES_HOST", "localhost"),
			}},
		Auth: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", "your_jwt_secret_key"),
		},
	}

	return config
}
