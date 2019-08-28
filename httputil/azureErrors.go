package httputil

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func RelevantError(httpError error, resp *http.Response) error {
	if httpError != nil {
		return httpError
	}
	if code := resp.StatusCode; code < 200 || code > 299 {
		return errors.New(fmt.Sprintf("Received %s from azure request", strconv.Itoa(code)))
	}
	return nil
}
