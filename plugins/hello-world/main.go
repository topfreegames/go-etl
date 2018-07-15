package main

import (
	"github.com/dailyburn/ratchet"
)

type etl string

func (e etl) Extract() ratchet.DataProcessor {
	return &Extractor{}
}

func (e etl) Transform() ratchet.DataProcessor {
	return &Transformer{}
}

func (e etl) Load() ratchet.DataProcessor {
	return &Loader{}
}

// ETL is the exported symbol of this plugin
var ETL etl
