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

	if font, err = OpenFont("../assets/test.ttf", 32); err != nil {
		t.Errorf("Failed to open font: %s\n", err)
	}
	defer font.Close()

	for i := 0; i < 10000; i++ {
		if solid, err = font.RenderUTF8_Solid(randomString(t), sdl.Color{255, 0, 0, 255}); err != nil {
			t.Errorf("Failed to render text: %s\n", err)
		}
		defer solid.Free()
	}
}
