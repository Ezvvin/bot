package domain

import (
	"fmt"
)

var (
	ErrLogger_Config    error = fmt.Errorf("logger config error")
	ErrTelegramBot_Init error = fmt.Errorf("telegram Bot init error")
)
