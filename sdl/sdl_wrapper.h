#if defined(__WIN32)
	#include <SDL2/SDL.h>
	#include <stdlib.h>
#elif defined(__linux__)
	#include <SDL2/SDL.h>
#else
	#include <SDL.h>
#endif
