package uptime

import (
	"time"
)

// Service uptime definition
type Service struct {
	start time.Time
}

// New uptime service
func New() *Service {
	return &Service{
		start: time.Now(),
	}
}

func (s *Service) GetDuration() time.Duration {
	return time.Since(s.start)
}
