package main

import (
	"log"
	"time"

	"github.com/topfreegames/go-etl/models"
)

// Transformer implements ratchet.DataProcessor
type Transformer struct{}

// ProcessData implementation
func (t *Transformer) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	log.Print("executing transformer")
	time.Sleep(10 * time.Second)
	outputChan <- models.Data("Hello World")
}
