package utils

import (
	"log"
	"strings"
)

//HandleError will panic and log errors
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}

}

//LowercaseString will return a lowercased string value
func LowercaseString(str string) string {
	return strings.ToLower(str)
}
