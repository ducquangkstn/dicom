package dicomio

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrorLimitReached = errors.New("buffer limit reached")

type Reader interface {
	io.Reader
	ReadN(n uint32) ([]byte, error)
	ReadUInt16() (uint16, error)
	ReadUInt32() (uint32, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadString(n uint32) (string, error)
	Skip(n uint) error
	PushLimit(n uint)
	PopLimit()
}

type reader struct {
	in        io.Reader
	bo        binary.ByteOrder
	limit     int64
	bytesRead int64
}

// TODO: implement reader