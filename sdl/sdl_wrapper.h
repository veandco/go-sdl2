#if defined(_WIN32)
	#include <SDL2/SDL.h>
	#include <stdlib.h>
#else
	#include <SDL.h>
#endif

#if !defined(SDL_2_0_10)
#define SDL_2_0_10

#if !(SDL_VERSION_ATLEAST(2,0,10))

#if defined(WARN_OUTDATED)
#pragma message("SDL_FPoint is not supported before SDL 2.0.10")
#endif

typedef struct SDL_FPoint
{
    float x;
    float y;
} SDL_FPoint;

#if defined(WARN_OUTDATED)
#pragma message("SDL_FRect is not supported before SDL 2.0.10")
#endif

typedef struct SDL_FRect
{
    float x;
    float y;
    float w;
    float h;
} SDL_FRect;


#endif

#endif
