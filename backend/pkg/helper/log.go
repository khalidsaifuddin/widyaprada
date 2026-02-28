package helper

import (
	"encoding/json"
	"log"
)

func DebugLogJson(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error marshalling to JSON: %v", err)
		return ""
	}
	return string(b)
}

func LogError(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func LogInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func LogDebug(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}
