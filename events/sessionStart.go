package events

import (
	"encoding/json"
	"time"
)

type SessionStart struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
}

func (u SessionStart) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "user",
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
	}
	return json.Marshal(data)
}

func NewSessionStart(userId string, sessionId string, created time.Time, platform ...Platform) SessionStart {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return SessionStart{
		UserId:      userId,
		UserSession: sessionId,
		Created:     created,
		Platform:    p,
	}
}
