package interal

import "time"

type Server struct {
	startTime time.Time
	version   string
}

func NewServer() *Server {
	return &Server{
		startTime: time.Now(),
		version:   "0.0.1-development",
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
