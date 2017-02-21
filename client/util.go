package client

import (
	"fmt"
	"net/http"
)

var (
	// ErrResponseNotOk is returned when an API response should not be considered valid.
	ErrResponseNotOk = fmt.Errorf("API response not OK")
)

func isResponseStatusOK(resp *http.Response) bool {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return false
	}
	return true
}
