#ifndef _GO_SDL_EVENTS_H
#define _GO_SDL_EVENTS_H

extern void setEventFilter();
extern void clearEventFilter();
extern void filterEvents(void *userdata);

#endif
