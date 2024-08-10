# Transmition Control Protocol (TCP)

Allows you to reliably stream data between nodes on a network.

TCP overcomes the effects of packet loss and receiving packets out of order.

Packet loss occurs when data fails to reach its destination.

TCP adapts its data transfer rate to make sure it transmits data as fast as possible while keeping dropped packets to a minimum. This is known as flow control. It makes up for the deficiencies of the underlying network media. 
TCPs reliability is generally dependant on its underlying network hardware.

TCP keeps track of received packets and retransmits unacknowledged packets.
Recipients can receive packets out of sequence. TCP organises unordered packets and processes them in sequence. 

## Working with TCP Sessions

TCP sessions allow you to deliver a stream of data of any size to a recipient and receive confirmation that the recipient received the data. Streaming allows the sender to receive feedback from the recipient while the transfer is taking place so errors can be corrected in real-time.

### Establishing a session with a TCP handshake

TCP connections use a three-way handshake between the client and server. The handshake creates an established TCP session over which to exchange data. The handshake process consists of three messages:

```mermaid
sequenceDiagram
Client ->> Server: SYN
Server ->> Client: SYN / ACK
Client ->> Server: SYN
```


