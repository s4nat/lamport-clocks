package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const NUMBER_OF_CLIENTS int = 10

type server struct {
	clientArray   []*client
	logicalTS     int
	serverChannel chan message
	allMessages   []message
	mux           sync.Mutex
}

type client struct {
	name          string
	server        *server
	logicalTS     int
	clientChannel chan message
}

type message struct {
	senderName    string
	messageString string
	logicalTS     int
}

func NewServer() *server {
	clientArray := []*client{}
	serverChannel := make(chan message)
	allMessages := []message{}
	s := server{clientArray, 0, serverChannel, allMessages, sync.Mutex{}}
	return &s
}

func NewClient(name string, s *server) *client {
	clientChannel := make(chan message)
	c := client{name, s, 0, clientChannel}
	return &c
}

func (s *server) registerClient(c *client) {
	s.clientArray = append(s.clientArray, c)
}

func (c *client) clientSend(serverChannel chan message, wg *sync.WaitGroup) {
	c.logicalTS++
	msg := message{
		senderName:    c.name,
		messageString: fmt.Sprintf("Hello from %s", c.name),
		logicalTS:     c.logicalTS,
	}
	fmt.Printf("\n%s sent: %s with TS %d\n", c.name, msg.messageString, msg.logicalTS)
	serverChannel <- msg
	defer wg.Done()
}

func (c *client) clientListen(wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range c.clientChannel {
		c.logicalTS = max(c.logicalTS, msg.logicalTS) + 1
		fmt.Printf("%s received a message: %s with TS %d\n", c.name, msg.messageString, msg.logicalTS)

		// Add the message to the global slice
		c.server.mux.Lock()
		c.server.allMessages = append(c.server.allMessages, msg)
		c.server.mux.Unlock()
	}
}

func (s *server) listenAndBroadcast() {
	for msg := range s.serverChannel {
		s.logicalTS = max(s.logicalTS, msg.logicalTS) + 1
		fmt.Printf("Server received message from %s: %s with TS %d\n", msg.senderName, msg.messageString, msg.logicalTS)
		if rand.Float32() < 0.5 {
			for _, client := range s.clientArray {
				if client.name != msg.senderName {
					client.clientChannel <- msg
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTotalOrder(s *server) {
	// Sort the slice of messages based on their logical timestamps
	sort.Slice(s.allMessages, func(i, j int) bool {
		return s.allMessages[i].logicalTS < s.allMessages[j].logicalTS
	})

	// Print the sorted messages
	fmt.Println("\n\nTOTAL ORDER:")
	for _, msg := range s.allMessages {
		fmt.Printf("Sender: %s, Message: %s, Timestamp: %d\n", msg.senderName, msg.messageString, msg.logicalTS)
	}
}

func main() {

	var wg sync.WaitGroup
	s := NewServer()

	// Registering clients
	for i := 1; i <= NUMBER_OF_CLIENTS; i++ {
		c := NewClient(fmt.Sprintf("Client %d", i), s)
		s.registerClient(c)
	}

	// Start the server's listener goroutine
	go s.listenAndBroadcast()

	// Go routine for each client to send messages periodically and listen to broadcasted messages
	for _, c := range s.clientArray {
		wg.Add(2)

		go func(c *client) {
			defer wg.Done()
			for i := 0; i < 5; i++ { // Let's assume each client will send 5 messages in total
				c.clientSend(s.serverChannel, &wg)
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Random sleep before sending next message
			}
		}(c)

		go c.clientListen(&wg)
	}

	wg.Wait()

	// Close the serverChannel after all clients have sent their messages
	close(s.serverChannel)

	// Print the total order of messages
	printTotalOrder(s)
}
