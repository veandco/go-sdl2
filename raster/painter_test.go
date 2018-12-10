package raster

import (
	"testing"

	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/golang/freetype/raster"
)

func TestImagePainter(t *testing.T) {
	// Spans to render
	ss := []raster.Span{
		raster.Span{0, 4, 6, 0xffff},
		raster.Span{1, 3, 7, 0xf0f0},
		raster.Span{2, 2, 8, 0x8f8f},
		raster.Span{3, 1, 9, 0x8080},
		raster.Span{4, 0, 10, 0x0f0f},
	}

	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	pt := NewImagePainter(img)
	pt.SetColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	pt.Paint(ss, false)

	// Write to PNG
	f, err := os.Create("output.png")
	if err != nil {
		t.Error(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		t.Error(err)
	}
	if err := f.Close(); err != nil {
		t.Error(err)
	}
}
