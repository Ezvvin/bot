package domain

import (
	"fmt"
)

var (
	LoggerConfigError     error = fmt.Errorf("logger config error")
	TelegramBotError_Init error = fmt.Errorf("telegram Bot init error")
)
