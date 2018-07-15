package worker

import (
	"log"
	"time"

	"github.com/Henrod/go-etl/models"
)

// Worker executes etl every period of time
type Worker struct {
	Period time.Duration
	Job    *models.Job
}

// Start runs the worker every period
func (w *Worker) Start(done chan struct{}) {
	ticker := time.NewTicker(w.Period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("executing job %s", w.Job.Name)
			err := w.Job.Execute()
			w.handleErr(err)
		case <-done:
			log.Print("terminating worker")
			return
		}
	}
}

func (w *Worker) handleErr(err error) {
	if err != nil {
		log.Printf("job failed: %s", err.Error())
	} else {
		log.Print("executed job")
	}
}
