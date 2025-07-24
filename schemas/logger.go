package schemas

import (
	"time"
)

type LogLevel string

const (
	Debug   LogLevel = "DEBUG"
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARNING"
	Error   LogLevel = "ERROR"
)

type LogMessage struct {
	Level     LogLevel  `json:"level"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Service   string    `json:"service"`
}
