# lamport-clocks


![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image2.png){width="0.6944444444444444in"
height="0.18055555555555555in"}

> **Lamport and Vector Clock - readMe**
>
> This project simulates a client-server architecture where several
> clients are registered to a server. Periodically, each client sends
> messages to the server. Upon receiving a message, the server may
> either forward the message to all other registered clients or drop the
> message altogether. The project implements both Lamport\'s logical
> clock and Vector clocks to order and manage messages.
>
> **Getting Started**
>
> **Running the simulation**

  ------------------------------------------------------------------------------------------------------------------------------------
  ![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image1.png){width="5.555555555555555e-2in"   Build the Go program and run it
  height="5.555555555555555e-2in"}                                                                 
  ------------------------------------------------------------------------------------------------ -----------------------------------

  ------------------------------------------------------------------------------------------------------------------------------------

+-----------------------------------------------------------------------+
| > cd cmd/logical_clock                                                |
+=======================================================================+
+-----------------------------------------------------------------------+

+-----------------------------------------------------------------------+
| > go build && ./logical_clock                                         |
+=======================================================================+
+-----------------------------------------------------------------------+

+-----------------------------------------------------------------------+
| > cd cmd/vector_clock                                                 |
+=======================================================================+
+-----------------------------------------------------------------------+

+-----------------------------------------------------------------------+
| > go build && ./vector_clock                                          |
+=======================================================================+
+-----------------------------------------------------------------------+

> Note: For Windows platforms, run the executable file (e.g.,
> logical_clock.exe ).
>
> **Overview of the Code**
>
> **Data Structures**
>
> server : Represents the main server which listens to client messages
> and might broadcast them to other clients.

Lamport and Vector Clock - readMe 1

![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image3.png){width="0.6944444444444444in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image4.png){width="0.7638888888888888in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image5.png){width="1.0555555555555556in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image6.png){width="1.0555555555555556in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image7.png){width="1.1111111111111112in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image8.png){width="1.25in"
height="0.18055555555555555in"}![](vertopal_0bf5b266ef2e4bbbbb526e4d0b04e65c/media/image9.png){width="1.6666666666666667in"
height="0.18055555555555555in"}

> client : Represents a client which sends messages to the server and
> listens for broadcasted messages.
>
> message : Represents a message with its metadata including sender,
> receiver, actual content, and a logical timestamp.
>
> **Key Functions**
>
> NewServer() : Initializes a new server.
>
> NewClient() : Initializes a new client.
>
> clientSend() : Handles the client-side logic when sending a message.
> It increments the logical timestamp and sends the message to the
> server.
>
> clientListen() : Handles the client-side logic when listening to
> messages. It updates its local logical timestamp based on received
> messages and detects causality violations using vector clocks.
>
> listenAndBroadcast() : Handles the server-side logic. It listens to
> incoming messages from clients, updates its own logical timestamp, and
> decides whether to broadcast the message to other clients.
>
> **Logical Clocks**
>
> The program uses two kinds of logical clocks:
>
> 1\. **Lamport\'s Logical Clock**: It\'s a simple counter that ensures
> a total order of events in a distributed system.
>
> 2\. **Vector Clock**: A more complex structure that can capture
> causality between events. It uses an integer array where each index
> represents a client.
>
> **Understanding the Output**
>
> When you run the simulation, you\'ll observe a series of messages
> indicating:
>
> 1\. Which client has sent a message and its content.
>
> 2\. Which messages the server has received and whether it broadcasts
> them.
>
> 3\. Which messages each client receives.

Lamport and Vector Clock - readMe 2

> At the end of the simulation, the program will print a list of
> messages in the order determined by the logical clocks

Lamport and Vector Clock - readMe 3
