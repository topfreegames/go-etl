package app

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/topfreegames/go-etl/reader"
	"github.com/topfreegames/go-etl/worker"
)

// App starts all workers
type App struct {
	workers []worker.Worker
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
	log.Print("reading config")
	workers, err := reader.Read()
	if err != nil {
		return err
	}

	log.Print("configuring worker")
	workers = workers.Configure()

	log.Print("success on worker")
	a.workers = workers
	return nil
}

// Start starts all workers
func (a *App) Start() {
	done := make(chan struct{})

	wg.Add(len(a.workers))
	for _, w := range a.workers {
		go a.startWorker(w, done)
	}

	a.wait(done)
}

func (a *App) startWorker(worker worker.Worker, done chan struct{}) {
	worker.Start(done)
	wg.Done()
}

func (a *App) wait(done chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	a.terminate(done)

	wg.Wait()
}

func (a *App) terminate(done chan struct{}) {
	log.Print("terminating app...")
	log.Print("waiting for graceful shutdown")
	close(done)
}
