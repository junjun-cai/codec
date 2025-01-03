package quotedprintable

import (
	"bytes"
	"mime/quotedprintable"
)

var StdCodec, _ = NewCodec()

type quotedprintableCodec struct {
}

func NewCodec() (*quotedprintableCodec, error) {
	return &quotedprintableCodec{}, nil
}

func (q *quotedprintableCodec) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := quotedprintable.NewWriter(&buf)
	_, err := writer.Write(src)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (q *quotedprintableCodec) Decode(src []byte) ([]byte, error) {
	reader := quotedprintable.NewReader(bytes.NewReader(src))
	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
