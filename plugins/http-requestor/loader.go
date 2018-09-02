package main

import (
	"fmt"
	"log"

	"github.com/topfreegames/go-etl/models"
)

// Loader implements ratchet.DataProcessor
type Loader struct{}

// ProcessData implementation
func (l *Loader) ProcessData(d models.Data, outputChan chan models.Data, killChan chan error) {
	log.Print("executing loader")
	fmt.Printf("loader: %s\n", string(d))
}
