package main

import (
	"github.com/Henrod/go-etl/processors"
	"github.com/dailyburn/ratchet"
)

type etl string

func (e etl) Extract() ratchet.DataProcessor {
	return processors.NewHTTPRequestor("GET", "http://localhost:8080")
}

func (e etl) Transform() ratchet.DataProcessor {
	return &Transformer{}
}

func (e etl) Load() ratchet.DataProcessor {
	return &Loader{}
}

// ETL is the exported symbol of this plugin
var ETL etl
