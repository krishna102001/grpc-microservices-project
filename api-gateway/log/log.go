package log

import logger "github.com/charmbracelet/log"

func Info(detail string) {
	logger.Info(detail)
}

func Error(detail string) {
	logger.Error(detail)
}

func Fatal(detail string) {
	logger.Fatal(detail)
}
