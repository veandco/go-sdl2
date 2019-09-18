#include "_cgo_export.h"
#include "events.h"

SDL_Event event;

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
	SDL_FilterEvents((SDL_EventFilter)goEventFilterCallback, userdata);
}

void addEventWatch(void *userdata)
{
	SDL_AddEventWatch((SDL_EventFilter)goEventFilterCallback, userdata);
}

void delEventWatch(void *userdata)
{
	SDL_DelEventWatch((SDL_EventFilter)goEventFilterCallback, userdata);
}

int PollEvent()
{
	return SDL_PollEvent(&event);
}