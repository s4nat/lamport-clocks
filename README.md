**Lamport and Vector Clock - readMe**

This project simulates a client-server architecture where several clients are registered to a server. Periodically, each client sends messages to the server. Upon receiving a message, the server may either forward the message to all other registered clients or drop the message altogether. The project implements both Lamport's logical clock and Vector clocks to order and manage messages.

**Getting Started**

**Running the simulation**

![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.001.png) Build the Go program and run it

cd cmd/logical\_clock![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.002.png)

go build && ./logical\_clock

cd cmd/vector\_clock![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.003.png)

go build && ./vector\_clock

Note: For Windows platforms, run the executable file (e.g.,  logical\_clock.exe ).![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.004.png)

**Overview of the Code**

**Data Structures**

- server : Represents the main server which listens to client messages and might broadcast them to other clients.
- client : Represents a client which sends messages to the server and listens for broadcasted messages.
  - message : Represents a message with its metadata including sender, receiver, actual content, and a logical timestamp.![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.005.png)

**Key Functions**

- NewServer() : Initializes a new server.
  - NewClient() : Initializes a new client.
    - clientSend() : Handles the client-side logic when sending a message. It increments the logical timestamp and sends the message to the server.![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.006.png)
      - clientListen() : Handles the client-side logic when listening to messages. It updates its local logical timestamp based on received messages and detects causality violations using vector clocks.![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.007.png)
    - listenAndBroadcast() : Handles the server-side logic. It listens to incoming messages from clients, updates its own logical timestamp, and decides whether to broadcast the message to other clients.![](Aspose.Words.2e880312-3a81-42c7-9581-363ad111c21a.008.png)

**Logical Clocks**

The program uses two kinds of logical clocks:

1. **Lamport's Logical Clock**: It's a simple counter that ensures a total order of events in a distributed system.
1. **Vector Clock**: A more complex structure that can capture causality between events. It uses an integer array where each index represents a client.

**Understanding the Output**

When you run the simulation, you'll observe a series of messages indicating:

1. Which client has sent a message and its content.
1. Which messages the server has received and whether it broadcasts them.
1. Which messages each client receives.

At the end of the simulation, the program will print a list of messages in the order determined by the logical clocks
Lamport and V ector Clock - readMe 2
