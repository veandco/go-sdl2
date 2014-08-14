#include "_cgo_export.h"
#include "events.h"

void setEventFilter()
{
    SDL_SetEventFilter((SDL_EventFilter)goSetEventFilterCallback, NULL);
}

void clearEventFilter()
{
    SDL_SetEventFilter(NULL, NULL);
}

void filterEvents(void *userdata)
{
	SDL_FilterEvents((SDL_EventFilter)goFilterEventsCallback, userdata);
}