#if defined(__WIN32)
	#include <SDL2/SDL_ttf.h>
    #include <stdlib.h>
#elif defined(__APPLE__) && defined(__MACH__)
	#include <SDL2_ttf/SDL_ttf.h>
#else
	#include <SDL_ttf.h>
#endif
