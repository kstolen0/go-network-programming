# Unix Domain Sockets

Unix domain sockets are a form of inter-process communication (IPC) which uses the filesystem to determine a packet's destination address. 

Socket addressing allows individual services on the same node to listen for incomining traffic. 

Unix domain sockets apply the socket addressing principle to the file-system.  

Each domain socket has an associated file on the file system, which corresponds to a network socket's IP address and port number.

A service can listen to reads and writes of a given file.  

This protocol allows services to bypass the OS's network stack. 

Forgoeing Unix Domain Sockets for network sockets when communicating with local services can sacrifice multiple security advantages and performance gains.

Unix Domain Sockets have many benefits when communicating with local services but may not be appropriate if a service has to move across multiple nodes or has portability requirements.


