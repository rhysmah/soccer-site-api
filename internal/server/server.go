package interal

import "time"

type Server struct {
	startTime time.Time
	version   string
}

const (
	version = "0.0.1-development"
)

func NewServer() *Server {
	return &Server{
		startTime: time.Now(),
		version:   version,
	}
}

func (s *Server) GetStartTime() time.Time {
	return s.startTime
}

func (s *Server) GetUptime() time.Duration {
	return time.Since(s.startTime)
}

func (s *Server) GetVersion() string {
	return s.version
}
