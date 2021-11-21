package events

import (
	"encoding/json"
	"time"
)

type Error struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	Severity string
	Message  string
}

func (e Error) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "error",
		"device":       "unknown",
		"v":            2,
		"user_id":      e.UserId,
		"client_ts":    e.Created.Unix(),
		"sdk_version":  "rest api v2",
		"os_version":   "linux 1",
		"manufacturer": "unknown",
		"platform":     "linux",
		"session_id":   e.UserSession,
		"session_num":  1,
		// PROPERTIES
		"severity": e.Severity,
		"message":  e.Message,
	}
	return json.Marshal(data)
}

func NewError(userId string, sessionId string, platform ...Platform) Error {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return Error{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Platform:    p,
	}
}
