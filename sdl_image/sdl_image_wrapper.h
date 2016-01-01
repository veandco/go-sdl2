#if defined(__WIN32)
	#include <SDL2/SDL_image.h>
	#include <stdlib.h>
#elif defined(__linux__)
	#include <SDL2/SDL_image.h>
#else
	#include <SDL_image.h>
#endif
