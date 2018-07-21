package reader

import (
	"github.com/Henrod/go-etl/worker"
	"github.com/spf13/viper"
)

// ConfigReader implements reader.Reader by reading from a config file
type ConfigReader struct{}

// NewConfigReader returns a *ConfigReader
func NewConfigReader() *ConfigReader {
	return &ConfigReader{}
}

func (c *ConfigReader) Read() (workers worker.Workers, err error) {
	var specs []map[string]interface{}
	err = viper.UnmarshalKey("workers", &specs)
	if err != nil {
		return nil, err
	}

	workers = make(worker.Workers, len(specs))
	for idx, spec := range specs {
		workers[idx], err = worker.NewWorker(spec)
	}

	err = viper.UnmarshalKey("workers", &workers)
	if err != nil {
		return nil, err
	}

	err = workers.Validate()
	if err != nil {
		return nil, err
	}

	return workers, nil
}
