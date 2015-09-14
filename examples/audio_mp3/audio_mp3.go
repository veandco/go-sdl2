package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

func main() {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Println(err)
		return
	}

	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Println(err)
		return
	}
	defer sdl.Quit()
	defer mix.Quit()
	defer mix.CloseAudio()

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}

	if music, err := mix.LoadMUS("../../assets/test.mp3"); err != nil {
		log.Println(err)
	} else if err = music.Play(1); err != nil {
		log.Println(err)
	} else {
		sdl.Delay(5000)
		music.Free()
	}
}
