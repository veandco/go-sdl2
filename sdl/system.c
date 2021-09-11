#include "_cgo_export.h"
#include "system.h"

#if defined(_WIN32)
void SetWindowsMessageHook()
{
	SDL_SetWindowsMessageHook((SDL_WindowsMessageHook) goWindowsMessageHook, NULL);
}
#endif
