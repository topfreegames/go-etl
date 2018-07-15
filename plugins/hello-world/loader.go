package main

import (
	"fmt"
	"log"

	"github.com/dailyburn/ratchet/data"
)

// Loader implements ratchet.DataProcessor
type Loader struct{}

// ProcessData implementation
func (l *Loader) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	log.Print("executing loader")
	fmt.Printf("loader: %s\n", string(d))
}

// Finish implementation
func (l *Loader) Finish(outputChan chan data.JSON, killChan chan error) {}
