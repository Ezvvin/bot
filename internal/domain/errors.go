package domain

import (
	"fmt"
)

var (
	LoggerConfigError     error = fmt.Errorf("Logger config error")
	TelegramBotError_Init error = fmt.Errorf("Telegram Bot init error")
)
