#if defined(_WIN32)
	#include <SDL2/SDL.h>
	#include <stdlib.h>
#else
	#include <SDL.h>
#endif

#if !defined(SDL_WINDOW_ALLOW_HIGHDPI)
	#define SDL_WINDOW_ALLOW_HIGHDPI (0x00002000)
#endif
