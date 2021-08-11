package helpers

import (
	"bytes"
	"image"
	"image/color"
	"reflect"

	"github.com/nfnt/resize"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

// ScaleImage resizes an image to the given width and height.
func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	return img, w, h
}

// ConvertToAscii converts an image to ASCII.
func ConvertToAscii(img image.Image, w, h int) string {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}

	return string(buf.Bytes())
}
