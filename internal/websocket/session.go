package websocket

import "sync"

type Session struct {
	ID      string
	votes   map[string]int
	Clients map[*Client]bool
	mu      sync.Mutex
}

var sessions = make(map[string]*Session)

var sessionsMu sync.Mutex

func CreateSession(id string) *Session {
	session := Session{
		ID:      id,
		votes:   make(map[string]int),
		Clients: make(map[*Client]bool),
	}
	sessionsMu.Lock()
	sessions[id] = &session
	sessionsMu.Unlock()
	return &session
}

func GetSession(id string) *Session {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()
	return sessions[id]
}

func (s *Session) AddClient(client *Client) {
	s.mu.Lock()
	s.Clients[client] = true
	s.mu.Unlock()
}

func (s *Session) RemoveClient(client *Client) {
	s.mu.Lock()
	delete(s.Clients, client)
	s.mu.Unlock()
}

func (s *Session) Broadcast(message Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.Clients {
		client.send <- message
	}
}
