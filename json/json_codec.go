package json

import (
	"encoding/json"
	"io"

	"github.com/edenzhong7/ifly/icodec"
)

type Codec struct{}

func NewCodec() Codec {
	return Codec{}
}

func (c Codec) Name() string {
	return "json"
}

func (c Codec) NewEncoder(w io.Writer) icodec.Encoder {
	return json.NewEncoder(w)
}

func (c Codec) NewDecoder(r io.Reader) icodec.Decoder {
	return json.NewDecoder(r)
}
