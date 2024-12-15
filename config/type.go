package config

type Config struct {
	Application         ApplicationConfig
	ClickerGameDatabase DBConfig
	Telegram            TelegramConfig
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

type TelegramConfig struct {
	BotToken        string
	ChannelUsername string
}
