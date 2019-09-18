#ifndef _GO_SDL_EVENTS_H
#define _GO_SDL_EVENTS_H

#if defined(_WIN32)
	#include <SDL2/SDL_events.h>
#else
	#include <SDL_events.h>
#endif

extern SDL_Event event;

extern void setEventFilter();
extern void clearEventFilter();
extern void filterEvents(void *userdata);
extern void addEventWatch(void *userdata);
extern void delEventWatch(void *userdata);
extern int PollEvent();

#endif
