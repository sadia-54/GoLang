package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func InitLogger() {
	// Open file for logging (append mode)
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Console writer (human-readable) for development
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Multi output: file + console
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Create logger
	l := zerolog.New(multi).With().Timestamp().Logger()

	Logger = &l

	// Set global log level (InfoLevel or DebugLevel for dev)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
