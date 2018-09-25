## 2.0.8

### Hints

[x] SDL_HINT_IOS_HIDE_HOME_INDICATOR
[x] SDL_HINT_RETURN_KEY_HIDES_IME
[x] SDL_HINT_TV_REMOTE_AS_JOYSTICK
[x] SDL_HINT_VIDEO_X11_NET_WM_BYPASS_COMPOSITOR
[x] SDL_HINT_VIDEO_DOUBLE_BUFFER

### Surface

[x] SDL_SetYUVConversionMode()
[x] SDL_GetYUVConversionMode()

### Android

[x] SDL_IsAndroidTV()

### Mac OS X / iOS / tvOS

[x] SDL_RenderGetMetalLayer()
[x] SDL_RenderGetMetalCommandEncoder()

### Windows UWP

[ ] SDL_WinRTGetDeviceFamily()

## 2.0.7

### General

Audio

- [x] SDL_NewAudioStream()
- [x] SDL_AudioStreamPut()
- [x] SDL_AudioStreamGet()
- [x] SDL_AudioStreamAvailable()
- [x] SDL_AudioStreamFlush()
- [x] SDL_AudioStreamClear()
- [x] SDL_FreeAudioStream()

Joystick

- [x] SDL_LockJoysticks()
- [x] SDL_UnlockJoysticks()

Stdinc

- [ ] SDL_GetMemoryFunctions()
- [ ] SDL_SetMemoryFunctions()
- [ ] SDL_GetNumAllocations()

## 2.0.6

### General

Blend Mode

- [x] SDL_ComposeCustomBlendMode()

CPU Info

- [x] SDL_HasNEON()

Game Controller

- [x] SDL_GameControllerGetVendor()
- [x] SDL_GameControllerGetProduct()
- [x] SDL_GameControllerGetProductVersion()
- [x] SDL_GameControllerNumMappings()
- [x] SDL_GameControllerMappingForIndex()

Hints

- [x] SDL_HINT_AUDIO_RESAMPLING_MODE
- [x] SDL_HINT_RENDER_LOGICAL_SIZE_MODE
- [x] SDL_HINT_MOUSE_NORMAL_SPEED_SCALE
- [x] SDL_HINT_MOUSE_RELATIVE_SPEED_SCALE
- [x] SDL_HINT_TOUCH_MOUSE_EVENTS

Joystick

- [x] SDL_JoystickGetDeviceVendor()
- [x] SDL_JoystickGetDeviceProduct()
- [x] SDL_JoystickGetDeviceProductVersion()
- [x] SDL_JoystickGetDeviceType()
- [x] SDL_JoystickGetDeviceInstanceID()
- [x] SDL_JoystickGetVendor()
- [x] SDL_JoystickGetProduct()
- [x] SDL_JoystickGetProductVersion()
- [x] SDL_JoystickGetType()
- [x] SDL_JoystickGetAxisInitialState()

RW Ops

- [x] SDL_LoadFile()
- [x] SDL_LoadFile_RW()

Surface

- [x] SDL_DuplicateSurface()

Vulkan

- [x] SDL_Vulkan_LoadLibrary()
- [x] SDL_Vulkan_GetVkGetInstanceProcAddr()
- [x] SDL_Vulkan_GetInstanceExtensions()
- [x] SDL_Vulkan_CreateSurface()
- [x] SDL_Vulkan_GetDrawableSize()
- [x] SDL_Vulkan_UnloadLibrary()

### Windows

Hints

- [x] SDL_HINT_WINDOWS_INTRESOURCE_ICON
- [x] SDL_HINT_WINDOWS_INTRESOURCE_ICON_SMALL

## Miscellaneous

- [ ] Add ability to set window title bar color on runtime
