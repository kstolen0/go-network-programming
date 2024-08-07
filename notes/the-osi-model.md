# The Open Systems Interconnection Reference Model

Created in the 1970s, the OSI model serves as a framework for the development and communication about protocols. 

Protocols are rules and procedures that determine the format and order of data sent over a network. 

## The OSI Layers

| Layer 7 | Application |
| :---- | :----: |
| Layer 6 | Presentation |
| Layer 5 | Session |
| Layer 4 | Transport | 
| Layer 3 | Network | 
| Layer 2 | Datalink | 
| Layer 1 | Physical | 

### Layer 7 - The Application layer

Network applications and libraries most often interact with the application layer.
examples of these applications are: Web browsers, Skype, Torrent clients

### Layer 6 - The Presentation layer

Prepares data for the network layer when that data is moving down the stack and presents data to the application layer when the data moves up the stack.
examples of layer 6 functions are: encryption/decryption, data encoding

### Layer 5 - The Session layer

Manages the connection life cycle between nodes on a network. Responsible for establishing the connection, managing time-outs, coordinating the mode of operation, and terminating the connection.

### Layer 4 - Transport layer

Controls nd coordinates the transfer of data between two nodes while maintaining the reliability of the transfer, such as correcting errors, controlling the speed of data transfer, chunking/segmenting data, retransmitting missing data, and acknowledge received data.

### Layer 3 - Network layer

Responsible for transmitting data between nodes. Allows you to send data to a network address without having a direct point to point connection to the remote node. 

### Layer 2 - Data Link layer

Handles data transfers between two directly connected nodes. Protocols on this layer attempt to correct errors on the physical layer.

### Layer 1 - Physical layer

Converts bits from the network stack to electrical, optic, or radio signals and from the physical medium  back to the data layer. This layer controls the bit rate. 

