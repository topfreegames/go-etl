package main

import (
	"log"

	"github.com/dailyburn/ratchet/data"
)

// Extractor implements ratchet.DataProcessor
type Extractor struct{}

// ProcessData implementation
func (e *Extractor) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	log.Print("executing extractor")
	outputChan <- data.JSON("Hello World")
}

// Finish implementation
func (e *Extractor) Finish(outputChan chan data.JSON, killChan chan error) {}
