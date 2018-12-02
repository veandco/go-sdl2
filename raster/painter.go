package raster

import (
	"github.com/golang/freetype/raster"
	"image/color"
	"image/draw"
)

type ImagePainter struct {
	Image draw.Image
	c     color.Color
}

func (p *ImagePainter) Paint(ss []raster.Span, done bool) {
	// Convert color to RGBA
	dr, dg, db, da := p.c.RGBA() // 16 bit values
	// Alpha mask
	const m = 1<<16 - 1
	// Draw spans
	b := p.Image.Bounds()
	for _, s := range ss {
		if s.Y < b.Min.Y || s.Y >= b.Max.Y {
			continue
		}
		if s.X0 < b.Min.X {
			s.X0 = b.Min.X
		}
		if s.X1 > b.Max.X {
			s.X1 = b.Max.X
		}
		if s.X0 >= s.X1 {
			continue
		}
		for x := s.X0; x < s.X1; x++ {
			y := s.Y - b.Min.Y
			var ma uint32 = s.Alpha // 16 bit value
			// Get destination pixel color in RGBA64
			sr, sg, sb, sa := p.Image.At(x, y).RGBA() // 16 bit values
			// Compute destination color in RGBA64
			var a uint32 = (m - (da * ma / m))
			rr := uint16((dr*ma + sr*a) / m)
			gg := uint16((dg*ma + sg*a) / m)
			bb := uint16((db*ma + sb*a) / m)
			aa := uint16((da*ma + sa*a) / m)
			// Use image model to convert
			p.Image.Set(x, y, color.RGBA64{rr, gg, bb, aa})
		}
	}
}

func (p *ImagePainter) SetColor(c color.Color) {
	p.c = c
}

func NewImagePainter(m draw.Image) *ImagePainter {
	return &ImagePainter{Image: m}
}
