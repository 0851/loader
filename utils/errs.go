package utils

import (
	"errors"
	"log"
	errs "github.com/pkg/errors"
)

var (
	GetHttpError = errors.New("get http server error")
)

func Failure(err error, message string) {
	if err == nil {
		return
	}
	err = errs.Wrap(err, message)
	log.Fatalf("loader catch error: %+v\n", err)
}
