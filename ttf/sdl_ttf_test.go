package ttf

import (
	"math/rand"
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ `!@#$%^&*()-=~_+[]\\{}|;':\",./<>?"

func randomString(t *testing.T) string {
	s := ""
	length := 16 + rand.Int()%16

	for i := 0; i < length; i++ {
		s += string(characters[rand.Int()%len(characters)])
	}
	t.Log("text:", s)

	return s
}

func TestTTF(t *testing.T) {
	var font *Font
	var solid *sdl.Surface
	var err error

	if err := Init(); err != nil {
		t.Errorf("Failed to initialize TTF: %s\n", err)
	}

	if font, err = OpenFont("../.go-sdl2-examples/assets/test.ttf", 32); err != nil {
		t.Errorf("Failed to open font: %s\n", err)
	}
	defer font.Close()

	for i := 0; i < 10000; i++ {
		if solid, err = font.RenderUTF8Solid(randomString(t), sdl.Color{255, 0, 0, 255}); err != nil {
			t.Errorf("Failed to render text: %s\n", err)
		}
		defer solid.Free()
	}

	gm, err := font.GlyphMetrics('A')
	if err != nil {
		t.Errorf("Failed to get glyph metrics: %s\n", err)
	}
	expectMetrics := GlyphMetrics{0, 22, 0, 23, 22}
	if *gm != expectMetrics {
		t.Errorf("GlyphMetrics got %v - want %v", *gm, expectMetrics)
	}
}

func TestGlyphIsProvided(t *testing.T) {
	var font *Font
	var err error

	if err := Init(); err != nil {
		t.Errorf("Failed to initialize TTF: %s\n", err)
	}

	if font, err = OpenFont("../.go-sdl2-examples/assets/test.ttf", 32); err != nil {
		t.Errorf("Failed to open font: %s\n", err)
	}

	asciiRune := 'A'

	if !font.GlyphIsProvided(uint16(asciiRune)) {
		t.Errorf("GlyphIsProvided(): rune not found in font that includes glyph")
	}

	if !font.GlyphIsProvided32(uint32(asciiRune)) {
		t.Errorf("GlyphIsProvided32(): rune not found in font that includes glyph")
	}

	emojiRune := '🤖'

	if font.GlyphIsProvided(uint16(emojiRune)) {
		t.Errorf("GlyphIsProvided(): rune found in font that doesn't support emojis")
	}

	if font.GlyphIsProvided32(uint32(emojiRune)) {
		t.Errorf("GlyphIsProvided32(): rune found in font that doesn't support emojis")
	}
}
