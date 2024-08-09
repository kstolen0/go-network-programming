# IPv6 Addressing

IPv6 addresses are 128 bit numbers arranged in eight colon-delimited groups of 16 bits (aka hextets) and are typically represented in lowercase hexadecimal values. 

There are a lot of IPv6 addresses.

## IPv6 Network and Host Addresses

As with IPv4 addresses, IPv6 addresses also have a Network ID and Host ID.
The Host ID is commonly known as the interface ID.
Both the network and host IDs are 64 bits.
The first 48 bits of the network address are known as the Global Routing Prefix (GRP).
The last 16 bits are called the subnet ID.

The GRP is used for globally subdividing the IPv6 address space and routing traffic between these groups.
The subnet ID is used to further subdivide each GRP group into site-specific networks.

### IPv6 Address Categories

#### Unicast Addresses

A unicast address uniquely identifies a node.

#### Multicast Addresses

Represents a group of nodes. Whereas IPv4 broadcast addresses will propagate a message to all addresses on a network, multicast addresses can simultaneously deliver a message to a subset of network addresses, not just all of them.

#### Anycast Addresses

Ipv6 supports multiple nodes using the same network address. Anycast addresses represents a group of nodes listening on the same address, where the nearist node to the sender will receive the message. 
The router determines which node receives the message, which is usually the node with the least latency between the origin and destination. 


### Advantages over IPv4

#### Simplified header format for more efficient routing

The IPv4 header contains mandatory but rarely used fields. IPv6 makes these fields optional.
The Ipv6 header is also extensible. 

#### Stateless address autoconfiguration

Unlike in Ipv4, nodes using IPv6 can automatically configure or derive their IPv6 addeRsses through stateless address autoconfiguration (SLAAC).
This initially had a security flaw, wherin the host portion of the  Ipv6 address contained the nodes NIC's MAC address, which allows anyone to track your online activity. 
A privacy extension was added to SLAAC to randomize the interface ID.


