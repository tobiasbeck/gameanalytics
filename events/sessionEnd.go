package events

import (
	"encoding/json"
	"time"
)

type SessionEnd struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	Length int
}

func (u SessionEnd) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "session_end",
		"device":       "unknown",
		"v":            2,
		"user_id":      u.UserId,
		"client_ts":    u.Created.Unix(),
		"sdk_version":  "rest api v2",
		"os_version":   "linux 1",
		"manufacturer": "unknown",
		"platform":     "linux",
		"session_id":   u.UserSession,
		"session_num":  1,
		// PROPERTIES
		"length": u.Length,
	}
	return json.Marshal(data)
}

func NewSessionEnd(userId string, sessionId string, created time.Time, platform ...Platform) SessionEnd {
	now := time.Now()
	duration := now.Sub(created)
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return SessionEnd{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Length:      int(duration.Seconds()),
		Platform:    p,
	}
}
