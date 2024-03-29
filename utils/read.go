package utils

import (
	"io/ioutil"
	"net/http"
	"os"
)

func readAsHttp(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer func() {
		err := resp.Body.Close()
		Failure(err, "read file error")
	}()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		return body, err
	} else {
		return nil, GetHttpError
	}
}
func readAsFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	defer func() {
		err := f.Close()
		Failure(err, "read file error")
	}()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(f)
	return body, err
}
