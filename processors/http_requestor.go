package processors

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dailyburn/ratchet/data"
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
func (h *HTTPRequestor) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	log.Print("calling http requestor process")

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

// Finish implementation
func (h *HTTPRequestor) Finish(outputChan chan data.JSON, killChan chan error) {}
