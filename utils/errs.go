package utils

import (
	"errors"
	"log"
)

var (
	GetFileError = errors.New("get files error")
	GetHttpError = errors.New("get http server error")
)

func Failure(err error) {
	log.Fatalf("loader catch error: %v\n", err)
}
