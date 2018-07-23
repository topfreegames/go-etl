package worker

import (
	"errors"
	"log"
	"time"

	"github.com/topfreegames/go-etl/models"
)

// PeriodicWorker executes etl every period of time
type PeriodicWorker struct {
	Period time.Duration
	Job    *models.Job
}

// NewPeriodicWorker returns an PeriodicWorker
func NewPeriodicWorker() Worker {
	return &PeriodicWorker{}
}

// GetJob returns the job
func (p *PeriodicWorker) GetJob() *models.Job {
	return p.Job
}

// Validate validates spec
func (p *PeriodicWorker) Validate() error {
	if p.Period.Seconds() <= 0 {
		return errors.New("invalid period")
	}

	return nil
}

// Start runs the worker every period
func (p *PeriodicWorker) Start(done chan struct{}) {
	ticker := time.NewTicker(p.Period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("executing job %s", p.Job.Name)
			err := p.Job.Execute()
			handleErr(err)
		case <-done:
			log.Print("terminating worker")
			return
		}
	}
}
