package processors

import (
	"log"

	"github.com/dailyburn/ratchet/data"
)

// Logger is a processor that just logs and passes input to output
type Logger struct{}

// ProcessData implementation
func (l *Logger) ProcessData(
	d data.JSON,
	outputChan chan data.JSON,
	killChan chan error,
) {
	log.Printf("logger worker: %s", string(d))
	outputChan <- d
}

// Finish implementation
func (l *Logger) Finish(
	outputChan chan data.JSON,
	killChan chan error,
) {
}
