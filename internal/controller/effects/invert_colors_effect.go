package effects

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
)

type InvertColorsEffect struct {
}

func (i *InvertColorsEffect) Apply(data []byte) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, format, err
	}

	invertedImg := InvertColors(img)

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, invertedImg, nil); err != nil {
		return nil, format, err
	}

	return buf.Bytes(), format, nil
}

func InvertColors(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	invertedImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			invertedColor := color.RGBA{R: uint8(255 - r/256), G: uint8(255 - g/256), B: uint8(255 - b/256), A: uint8(a / 256)}
			invertedImage.Set(x, y, invertedColor)
		}
	}

	return invertedImage
}

func NewInvertColorsEffect() *InvertColorsEffect {
	return &InvertColorsEffect{}
}
