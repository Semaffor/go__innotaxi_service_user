package helpers

import (
	"log"
	"strconv"
)

func ConvertToInt(value string, defaultValue int) int {
	if value == "" {
		log.Printf("Assigned value is blank, setted default: %d", defaultValue)

		return defaultValue
	}
	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error when converting: %s, default: %d", value, defaultValue)

		return defaultValue
	}

	return convertedValue
}
