#if defined(__WIN32)
	#include <SDL2/SDL_mixer.h>
#elif defined(__linux__)
	#include <SDL2/SDL_mixer.h>
#else
	#include <SDL_mixer.h>
#endif
