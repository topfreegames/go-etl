package main

import (
	"github.com/topfreegames/go-etl/models"
	"github.com/topfreegames/go-etl/processors"
)

type etl string

func (e etl) Extract() models.DataProcessor {
	return &processors.Logger{}
}

func (e etl) Transform() models.DataProcessor {
	return &processors.Logger{}
}

func (e etl) Load() models.DataProcessor {
	return &processors.Logger{}
}

// ETL is the exported symbol of this plugin
var ETL etl
