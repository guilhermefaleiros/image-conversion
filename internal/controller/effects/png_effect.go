package effects

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"image/png"
)

type PNGEffect struct {
}

func (j *PNGEffect) Apply(data []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, format, err
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, format, err
	}

	return buf.Bytes(), "png", nil
}

func NewPNGEffect() *PNGEffect {
	return &PNGEffect{}
}
