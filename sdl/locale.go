package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

typedef struct SDL_Locale
{
    const char *language;
    const char *country;
} SDL_Locale;

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetPreferredLocales is not supported before SDL 2.0.14")
#endif

static SDL_Locale * SDL_GetPreferredLocales(void)
{
	return NULL;
}

#endif
*/
import "C"

type Locale struct {
	Language string // A language name, like "en" for English.
	Country  string // A country, like "US" for America. Can be empty.
}

func GetPreferredLocales() (locale Locale, err error) {
	_locale := C.SDL_GetPreferredLocales()
	if _locale == nil {
		err = GetError()
		return
	}
	locale.Language = C.GoString(_locale.language)
	locale.Country = C.GoString(_locale.country)
	return
}
