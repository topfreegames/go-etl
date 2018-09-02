package processors

import "github.com/topfreegames/go-etl/models"

// Null is a processor that just passes input to output
type Null struct{}

// ProcessData implementation
func (*Null) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	outputChan <- d
}
