package events

import (
	"encoding/json"
	"strings"
	"time"
)

type PStatus = string

const StatusStart PStatus = "Start"
const StatusFail PStatus = "Fail"
const StatusComplete PStatus = "Complete"

type Progression struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	ProgressionStatus PStatus
	Progression1      string
	Progression2      string
	Progression3      string
	AttemptNum        string
	Score             string
}

func (e Progression) eventId() string {
	event := make([]string, 0, 4)
	event = append(event, e.ProgressionStatus, e.Progression1)
	if e.Progression2 != "" {
		event = append(event, e.Progression2)
	}
	if e.Progression3 != "" {
		event = append(event, e.Progression3)
	}
	return strings.Join(event, ":")
}

func (e Progression) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "progression",
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
		"event_id":    e.eventId(),
		"attempt_num": e.AttemptNum,
		"score":       e.Score,
	}
	filterOptionalFields(data, "attempt_num", "score")
	return json.Marshal(data)
}

func NewProgression(userId string, sessionId string, platform ...Platform) Progression {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return Progression{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Platform:    p,
	}
}
