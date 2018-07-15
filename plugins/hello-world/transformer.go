package main

import (
	"log"
	"time"

	"github.com/dailyburn/ratchet/data"
)

// Transformer implements ratchet.DataProcessor
type Transformer struct{}

// ProcessData implementation
func (t *Transformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	log.Print("executing transformer")
	time.Sleep(10 * time.Second)
	outputChan <- data.JSON("Hello World")
}

// Finish implementation
func (t *Transformer) Finish(outputChan chan data.JSON, killChan chan error) {}
