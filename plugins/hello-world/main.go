package main

import (
	"github.com/topfreegames/go-etl/models"
)

type etl string

func (e etl) Extract() models.DataProcessor {
	return &Extractor{}
}

func (e etl) Transform() models.DataProcessor {
	return &Transformer{}
}

func (e etl) Load() models.DataProcessor {
	return &Loader{}
}

// ETL is the exported symbol of this plugin
var ETL etl
