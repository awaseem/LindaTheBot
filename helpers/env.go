package helpers

import "os"

// GetEnvOrElse get enviroment variable or else
func GetEnvOrElse(key string, elseVal string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return elseVal
}
