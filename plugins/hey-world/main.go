package main

import "fmt"

type etl string

func (e etl) Extract() {
	fmt.Println("Hey World")
}

// ETL is the exported symbol of this plugin
var ETL etl
