package uptime

import (
	"time"
)

type Service struct {
	start time.Time
}

func New() *Service {
	return &Service{
		start: time.Now(),
	}
}

func (s *Service) GetDuration() time.Duration {
	return time.Since(s.start)
}
