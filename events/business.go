package events

import (
	"encoding/json"
	"time"
)

type Business struct {
	UserId      string
	UserSession string
	Created     time.Time
	Platform    Platform
	// PROPERTIES
	EventId        string
	Amount         int
	Currency       string
	TransactionNum string
	CartType       string
	ReceiptInfo    string
}

func (e Business) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"category":     "business",
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
		"event_id":        e.EventId,
		"amount":          e.Amount,
		"currency":        e.Currency,
		"transaction_num": e.TransactionNum,
		"cart_type":       e.CartType,
		"receipt_info":    e.ReceiptInfo,
	}
	filterOptionalFields(data, "cart_type", "receipt_info")
	return json.Marshal(data)
}

func NewBusiness(userId string, sessionId string, platform ...Platform) Business {
	var p Platform = nil
	if len(platform) > 0 {
		p = platform[0]
	}
	return Business{
		UserId:      userId,
		UserSession: sessionId,
		Created:     time.Now(),
		Platform:    p,
	}
}
