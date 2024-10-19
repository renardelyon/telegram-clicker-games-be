package config

type Config struct {
	Application         ApplicationConfig
	ClickerGameDatabase DBConfig
}

type ApplicationConfig struct {
	ServiceName    string
	ServiceVersion string
	ServerPort     string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Database string
}
