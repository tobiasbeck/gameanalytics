package events

import (
	"encoding/json"
	"strings"
	"time"
)

type Flow = string

const Sink Flow = "Sink"
const Source Flow = "Source"

type Ressource struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	FlowType Flow
	Currency string
	ItemType string
	ItemId   string
	Amount   int
}

func (e Ressource) eventId() string {
	return strings.Join([]string{e.FlowType, e.Currency, e.ItemType, e.ItemId}, ":")
}

func (e Ressource) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "resource",
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
		"amount":   e.Amount,
	}
	return json.Marshal(data)
}

func NewRessource(userId string, sessionId string, platform ...Platform) Ressource {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return Ressource{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Platform:    p,
	}
}
