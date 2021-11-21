package events

import (
	"encoding/json"
	"strings"
	"time"
)

type Design struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	ProgressionStatus PStatus
	IDParts           []string
	Value             float64
}

func (e Design) eventId() string {
	parts := e.IDParts
	if len(parts) > 5 {
		parts = parts[0:5]
	}
	return strings.Join(parts, ":")
}

func (e Design) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "design",
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
		"event_id": e.eventId(),
		"value":    e.Value,
	}
	filterOptionalFields(data, "value")
	return json.Marshal(data)
}

func NewDesign(userId string, sessionId string, platform ...Platform) Design {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return Design{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Platform:    p,
	}
}
