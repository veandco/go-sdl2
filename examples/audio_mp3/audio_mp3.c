#include <SDL2/SDL.h>
#include <SDL2/SDL_mixer.h>

int main()
{
	Mix_Music *music = 0;
	int ret = 0;

	ret = SDL_Init(SDL_INIT_AUDIO);
	if (ret) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return 1;
	}

	ret = Mix_Init(MIX_INIT_MP3);
	if ((ret & MIX_INIT_MP3) != MIX_INIT_MP3) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return 2;
	}

	ret = Mix_OpenAudio(22050, MIX_DEFAULT_FORMAT, 2, 4096);
	if (ret < 0) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return 3;
	}

	music = Mix_LoadMUS("../../assets/test.mp3");
	if (!music) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return 4;
	}

	ret = Mix_PlayMusic(music, 1);
	if (ret == -1) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return 5;
	}

	SDL_Delay(5000);

	Mix_FreeMusic(music);
	Mix_CloseAudio();
	Mix_Quit();
	SDL_Quit();
}
