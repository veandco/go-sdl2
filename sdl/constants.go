package sdl

const (
	/* SDL.h */
	INIT_TIMER			= 0x00000001
	INIT_AUDIO			= 0x00000010
	INIT_VIDEO			= 0x00000020
	INIT_JOYSTICK			= 0x00000200
	INIT_HAPTIC			= 0x00001000
	INIT_GAMECONTROLLER		= 0x00002000      /**< turn on game controller also implicitly does JOYSTICK */
	INIT_NOPARACHUTE		= 0x00100000      /**< Don't catch fatal signals */
	INIT_EVERYTHING			= INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_JOYSTICK |
					  INIT_HAPTIC | INIT_GAMECONTROLLER

	WINDOWPOS_UNDEFINED_MASK	= 0x1FFF0000
	WINDOWPOS_UNDEFINED		= WINDOWPOS_UNDEFINED_MASK | 0

	/* SDL_video.h */
	WINDOW_FULLSCREEN		= 0x00000001         /**< fullscreen window */
	WINDOW_OPENGL			= 0x00000002             /**< window usable with OpenGL context */
	WINDOW_SHOWN			= 0x00000004              /**< window is visible */
	WINDOW_HIDDEN			= 0x00000008             /**< window is not visible */
	WINDOW_BORDERLESS		= 0x00000010         /**< no window decoration */
	WINDOW_RESIZABLE		= 0x00000020          /**< window can be resized */
	WINDOW_MINIMIZED		= 0x00000040          /**< window is minimized */
	WINDOW_MAXIMIZED		= 0x00000080          /**< window is maximized */
	WINDOW_INPUT_GRABBED		= 0x00000100      /**< window has grabbed input focus */
	WINDOW_INPUT_FOCUS		= 0x00000200        /**< window has input focus */
	WINDOW_MOUSE_FOCUS		= 0x00000400        /**< window has mouse focus */
	WINDOW_FULLSCREEN_DESKT		= WINDOW_FULLSCREEN | 0x00001000
	WINDOW_FOREIGN			= 0x00000800             /**< window not created by SDL */
)
