package example

import (
	"io"
	"sync"
)

type ChannelBefore struct {
	mu      sync.Mutex
	clients map[string]io.ReadWriter
}

func (channel *ChannelBefore) Connect(name string, client io.ReadWriter) {
	channel.mu.Lock()
	defer channel.mu.Unlock()

	channel.broadcast(name + " connected")
	channel.clients[name] = client
}

func (channel *ChannelBefore) Disconnect(name string) {
	channel.mu.Lock()
	defer channel.mu.Unlock()
	channel.disconnect(name)
}

func (channel *ChannelBefore) disconnect(name string) {
	// channel.mu must be held
	delete(channel.clients, name)
	channel.broadcast(name + " disconnected")
}

func (channel *ChannelBefore) broadcast(message string) {
	// channel.mu must be held
	for name, client := range channel.clients {
		n, err := client.Write([]byte(message))
		if err != nil || n != len(message) {
			channel.disconnect(name)
		}
	}
}
