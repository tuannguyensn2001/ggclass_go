package bytepkg

import (
	"bytes"
	"encoding/gob"
)

func Encode(data interface{}) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}

	return network.Bytes(), nil
}
