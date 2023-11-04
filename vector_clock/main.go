package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUMBER_OF_CLIENTS int = 10

type server struct {
	clientArray   []client
	logicalTS     []int
	serverChannel chan message
	allMessages   []message
	mux           sync.Mutex
}

type client struct {
	name          string
	server        *server
	index         int // client's position in the vector clock
	logicalTS     []int
	clientChannel chan message
}

type message struct {
	senderName    string
	messageString string
	logicalTS     []int
}

func NewServer(numClients int) *server {
	clientArray := []client{}
	serverChannel := make(chan message)
	allMessages := []message{}
	logicalTS := make([]int, numClients)
	s := server{clientArray, logicalTS, serverChannel, allMessages, sync.Mutex{}}
	return &s
}

func NewClient(name string, index int, s *server) *client {
	clientChannel := make(chan message)
	logicalTS := make([]int, len(s.logicalTS))
	c := client{name, s, index, logicalTS, clientChannel}
	return &c
}

func (s *server) registerClient(c *client) {
	s.clientArray = append(s.clientArray, *c)
}

func (c *client) clientSend(serverChannel chan message, wg *sync.WaitGroup, serverWg *sync.WaitGroup) {
	defer wg.Done()
	defer serverWg.Done()
	c.logicalTS[c.index]++
	msg := message{
		senderName:    c.name,
		messageString: fmt.Sprintf("Hello from %s", c.name),
		logicalTS:     c.logicalTS,
	}
	fmt.Printf("\n%s sent: %s with VC %v\n", c.name, msg.messageString, msg.logicalTS)
	serverChannel <- msg
	defer wg.Done()
}

func (c *client) clientListen(wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range c.clientChannel {
		c.logicalTS = mergeVC(c.logicalTS, msg.logicalTS)
		c.logicalTS[c.index]++
		fmt.Printf("%s received a message: %s with VC %v\n", c.name, msg.messageString, msg.logicalTS)

		c.server.mux.Lock()
		c.server.allMessages = append(c.server.allMessages, msg)
		c.server.mux.Unlock()
	}
}

func (s *server) listenAndBroadcast(serverWg *sync.WaitGroup) {
	for {
		select {
		case msg, open := <-s.serverChannel:
			if !open {
				return
			}
			s.logicalTS = mergeVC(s.logicalTS, msg.logicalTS)
			fmt.Printf("Server received message from %s: %s with VC %v\n", msg.senderName, msg.messageString, msg.logicalTS)
			if rand.Float32() < 0.5 {
				for _, client := range s.clientArray {
					if client.name != msg.senderName {
						client.clientChannel <- msg
					}
				}
			}
		case <-time.After(5 * time.Second): // Timeout to prevent deadlock
			if serverWg != nil {
				serverWg.Wait()
			}
			close(s.serverChannel)
			return
		}
	}
}

func detectCausalityViolation(s *server) {
	for i := 0; i < len(s.allMessages); i++ {
		for j := i + 1; j < len(s.allMessages); j++ {
			a, b := s.allMessages[i].logicalTS, s.allMessages[j].logicalTS
			if (a[0] > b[0] && a[1] < b[1]) || (a[0] < b[0] && a[1] > b[1]) {
				fmt.Printf("Potential causality violation between messages: %s (VC: %v) and %s (VC: %v)\n",
					s.allMessages[i].messageString, a,
					s.allMessages[j].messageString, b)
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

func mergeVC(local, received []int) []int {
	for i := 0; i < len(local); i++ {
		local[i] = max(local[i], received[i])
	}
	return local
}

func main() {

	var wg sync.WaitGroup
	s := NewServer(NUMBER_OF_CLIENTS)

	// Registering clients
	for i := 1; i <= NUMBER_OF_CLIENTS; i++ {
		c := NewClient(fmt.Sprintf("Client %d", i), i-1, s)
		s.registerClient(c)
	}

	serverWg := &sync.WaitGroup{}
	serverWg.Add(NUMBER_OF_CLIENTS)
	// Start the server's listener goroutine
	go s.listenAndBroadcast(serverWg)

	// Go routine for each client to send messages periodically and listen to broadcasted messages
	for _, c := range s.clientArray {
		wg.Add(2)

		go func(c client) {
			defer wg.Done()
			for i := 0; i < 5; i++ { // Let's assume each client will send 5 messages in total
				go c.clientSend(s.serverChannel, &wg, serverWg)
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Random sleep before sending next message
			}
		}(c)

		go c.clientListen(&wg)
	}

	wg.Wait()

	// Close the serverChannel after all clients have sent their messages
	close(s.serverChannel)

	// Detect causality violations based on vector clocks
	detectCausalityViolation(s)
}
