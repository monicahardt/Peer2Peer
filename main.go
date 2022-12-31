package main

import (
	proto "Peer2peer/grpc"
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type Peer struct {
	proto.UnimplementedReceiveServer
	id               int32
	lamport          int32
	amountOfRequests map[int32]int32
	state            State
	clients          map[int32]proto.ReceiveClient
	ctx              context.Context
	port 			int
}

var criticalSection = false;
var port = flag.Int("port", 0, "server port number")

func main() {
	//setting the log file
	f, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	
	// arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	// ownPort := int32(arg1) + 5000

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &Peer{
		id:               int32(*port),
		port: 			*port,
		amountOfRequests: make(map[int32]int32),
		clients:          make(map[int32]proto.ReceiveClient),
		ctx:              ctx,
	}

	go startPeerServer(p)
	go p.connectToOtherPeers()

	for {

	}
}

func startPeerServer(p *Peer) {

	grpcServer := grpc.NewServer()                                           // create a new grpc server
	listen, err := net.Listen("tcp", "localhost:"+strconv.Itoa(p.port)) // creates the listener

	if err != nil {
		log.Fatalln("Could not start listener")
	}

	log.Printf("Started peer at port %v", p.port)
	fmt.Printf("Started peer at port %v", p.port)

	proto.RegisterReceiveServer(grpcServer, p)
	serverError := grpcServer.Serve(listen)

	if serverError != nil {
		log.Printf("Could not register server")
	}

}

func (p *Peer) connectToOtherPeers(){
	//hardcoded to connect to peers with ports 5001, 5002 and 5003
	for i := 0; i < 3; i++ {
		port := int32(5001) + int32(i)

		//if it is the peers own port we just continue
		if port == int32(p.port) {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("Trying to dial: %v\n", port)
		fmt.Printf("Trying to dial: %v\n", port)

		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}

		defer conn.Close()
		c := proto.NewReceiveClient(conn)
		//adding the new connection to another peer (client) to the map of peers/clients
		p.clients[port] = c
		fmt.Printf("Connection made to peer with id: %v\n", port)
	}
	//wait for someone to want to enter
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.Enter()
	}
}


func (p *Peer) Receive(ctx context.Context, req *proto.Request) (*proto.Reply, error) {
	if(p.lamport < req.Lamport) { p.lamport = req.Lamport; }
	p.lamport++;

	if p.state == held || (p.state == wanted && (req.Lamport > p.lamport)) {
		for criticalSection {
			time.Sleep(2*time.Second);
		}
		p.lamport++;
		rep := &proto.Reply{Id: p.id, Lamport: p.lamport};
		return rep, nil;
	} else {
		p.lamport++;
		rep := &proto.Reply{Id: p.id, Lamport: p.lamport};
		return rep, nil;
	}
}

func (p *Peer) Enter() {
	p.state = wanted
	for _, client := range p.clients {
		p.lamport++;
		log.Printf("Peer with id: %v requested access to the critical section with lamporttime: %v", p.id,p.lamport)
		request := &proto.Request{Id: p.id, Lamport: p.lamport}
		rep,_:= client.Receive(p.ctx, request)
		if(rep.Lamport > p.lamport) {p.lamport = rep.Lamport}
		p.lamport++;
		log.Printf("Peer with id: %v received reply to access with lamporttime: %v", p.id,p.lamport)

	}
	//recieved all replies
	p.state = held
	criticalSection = true;
	p.lamport++;
	log.Printf("**** Peer with id: %v entered the critical section with lamporttime: %v ****", p.id,p.lamport)
	time.Sleep(5 * time.Second)
	p.Exit();
}

func (p *Peer) Exit(){
	p.lamport++;
	p.state = released
	criticalSection = false;
	log.Printf("---- Peer with id: %v exited the critical section with lamporttime: %v ----", p.id, p.lamport)
}

// our enum for State
type State string
const (
	released       = "released"
	held           = "held"
	wanted         = "wanted"
	undefinedState = "illegal"
)
