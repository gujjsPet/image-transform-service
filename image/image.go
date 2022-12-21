package image

import (
	"image"
	"image/draw"
	"io"

	_ "image/gif"
	_ "image/jpeg"
	png "image/png"
)

func DecodeImage(r io.Reader) (image.Image, error) {
	m, _, err := image.Decode(r)
	return m, err
}

func EncodeImage(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}

func CropImage(src image.Image, x0, y0, x1, y1 int) image.Image {

	r := image.Rect(x0, y0, x1, y1)
	dr := r.Sub(r.Min)
	dst := image.NewRGBA(dr)

	draw.Draw(dst, dst.Bounds(), image.Transparent, image.Point{}, draw.Src)
	draw.Draw(dst, dst.Bounds(), src, r.Min, draw.Src)

	return dst
}
