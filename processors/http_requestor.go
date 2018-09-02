package processors

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/topfreegames/go-etl/models"
)

// HTTPRequestor is a DataProcessor that makes a HTTP request
type HTTPRequestor struct {
	method, url string
}

// NewHTTPRequestor is the HTTPRequestor constructor
func NewHTTPRequestor(method, url string) *HTTPRequestor {
	return &HTTPRequestor{method, url}
}

// ProcessData implementation
func (h *HTTPRequestor) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	fmt.Println("calling http requestor process")

	req, err := http.NewRequest(h.method, h.url, nil)
	if err != nil {
		killChan <- err
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		killChan <- err
		return
	}

	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		killChan <- err
		return
	}

	outputChan <- bts
}
