#include "_cgo_export.h"
#include "events.h"

void setGoEventFilter()
{
    SDL_SetEventFilter((SDL_EventFilter)goEventFilter, NULL);
}

void clearGoEventFilter()
{
    SDL_SetEventFilter(NULL, NULL);
}
