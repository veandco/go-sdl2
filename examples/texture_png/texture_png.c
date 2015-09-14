// Test program in C to see behavior of IMG_Load()

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

static const int WIN_WIDTH = 800;
static const int WIN_HEIGHT = 600;

int main()
{
	SDL_Window *window = SDL_CreateWindow(
			"Go-SDL2 Texture - C version",
			SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED,
			WIN_WIDTH, WIN_HEIGHT,
			SDL_WINDOW_SHOWN);
	if (!window) {
		fprintf(stderr, "Failed to create window: %s\n", SDL_GetError());
		return 1;
	}

	SDL_Renderer *renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);
	if (!renderer) {
		fprintf(stderr, "Failed to create renderer: %s\n", SDL_GetError());
		return 2;
	}

	SDL_Surface *image = IMG_Load("../../assets/test.png");
	if (!image) {
		fprintf(stderr, "Failed to load PNG: %s\n", SDL_GetError());
		return 3;
	}

	SDL_Texture *texture = SDL_CreateTextureFromSurface(renderer, image);
	if (!texture) {
		fprintf(stderr, "Failed to cerate texture: %s\n", SDL_GetError());
		return 4;
	}

	SDL_Rect src = {0, 0, 512, 512};
	SDL_Rect dst = {100, 50, 512, 512};
	SDL_Rect bound = {0, 0, WIN_WIDTH, WIN_HEIGHT};

	SDL_RenderClear(renderer);
	SDL_SetRenderDrawColor(renderer, 255, 0, 0, 255);
	SDL_RenderFillRect(renderer, &bound);
	SDL_RenderCopy(renderer, texture, &src, &dst);
	SDL_RenderPresent(renderer);

	SDL_Delay(2000);

	SDL_DestroyTexture(texture);
	SDL_FreeSurface(image);
	SDL_DestroyRenderer(renderer);
	SDL_DestroyWindow(window);

	return 0;
}
