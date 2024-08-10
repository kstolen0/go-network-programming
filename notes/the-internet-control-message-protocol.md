# The Internet Control Message Protocol (ICMP)

The IP protocol relies on ICMP to provide feedback about the local nework.

ICMP can inform clients about: 
* network problems
* unreachable nodes or networks
* local network configuration
* proper traffic routes
* network timeouts

IPv4 and IPv6 have their own ICMP implementations (ICMPv4 and ICMPv6). 

Network events often result in ICMP response messages. 

Routers use ICMP to help inform clients of better routes to their destination node. 
If data is sent to a router that isn't the appropriate or best router to handle traffic for a destination it may reply with a ICMP redirect message after forwarding the data to the correct router.

You can dermine if a node is online and reachable by using an ICMP echo request (aka `ping`). 

If the node is reachable and receives the ping, it will reply with an echo reply message (aka `pong`).

