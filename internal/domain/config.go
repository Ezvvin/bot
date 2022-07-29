package domain

type Config struct {
	LogLevel    string `default:"DEBUG"`
	AdminChat   int64  `default:""`
	AdminChanel int64  `default:""`
	ApiKeyMap   string `default:""`

	BotConfig BotConfig
}
type BotConfig struct {
	Token    string `default:""`
	PayToken string `default:""`
	Timeout  int    `default:"60"`
	Debug    bool   `default:"false"`
}
