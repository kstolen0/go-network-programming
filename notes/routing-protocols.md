# Routing Protocols

### Terms

. | .
--- | ---
Autonomous System | An organisation that manages one or more networks (e.g. And ISP)
Autonomous System Number (ASN) | A unique identifier assigned to an Autonomous System (See rfc 1930) used to broadcast network information to other Autonomous Systems using an external routing protocol
DNS | Domain Name System

## Border Gateway Protocol (BGP)

Allows an ASNassigned ISP to exchange routing information. 
BGP relies on trust between ISPs. 
Misconfiguring BGP (route leaks) can result in large public network outages.

## Name and Address Resolution

DNS is a method of matching IP addresses to human readable domain names. 
All domains are children of top-level domains, e.g. .com, .net, .org, etc.

When entering a domain name in a browser, the client will consult the relevant domain name resolver. 
1. The resolver will start by asking one of the IANA-maintained root name servers for the IP address of the domain.
2. The root name server will examine the top level domain and return the resolver for that name server.
3. The resolver will then examine ask that name server for the the IP address which will examine the domain portion and direct the resolver to ask the domain's name server. 
4. Finally the resolver will ask the domain name server for the IP address corresponding to the domain name. 
5. The browser will establish a connection to this IP address and retrieve the web page.

### Domain Name Resource Records

Domain name servers maintain resource records for the domains they serve.
These records contain domain-specific information which are used to satisfy domain name queries.

Some common records are:

#### Address (A)

The most common record. This resolves to one or more IPv4 addresses.
When the browser asks to resolve the IP address for a domain name, the resolver ultimately asks for the A resource record.

The AAAA record is the IPv6 equivalent of the A record.

#### Start of Authority (SOA)

Contains the authorative and administrative details about the domain.
All domains must have an SOA record.

The SOA record includes:
1. the primary name server 
2. the administrators email address
3. fields used by secondary name servers

#### Name Server (NS)

Returns the authoritative name servers for the domain name.
Authortative name servers are the name servers able to provide answers for the domain name.

#### Canonical Name (CNAME)

Points one domain at another. 
These can make administration easier. 

#### Mail Exchange (MX)

Specifies the mail server hostnames that should be contacted when sending email to recipients at the domain.

#### Pointer (PTR)

Allows you to perform a reverse lookup by accepting an IP address and returning its corresponding domain name.

#### Text (TXT)

Allows the domain owner to return arbitrary text. 
These records can contain values that prove ownership of the domain, values remote servers can use to authorize email, and anything else.



