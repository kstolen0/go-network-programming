# The TCP/IP Model

Transmission Control Protol over Internet Protocol consists of four named layers: 
1. application
2. transport
3. internet 
4. link

This model simplifies the OSI model's application, presentation, and session layers into a single application layer.
Simlarly, the Link layer of the TCP/IP model spans both the data link and physical layers of the OSI model.
The transport and internet layers share a one-to-one relationship witht the transport and network layers of the OSI model, respectively.


## The Application Layer

Interacts directly with software applications. Most web software uses protocols defined for this layer.
Examples of these protocols are: 
* HTTP
* FTP
* SMTP
* DHCP
* DNS

## The Transport Layer

Handles the transfer of data between two nodes. These protocols help to ensure data integrity. 
Examples of these protocols are:
* TCP 
* UDP (User Datagram Protocol)

## The Internet Layer

Responsible for routing packets of data form the upper layers between the origin node and the destination node. 
Examples of Internet layer protocols are:
* IPv4 (Internet Protocol version 4)
* IPv6 (Internet Protocol version 6)
* BGP (Border Gateway Protocol)
* ICMP (Internet Control Message Protocol)
* IGMP (Internet Group Management Protocol)
* IPsec (Internet Protocol Security Suite)

## The Link Layer

The interface between the core TCP/IP protocols and the physical media. 
This layer uses the Address Resolution Protocol (ARP) to translate a node's IP address to the MAC address of its network interface.

See also: 

* RFC 1122 for the lower three layers of TCP/IP https://datatracker.ietf.org/doc/html/rfc1122/
* RFC 1123 for the application layer oF TCP/IP https://datatracker.ietf.org/doc/html/rfc1123/
