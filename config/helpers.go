package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorMessages struct {
	ErrorMessage string
	ErrorType   string
}

func GetErrorMessages(err error) ErrorMessages {
	var errorMessages ErrorMessages
	switch err := err.(type) {
	case validator.ValidationErrors:
		for _, e := range err {
			errorMessages.ErrorMessage += fmt.Sprintf("%s %s\n", e.Field(), e.Error())
		}
	case string:
		errorMessages.ErrorMessage = err
	default:
		log.Println(err)
	}
	return errorMessages
}

func GetDateTimeFromTS(ts int64) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", ts/86400%10000, ts/86400%100%100, ts/3600%24, ts/3600%60, ts/60%60, ts%60)
}

func GetFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func GetFileSize(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := info.Size()
	return int(size)
}

func GetIntFromEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return i
}