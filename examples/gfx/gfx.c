#include <SDL2/SDL.h>
#include <SDL2/SDL2_gfxPrimitives.h>
#include <stdlib.h>

#define WIDTH (800)
#define HEIGHT (600)

int main()
{
	SDL_Window *window;
	SDL_Renderer *renderer;
	int x, y;

	window = SDL_CreateWindow("SDL2 GFX", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_OPENGL);
	if (!window) {
		fprintf(stderr, "error: %s\n", SDL_GetError());
		return -1;
	}

	renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);
	if (!renderer) {
		fprintf(stderr, "error: %s\n", SDL_GetError());
		return -1;
	}

	for (y = 0; y < HEIGHT; y++) {
		for (x = 0; x < WIDTH; x++) {
			pixelColor(renderer, x, y, rand());
		}
	}
	SDL_RenderPresent(renderer);
	SDL_Delay(3000);

	SDL_DestroyRenderer(renderer);
	SDL_DestroyWindow(window);
}
