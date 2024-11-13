package env

import (
	"os"
	"strconv"
)

// used to fetch env variables with a fallback option. helps us ensure there is a default value if the variable is not set or has incorrect format

// retrieves a string env variable by its key
// key = name of variable we want to get
// fallback = the value to return if the env is not set
func GetString(key, fallback string) string {

	// check if the env variable exists
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// retrieves an int env variable by its key
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
