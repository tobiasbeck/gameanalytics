package gameanalytics

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
	"tobiasbeck/gameanalytics/events"

	"github.com/google/uuid"
)

var sessions = map[string]*Session{}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Event interface {
	json.Marshaler
}

// Client is the api client for gameanalytics. Please instanciate with NewClient
type Client struct {
	key        string
	secret     string
	sessions   map[string]*Session
	batchLock  sync.Mutex
	eventBatch []Event
	client     http.Client
	platform   Platform
}

type NewClientOption = func(*Client)

func WithPlatform(p Platform) NewClientOption {
	return func(c *Client) {
		c.platform = p
	}
}

// NewClient returns a new client instance for gameanalytics
func NewClient(key string, secret string, opts ...NewClientOption) *Client {
	client := &Client{
		key:        key,
		secret:     secret,
		sessions:   map[string]*Session{},
		eventBatch: []Event{},
		batchLock:  sync.Mutex{},
		client:     http.Client{},
		platform: Platform{
			platform:  runtime.GOOS,
			osVersion: "unknown",
		},
	}

	for _, opt := range opts {
		opt(client)
	}
	go handleEvents(client)
	return client
}

func handleEvents(a *Client) {
	for {
		<-time.After(20 * time.Second)
		if len(a.eventBatch) == 0 {
			continue
		}
		a.batchLock.Lock()
		events := a.eventBatch
		a.eventBatch = []Event{}
		a.batchLock.Unlock()
		body, err := json.Marshal(events)
		if err != nil {
			fmt.Printf("client err events: %s", err)
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		_, err = a.post(ctx, "/v2/"+a.key+"/events", string(body))
		if err != nil {
			fmt.Printf("client err events: %s", err)
		}
		cancel()
	}
}

// StartSession opens a session for a user.
// If a session for given user is already open existing session is returned instead and no new session will be opened
func (a *Client) StartSession(ctx context.Context, userId string, platform ...Platform) (*Session, error) {
	if session, ok := a.sessions[userId]; ok {
		return session, nil
	}
	var p Platform = a.platform
	if len(platform) >= 1 {
		p = platform[0]
	}
	data, err := a.sendInit(ctx, userId, p)
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	sess := &Session{
		ID:         id.String(),
		userId:     userId,
		configs:    data.Configs,
		configHash: data.ConfigHash,
		opened:     time.Now(),
		platform:   &p,
	}
	a.sessions[userId] = sess
	ss := events.NewSessionStart(userId, sess.ID, sess.opened)
	a.SendEvent(ss)
	return a.sessions[userId], nil
}

// EndSession ends the session for a user
func (a *Client) EndSession(userId string) error {
	session := a.Session(userId)
	if session == nil {
		return errors.New("Session does not exist")
	}
	e := events.NewSessionEnd(session.userId, session.ID, session.opened)
	a.SendEvent(e)
	return nil
}

// Session returns a users session or nil if no session is open
func (a *Client) Session(userId string) *Session {
	if session, ok := a.sessions[userId]; ok {
		return session
	}
	return nil
}

type initResponse struct {
	Configs    []SessionConfig `json:"configs"`
	ConfigHash string          `json:"configs_hash"`
}

func (a *Client) sendInit(ctx context.Context, userId string, p Platform) (*initResponse, error) {
	bodyData := map[string]interface{}{
		"user_id":     userId,
		"platform":    p.platform,
		"os_version":  p.osVersion,
		"sdk_version": "rest api v2",
		"random_salt": randStringBytes(8),
	}
	if p.build != "" {
		bodyData["build"] = p.build
	}

	body, _ := json.Marshal(bodyData)
	resp, err := a.post(ctx, "remote_configs/v1/init?game_key="+a.key, string(body))
	if err != nil {
		return nil, err
	}
	data := initResponse{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// SendEvent add a event to the queue to send to gameanalytics
func (a *Client) SendEvent(event Event) {
	a.batchLock.Lock()
	a.eventBatch = append(a.eventBatch, event)
	a.batchLock.Unlock()
}

func (a *Client) post(ctx context.Context, urlPath string, body string) ([]byte, error) {
	// fmt.Printf("PATH: %s\n", HOST+urlPath)
	req, err := http.NewRequest("POST", HOST+urlPath, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header = headers(body, a.secret)
	// fmt.Printf("REQH: %#v\n", req.Header)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
