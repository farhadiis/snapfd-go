package utils

import (
	"log"
	"os"
	"strconv"
)

func GetEnvInt(key string) int {
	val, err := strconv.Atoi(GetEnv(key))
	if err != nil {
		log.Fatalf("The '%s' can't convert to integer.", key)
	}
	return val
}

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("You must set your '%s' environmental variable.", key)
	}
	return val
}
