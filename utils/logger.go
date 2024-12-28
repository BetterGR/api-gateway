package utils

import (
	"log"
)

var (
	DebugMode = true
)

func Debug(format string, v ...interface{}) {
	if DebugMode {
		log.Printf("[DEBUG] "+format, v...)
	}
}
