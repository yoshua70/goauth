package main

import (
	"log"
)

// FailOnError function
// panics if the provided error is not nil. Used to reduce error handling
// boilerplate.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err)
	}
}
