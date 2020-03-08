package utils

import "log"

//HandleError will panic and log errors
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}

}
