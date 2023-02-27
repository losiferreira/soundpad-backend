package shared

import (
	"os"
	"strconv"
)

func GetOsBoolEnv(key string) bool {
	stringValue := os.Getenv(key)
	boolValue, _ := strconv.ParseBool(stringValue)
	return boolValue
}
