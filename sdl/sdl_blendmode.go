package sdl

// #include <SDL2/SDL.h>
import "C"

type BlendMode C.SDL_BlendMode

const (
	BLENDMODE_NONE			= 0x00000000     /**< No blending */
	BLENDMODE_BLEND			= 0x00000001    /**< dst = (src * A) + (dst * (1-A)) */
	BLENDMODE_ADD			= 0x00000002      /**< dst = (src * A) + dst */
	BLENDMODE_MOD			= 0x00000004       /**< dst = src * dst */
)
