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
	Code Code
	etl  ETL
}

// Configure loads the job plugin
func (j *Job) Configure() error {
	if j.Code == "" {
		return j.configurePlugin("plugins")
	}

	return j.configureCode()
}

func (j *Job) configureCode() error {
	err := j.Code.Configure(j.Name)
	if err != nil {
		return err
	}

	return j.configurePlugin("codes")
}

func (j *Job) configurePlugin(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	mod := fmt.Sprintf("%s/%s/%s/main.so", dir, path, j.Name)
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
