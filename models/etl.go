package models

// ETL defines an ETL
type ETL interface {
	Extract() DataProcessor
	Transform() DataProcessor
	Load() DataProcessor
}
