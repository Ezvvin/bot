package domain

type Config struct {
	LogLevel  string `default:"DEBUG"`
	BotConfig BotConfig
}
type BotConfig struct {
	Token   string `default:""`
	Timeout int    `default:"60"`
	Debug   bool   `default:"false"`
}
