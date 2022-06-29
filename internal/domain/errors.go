package domain

import (
	"errors"
	"fmt"
)

var (
	ErrLogger_Config    error = fmt.Errorf("logger config error")
	ErrTelegramBot_Init error = fmt.Errorf("telegram Bot init error")
	ErrCommand_Init     error = errors.New("error on '%s' command")
)
