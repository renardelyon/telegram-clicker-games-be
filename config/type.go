package config

type Config struct {
	Application         ApplicationConfig `envPrefix:"APPLICATION_"`
	ClickerGameDatabase DBConfig          `envPrefix:"CLICKERGAME_DB_"`
	Telegram            TelegramConfig    `envPrefix:"TELEGRAM_"`
}

type ApplicationConfig struct {
	ServiceName    string `env:"SERVICE_NAME" envDefault:"default-service"`
	ServiceVersion string `env:"SERVICE_VERSION" envDefault:"1.0.0"`
	ServerPort     string `env:"SERVER_PORT" envDefault:"8080"`
}

type DBConfig struct {
	Host     string `env:"HOST,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Database string `env:"DATABASE,required"`
}

type TelegramConfig struct {
	BotToken        string `env:"BOT_TOKEN,required"`
	ChannelUsername string `env:"CHANNEL_USERNAME,required"`
	BotApp          string `env:"BOT_APP,required"`
}
