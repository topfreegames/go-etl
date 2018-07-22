package processors

import (
	"github.com/dailyburn/ratchet/data"
)

// Null is a processor that just passes input to output
type Null struct {
}

// ProcessData implementation
func (*Null) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	outputChan <- d
}

// Finish implementation
func (*Null) Finish(outputChan chan data.JSON, killChan chan error) {}
