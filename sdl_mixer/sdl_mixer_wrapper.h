#if defined(__WIN32)
	#include <SDL2/SDL_mixer.h>
#elif defined(__APPLE__) && defined(__MACH__)
	#include <SDL2_mixer/SDL_mixer.h>
#else
	#include <SDL_mixer.h>
#endif
