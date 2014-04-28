#include <SDL2/SDL.h>

extern int _SDL_GetSystemRAM();
extern SDL_bool _SDL_HasAVX();

extern char* _SDL_GetBasePath();
extern char* _SDL_GetPrefPath(const char *org, const char *app);
