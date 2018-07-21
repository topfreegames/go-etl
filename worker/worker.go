package worker

import (
	"errors"
	"fmt"
	"log"

	"github.com/Henrod/go-etl/models"
)

// Worker has a start method
type Worker interface {
	Start(done chan struct{})
	GetJob() *models.Job
	Validate() error
}

// NewWorker returns a worker from spec
func NewWorker(spec map[string]interface{}) (Worker, error) {
	workers := map[string]func() Worker{
		"period":   NewPeriodicWorker,
		"schedule": NewScheduleWorker,
	}

	var ctor func() Worker

	for name, workerCtor := range workers {
		if _, ok := spec[name]; ok {
			if ctor != nil {
				return nil, errors.New("must have only one worker type")
			}
			ctor = workerCtor
		}
	}

	if ctor == nil {
		return nil, errors.New("worker has no type")
	}

	return ctor(), nil
}

func getWorkerType(spec map[string]interface{}, types []string) (string, error) {
	containedKey := ""

	for _, key := range types {
		if _, ok := spec[key]; ok {
			if len(containedKey) > 0 {
				return "", fmt.Errorf(
					"worker spec has two types: %s and %s", containedKey, key)
			}
		}
	}

	return containedKey, nil
}

func handleErr(err error) {
	if err != nil {
		log.Printf("job failed: %s", err.Error())
	} else {
		log.Print("executed job")
	}
}
