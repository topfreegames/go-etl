package worker

import "errors"

// Schedule represents hh:mm in a day
type Schedule struct {
	Hour   int
	Minute int
}

// Validate validates hour and minute
func (s *Schedule) Validate() error {
	if s.Hour < 0 || s.Hour > 23 {
		return errors.New("invalid hour")
	}

	if s.Minute < 0 || s.Minute > 59 {
		return errors.New("invalide minute")
	}

	return nil
}
