package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buffer := make([]byte, 1028)
	n, err := r.Read(buffer)
	if err != nil {
		return err
	}
	msg.Payload = buffer[:n]

	return nil
}
