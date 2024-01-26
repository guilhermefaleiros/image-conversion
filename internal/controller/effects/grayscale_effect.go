package effects

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
)

type GrayscaleEffect struct {
}

func (g *GrayscaleEffect) Apply(data []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, format, err
	}

	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, grayImg, nil); err != nil {
		return nil, format, err
	}

	return buf.Bytes(), format, nil
}

func NewGrayscaleEffect() *GrayscaleEffect {
	return &GrayscaleEffect{}
}
