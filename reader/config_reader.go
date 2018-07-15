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

func (c *ConfigReader) Read() (worker.Workers, error) {
	var workers []*worker.Worker
	err := viper.UnmarshalKey("workers", &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}
