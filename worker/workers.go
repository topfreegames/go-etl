package worker

// Workers is an array of workers
type Workers []*Worker

// Configure configures workers
func (w Workers) Configure() error {
	for _, worker := range w {
		err := worker.Job.Configure()
		if err != nil {
			return err
		}
	}
	return nil
}
