package main

import (
	"log"

	"github.com/aiden-janey/go_distributed_file_system/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	/**
	 * select{} lets you wait on multiple channel operations. Here
	 * (b/c it's empty) it blocks the current goroutine, so it isn't blocked.
	 */
	select {}
}
