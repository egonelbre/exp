package after

import (
	"io"

	"github.com/egonelbre/sync2"
)

type Channel struct {
	mu      sync2.Mutex
	clients map[string]io.ReadWriter
}

func (channel *Channel) Connect(name string, client io.ReadWriter) {
	channel.mu.Lock()
	defer channel.mu.Unlock()

	channel.broadcast(name + " connected")
	channel.clients[name] = client
}

func (channel *Channel) Disconnect(name string) {
	channel.mu.Lock()
	defer channel.mu.Unlock()
	channel.disconnect(name)
}

func (channel *Channel) disconnect(name string) {
	channel.mu.MustOwn()

	delete(channel.clients, name)
	channel.broadcast(name + " disconnected")
}

func (channel *Channel) broadcast(message string) {
	channel.mu.MustOwn()

	for name, client := range channel.clients {
		n, err := client.Write([]byte(message))
		if err != nil || n != len(message) {
			channel.disconnect(name)
		}
	}
}
