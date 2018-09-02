package models

// DataProcessor is a pipeline stage
type DataProcessor interface {
	ProcessData(
		input Data,
		outputChan chan Data,
		killChan chan error,
	)
}
