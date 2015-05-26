#if defined(__WIN32)
	#include <SDL2/SDL.h>
    #include <stdlib.h>
#elif defined(__APPLE__) && defined(__MACH__)
	#include <SDL2/SDL.h>
#else
	#include <SDL.h>
#endif
