package p2p

/**
* Message holds any arbitrary data that's being sent over
* each transport b/w two nodes in the network.
 */
type Message struct {
	Payload []byte
}
