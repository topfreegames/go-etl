package worker

import "log"

// Workers is an array of workers
type Workers []Worker

// Configure configures workers
func (w Workers) Configure() Workers {
	correctWorkers := Workers{}

	for _, worker := range w {
		err := worker.GetJob().Configure()
		if err != nil {
			name := worker.GetJob().Name
			msg := err.Error()
			log.Printf("failed to configure job %s: %s", name, msg)
		} else {
			correctWorkers = append(correctWorkers, worker)
		}
	}

	return correctWorkers
}

// Validate validates all workers in worker
func (w Workers) Validate() error {
	for _, worker := range w {
		if err := worker.Validate(); err != nil {
			return err
		}
	}

	return nil
}
