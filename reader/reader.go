package reader

import "github.com/topfreegames/go-etl/worker"

// Reader has the Read method that reads jobs from a source
type Reader interface {
	Read() (worker.Workers, error)
}
