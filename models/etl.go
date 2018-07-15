package models

import "github.com/dailyburn/ratchet"

// ETL defines an ETL
type ETL interface {
	Extract() ratchet.DataProcessor
}
