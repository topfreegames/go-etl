package models

import (
	"log"
)

// Pipeline contains stages and executes them in order
type Pipeline struct {
	stages []DataProcessor
}

// NewPipeline create pipeline with stages
func NewPipeline(stages ...DataProcessor) *Pipeline {
	return &Pipeline{
		stages: stages,
	}
}

// Run executes stages and return chan of error at first error on at
// the end of stages
func (p *Pipeline) Run() chan error {
	killChan := make(chan error)

	go p.run(killChan)

	return killChan
}

func (p *Pipeline) run(killChan chan error) {
	log.Print("starting pipeline")

	outputChan := make(chan Data, 1)
	outputChan <- NewData()

	for idx, stage := range p.stages {
		select {
		case err := <-killChan:
			killChan <- err
			return
		default:
			log.Printf("running stage %d", idx)
			stage.ProcessData(<-outputChan, outputChan, killChan)
		}
	}

	killChan <- nil
}
