# User Datagram Protocol (UDP)

UDP is a much simpler protocol compared to TCP however it is also much more unreliable. 
UDP is so simple that [its rfc](https://datatracker.ietf.org/doc/html/rfc768) is only about three pages long.
UDP does not provide any session support, or confirm if a destination is accessible. Recipients fo not automatically acknowledge packets.  
It does not manage congestion, control data flow, or retransmit packets.  
UDP does not guarantee that the destination will receive packets in the order they're sent.

The lack of these features is what makes UDP fast.

