package puny

import (
	"github.com/junjun-cai/codec/base"
	"golang.org/x/net/idna"
)

var StdCodec, _ = NewCodec()

type punyCodec struct {
	codec *idna.Profile
}

func NewCodec() (base.IEncoding, error) {
	return &punyCodec{codec: &idna.Profile{}}, nil
}

func (p *punyCodec) Encode(src []byte) ([]byte, error) {
	dst, err := p.codec.ToASCII(base.BytesToString(src))
	if err != nil {
		return nil, err
	}
	return base.StringToBytes(dst), nil
}

func (p *punyCodec) Decode(src []byte) ([]byte, error) {
	dst, err := p.codec.ToUnicode(base.BytesToString(src))
	if err != nil {
		return nil, err
	}
	return base.StringToBytes(dst), nil
}
