// internal/interfaces/websockets/leaderboard_socket.go
package websockets

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin returns true if the request Origin header is acceptable. For development, we'll accept any.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// LeaderboardHub maintains the set of active clients and broadcasts messages to the clients.
type LeaderboardHub struct {
	// Registered clients.
	Clients map[*LeaderboardClient]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *LeaderboardClient

	// Unregister requests from clients.
	Unregister chan *LeaderboardClient

	mu sync.Mutex // To protect access to the clients map
}

// LeaderboardClient is a middleman between the websocket connection and the hub.
type LeaderboardClient struct {
	Hub *LeaderboardHub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// NewLeaderboardHub creates a new Hub to manage leaderboard updates.
func NewLeaderboardHub() *LeaderboardHub {
	return &LeaderboardHub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *LeaderboardClient),
		Unregister: make(chan *LeaderboardClient),
		Clients:    make(map[*LeaderboardClient]bool),
	}
}

// Run starts the main Hub operations for handling client connections and broadcasting messages.
func (hub *LeaderboardHub) Run() {
	for {
		select {
		case client := <-hub.Register:
			hub.Clients[client] = true
		case client := <-hub.Unregister:
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.Send)
			}
		case message := <-hub.Broadcast:
			hub.mu.Lock()
			for client := range hub.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(hub.Clients, client)
				}
			}
			hub.mu.Unlock()
		}
	}
}

// ServeWs handles new WebSocket requests from clients, upgrading them from regular HTTP connections.
func ServeWs(hub *LeaderboardHub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &LeaderboardClient{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Start a goroutine for handling client messages.
	go client.writePump()
}

// writePump pumps messages from the hub to the websocket connection.
func (c *LeaderboardClient) writePump() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Close the writer to send the message.
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}