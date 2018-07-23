package main

import (
	"github.com/topfreegames/go-etl/processors"
	"github.com/dailyburn/ratchet"
)

type etl string

func (e etl) Extract() ratchet.DataProcessor {
	return processors.NewHTTPRequestor("GET", "http://localhost:8080")
}

func (e etl) Transform() ratchet.DataProcessor {
	return &processors.Logger{}
}

func (e etl) Load() ratchet.DataProcessor {
	return &processors.Null{}
}

// ETL is the exported symbol of this plugin
var ETL etl
