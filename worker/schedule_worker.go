package worker

import (
	"log"
	"time"

	"github.com/topfreegames/go-etl/models"
)

// ScheduleWorker executes etl every day at an schedule
type ScheduleWorker struct {
	Schedule *Schedule
	Job      *models.Job
	timer    *time.Timer
}

// NewScheduleWorker returns an ScheduleWorker
func NewScheduleWorker() Worker {
	return &ScheduleWorker{
		timer: time.NewTimer(1 * time.Second),
	}
}

// GetJob returns the job
func (s *ScheduleWorker) GetJob() *models.Job {
	return s.Job
}

// Validate validates worker configuration
func (s *ScheduleWorker) Validate() error {
	return s.Schedule.Validate()
}

// Start runs the worker every period
func (s *ScheduleWorker) Start(done chan struct{}) {
	for {
		select {
		case <-s.tick():
			log.Printf("executing job %s", s.Job.Name)
			err := s.Job.Execute()
			handleErr(err)
		case <-done:
			log.Print("terminating worker")
			return
		}
	}
}

func (s *ScheduleWorker) tick() <-chan time.Time {
	now := time.Now().UTC()
	nextSchedule := time.Date(
		now.Year(), now.Month(), now.Day(),
		s.Schedule.Hour, s.Schedule.Minute, 0, 0,
		time.UTC)

	if now.After(nextSchedule) {
		nextSchedule = nextSchedule.AddDate(0, 0, 1)
	}

	duration := nextSchedule.Sub(now)
	s.timer.Reset(duration)

	return s.timer.C
}
