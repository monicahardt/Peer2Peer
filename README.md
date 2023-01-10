# Peer2Peer

To run this code start three different peers on ports 5001, 5002 and 5003
as so

go run . -port 5001
go run . -port 5002
go run . -port 5003

pres enter at a terminal to request access to the critical section
Mutual exclusion could also be done with a channel

We have hardcoded that a peer should connect to three other peers. To add more peers to the program you must change the loop in the method connectToOtherPeers() to connect to more peers on different ports.

As of right now it is hardcoded to connect to peers with ports 5001, 5002 and 5003
