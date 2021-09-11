#ifndef _GO_SDL_SYSTEM_H
#define _GO_SDL_SYSTEM_H

#if defined(_WIN32)
	#include <SDL2/SDL_system.h>

	extern void SetWindowsMessageHook();
#endif

#endif
