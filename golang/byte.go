package main

import (
	"bytes"
	"encoding/binary"
)

type MyType struct {
	A int32
	B float64
}

func MyTypeSliceToBytesSlice(data []MyType) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, value := range data {
		if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
