package app

import (
	"sync"

	"github.com/Henrod/go-etl/reader"
	"github.com/Henrod/go-etl/worker"
)

// App starts all workers
type App struct {
	workers []*worker.Worker
}

var wg sync.WaitGroup

// NewApp reads spec and returns app
func NewApp(reader reader.Reader) (*App, error) {
	app := &App{}
	err := app.configure(reader)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) configure(reader reader.Reader) error {
	workers, err := reader.Read()
	if err != nil {
		return err
	}

	err = workers.Configure()
	if err != nil {
		return err
	}

	a.workers = workers
	return nil
}

// Start starts all workers
func (a *App) Start() {
	done := make(chan struct{})
	defer close(done)

	wg.Add(len(a.workers))
	for _, w := range a.workers {
		go a.startWorker(w, done)
	}

	wg.Wait()
}

func (a *App) startWorker(worker *worker.Worker, done chan struct{}) {
	worker.Start(done)
	wg.Done()
}
