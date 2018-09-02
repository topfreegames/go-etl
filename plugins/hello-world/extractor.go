package main

import (
	"log"

	"github.com/topfreegames/go-etl/models"
)

// Extractor implements ratchet.DataProcessor
type Extractor struct{}

// ProcessData implementation
func (e *Extractor) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	log.Print("executing extractor")
	outputChan <- models.Data("Hello World")
}
