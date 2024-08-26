package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
	Close() error
}

/**
 * Transport is anything that can handle communications b/w nodes in the network.
 * This can be in the form (TCP, UDP, Websockets, ...)
 */

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
