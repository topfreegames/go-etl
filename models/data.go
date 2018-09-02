package models

import "encoding/json"

// Data is what is passes by pipeline stages
type Data []byte

// NewData returns default data
func NewData() Data {
	return Data("GO")
}

// Unmarshal calls json.Unmarshal on Data into value
func (d Data) Unmarshal(value interface{}) error {
	return json.Unmarshal(d, value)
}
