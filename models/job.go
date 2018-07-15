package models

import (
	"fmt"
	"os"
	"plugin"

	"github.com/dailyburn/ratchet"
)

// Job has the ETL description
type Job struct {
	Name string
	etl  ETL
}

// Configure loads the job plugin
func (j *Job) Configure() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	mod := fmt.Sprintf("%s/plugins/%s/main.so", dir, j.Name)
	plug, err := plugin.Open(mod)
	if err != nil {
		return err
	}

	symETL, err := plug.Lookup("ETL")
	if err != nil {
		return err
	}

	etl, ok := symETL.(ETL)
	if !ok {
		return fmt.Errorf("plugin %s doesn't implement ETL interface", j.Name)
	}

	j.etl = etl
	return nil
}

// Execute ...
func (j *Job) Execute() error {
	pipeline := ratchet.NewPipeline(
		j.etl.Extract(),
		j.etl.Transform(),
		j.etl.Load(),
	)

	return <-pipeline.Run()
}
