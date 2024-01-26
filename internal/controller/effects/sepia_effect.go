package effects

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
)

type SepiaEffect struct {
}

func (s *SepiaEffect) Apply(data []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, format, err
	}

	bounds := img.Bounds()
	sepiaImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()
			newR, newG, newB := s.sepiaTone(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			sepiaColor := color.RGBA{newR, newG, newB, 0xff}
			sepiaImg.Set(x, y, sepiaColor)
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, sepiaImg, nil); err != nil {
		return nil, format, err
	}

	return buf.Bytes(), format, nil
}

func (s *SepiaEffect) sepiaTone(r, g, b uint8) (uint8, uint8, uint8) {
	newR := uint8(0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b))
	newG := uint8(0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b))
	newB := uint8(0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b))

	if newR > 255 {
		newR = 255
	}
	if newG > 255 {
		newG = 255
	}
	if newB > 255 {
		newB = 255
	}

	return newR, newG, newB
}

func NewSepiaEffect() *SepiaEffect {
	return &SepiaEffect{}
}
