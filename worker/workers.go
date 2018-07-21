package worker

// Workers is an array of workers
type Workers []Worker

// Configure configures workers
func (w Workers) Configure() error {
	for _, worker := range w {
		err := worker.GetJob().Configure()
		if err != nil {
			return err
		}
	}
	return nil
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
