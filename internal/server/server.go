package server

import "time"

const (
	DefaultVersion = "0.0.1-development"
)

type Server struct {
	startTime time.Time
	version   string
}

func NewServer() *Server {
	return &Server{
		startTime: time.Now(),
		version:   DefaultVersion,
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
