package processors

import (
	"log"

	"github.com/topfreegames/go-etl/models"
)

// Logger is a processor that just logs and passes input to output
type Logger struct{}

// ProcessData implementation
func (l *Logger) ProcessData(
	d models.Data,
	outputChan chan models.Data,
	killChan chan error,
) {
	log.Printf("logger worker: %s", string(d))
	outputChan <- d
}
