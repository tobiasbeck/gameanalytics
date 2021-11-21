package gameanalytics

import "time"

type SessionConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	EndTs int    `json:"end_ts"`
}

// Session is a users session which is currently open
type Session struct {
	ID         string
	userId     string
	configs    []SessionConfig
	configHash string
	opened     time.Time
	platform   *Platform
}

// Config returns a sessions config key if it exists
func (s *Session) Config(key string) *SessionConfig {
	for _, config := range s.configs {
		if config.Key == key {
			return &config
		}
	}
	return nil
}
