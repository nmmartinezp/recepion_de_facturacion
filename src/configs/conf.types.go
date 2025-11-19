package configs

type AppConfig struct {
	Port string
	Env  string
}

type DBConfig struct {
	Database string
	Port     string
	User     string
	Password string
	Host     string
}

type AuthConfig struct {
	JWTSecret string
}

type Configuration struct {
	App      AppConfig
	DB       map[string]DBConfig
	Auth     AuthConfig
	RabbitMQ RabbitMQConfig
}

type RabbitMQConfig struct {
	URL      string
	Exchange string
}
