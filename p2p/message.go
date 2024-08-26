package p2p

import "net"

/**
* Message holds any arbitrary data that's being sent over
* each transport b/w two nodes in the network.
 */
type RPC struct {
	From    net.Addr
	Payload []byte
}
