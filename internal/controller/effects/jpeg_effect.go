package effects

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
)

type JPEGEffect struct {
}

func (j *JPEGEffect) Apply(data []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, format, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, format, err
	}

	return buf.Bytes(), "jpeg", nil
}

func NewJPEGEffect() *JPEGEffect {
	return &JPEGEffect{}
}
