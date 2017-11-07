#include "_cgo_export.h"

#include "sdl_wrapper.h"
#include "hints.h"

void hintCallback(void *userdata, const char *name, const char *oldValue, const char *newValue)
{
	goHintCallback((char *) name, (char *) oldValue, (char *) newValue);
}

void addHintCallback(const char *name)
{
	SDL_AddHintCallback(name, hintCallback, NULL);
}

void delHintCallback(const char *name)
{
	SDL_DelHintCallback(name, hintCallback, NULL);}
