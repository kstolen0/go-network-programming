# The Internet Protocol

A set of rules that dictate the format of data sent over a network (the internet).

IP addresses identify nodes on a network at the internet layer of the [TCP/IP stack](the-tcp-ip-model.md) and are used to facilitate communication between nodes.

There are currently two versions of IP addresses: 
* IPv4
* IPv6

## IPv4 Addressing

A 32 bit number sequence arranged in four octets (8 bit groups) separated by decimal points.
There can be a total of just over four billion unique IPv4 addresses.

## Network and Host IDs

An IPv4 address consists of two components: 
* a network ID
* a host ID

The network ID informs network devices responsible for transfering packets (aka routers) toward their next destination about the next appropriate hop in the transmission. 

Routers accept data from a device, examine the network ID of the destination address, and forwards the data towards that destination. 

Once the data reaches the destination network, the router uses the Host ID to deliver the data to the recipient.

The Network ID always starts with the left-most bits and its size is determined by the size of the network it belongs to.

The remaining bits are designated to the host ID. 

## Subdiving IPv4 Addresses into Subnets

IPv4 allows one to subdivide a network into smaller groups to keep the network secure and easier to manage. These subdivisions are called `Subnets`. 

IP addresses in these subnets share the same Network ID but have unique Host IDs.

Identifying different networks allows you to control the flow of information between networks. 

### Allocating Networks with CIDR

Classless Inter-Domain Routing (CIDR) is used to indicate the number of bits in the network for a given IP address.

### Ports and Socket Addresses

Ports are used to uniquely identify data transmission between nodes for the purposes of multiplexing outgoing application data and demultiplexing incoming data back to an application. 
The combination of an IP address and port is known as a socket address.

Ports are a 16 bit unsigned integer. 
* 0 - 1023 are reserved for well-known ports and are assigned to common services by the Internet Assigned Numbers Authority (IANA)
* 1024 - 49151 are considered as semi-reserved for less common services
* 49152 - 65535 are ephemeral ports meant for client socket addresses


### Network Address Translation

Despite IPv4 having over four billion unique addresses, we've already run out... One method for addressing this shortage of address space is to use Network Address Translation (NAT).

NAT allows multiple nodes to share the same public IPv4 address. This can be achieved via a device which can keep track of incoming and outgoing traffic such as a firewall, load balancer, or a router. 

When applying NAT, the client node's private IP address is not visible or directly accessible to other nodes outside the network segment. 

### Unicasting, Multicasting, and Broadcasting

Sending packets from one IP address to another IP address is known as unicast addressing. 

TCP/IP supports multicasting, i.e. sending a message to a group of nodes.

Broadcasting is the ability to concurrently deliver a message to all IP addresses in a network. This is achieved by sending packets to the broadcast address of a subnet.

### Resolving the MAC address to a physical network connection

The MAC address is only relevant to the local network.

The Address Resolution Protocol (ARP) finds the appropriate MAC address for a given IP address (aka resolving). 
Nodes maintain ARP tables that map an IPv4 address to a MAC address. 
