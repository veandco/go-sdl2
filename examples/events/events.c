// author: Jacky Boen

#include <SDL2/SDL.h>

static const char *window_title = "SDL2 Events";
static const int window_width = 800, window_height = 600;
static SDL_Joystick *joysticks[16];

int main() {
	SDL_Window *window;
	SDL_bool running = SDL_TRUE;

	SDL_Init(SDL_INIT_EVERYTHING);

	window = SDL_CreateWindow(window_title, SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, window_width, window_height, SDL_WINDOW_SHOWN);
	if (!window) {
		fprintf(stderr, "%s\n", SDL_GetError());
		return -1;
	}

	SDL_JoystickEventState(SDL_ENABLE);

	while (running) {
		SDL_Event event;

		while (SDL_PollEvent(&event)) {
			switch (event.type) {
			case SDL_QUIT:
				running = SDL_FALSE;
				break;
			case SDL_MOUSEMOTION:
				fprintf(stdout, "[%d ms] MouseMotion\ttype:%d\twhich:\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					event.motion.timestamp, event.motion.type, event.motion.which, event.motion.x, event.motion.y, event.motion.xrel, event.motion.yrel);
				break;
			case SDL_MOUSEBUTTONDOWN:
			case SDL_MOUSEBUTTONUP:
				fprintf(stdout, "[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					event.button.timestamp, event.button.type, event.button.which, event.button.x, event.button.y, event.button.button, event.button.state);
				break;
			case SDL_MOUSEWHEEL:
				fprintf(stdout, "[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					event.wheel.timestamp, event.wheel.type, event.wheel.which, event.wheel.x, event.wheel.y);
				break;
			case SDL_KEYDOWN:
			case SDL_KEYUP:
				fprintf(stdout, "[%d ms] Keyboard\ttype:%d\twindowID:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					event.key.timestamp, event.key.type, event.key.windowID, event.key.keysym.sym, event.key.keysym.mod, event.key.state, event.key.repeat);
				break;
			case SDL_JOYAXISMOTION:
				fprintf(stdout, "[%d ms] JoyAxis\ttype:%d\twhich:%d\taxis:%d\tvalue:%d\n",
					event.jaxis.timestamp, event.jaxis.type, event.jaxis.which, event.jaxis.axis, event.jaxis.value);
				break;
			case SDL_JOYBALLMOTION:
				fprintf(stdout, "[%d ms] JoyBall\ttype:%d\twhich:%d\tball:%d\txrel:%d\tyrel:%d\n",
					event.jball.timestamp, event.jball.type, event.jball.which, event.jball.ball, event.jball.xrel, event.jball.yrel);
				break;
			case SDL_JOYBUTTONUP:
			case SDL_JOYBUTTONDOWN:
				fprintf(stdout, "[%d ms] JoyButton\ttype:%d\twhich:%d\tbutton:%d\tstate:%d\n",
					event.jbutton.timestamp, event.jbutton.type, event.jbutton.which, event.jbutton.button, event.jbutton.state);
				break;
			case SDL_JOYHATMOTION:
				fprintf(stdout, "[%d ms] JoyHat\ttype:%d\twhich:%d\that:%d\tvalue:%d\n",
					event.jhat.timestamp, event.jhat.type, event.jhat.which, event.jhat.hat, event.jhat.value);
				break;
			case SDL_JOYDEVICEADDED:
				joysticks[event.jdevice.which] = SDL_JoystickOpen(event.jdevice.which);
				if (joysticks[event.jdevice.which]) {
					fprintf(stdout, "Joystick %d connected\n", event.jdevice.which);
				}
				break;
			case SDL_JOYDEVICEREMOVED:
				if (!joysticks[event.jdevice.which]) {
					SDL_JoystickClose(joysticks[event.jdevice.which]);
				}
				fprintf(stdout, "Joystick %d disconnected\n", event.jdevice.which);
				break;
			default:
				fprintf(stdout, "Some event\n");
				break;
			}
		}
	}

	SDL_DestroyWindow(window);
	SDL_Quit();

	return 0;
}
