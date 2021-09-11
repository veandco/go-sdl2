## 2.0.16

[x] SDL_FlashWindow() to get a user’s attention
[x] SDL_GetAudioDeviceSpec() to get the preferred audio format of a device
[x] SDL_SetWindowAlwaysOnTop() to dynamically change the SDL_WINDOW_ALWAYS_ON_TOP flag for a window
[x] SDL_SetWindowKeyboardGrab() to support grabbing the keyboard independently of the mouse
[x] SDL_SoftStretchLinear() to do bilinear scaling between 32-bit software surfaces
[x] SDL_UpdateNVTexture() to update streaming NV12/21 textures
[x] SDL_GameControllerSendEffect() and SDL_JoystickSendEffect() to allow sending custom trigger effects to the DualSense controller
[x] SDL_GameControllerGetSensorDataRate() to get the sensor data rate for PlayStation and Nintendo Switch controllers

### Windows

[?] SDL_SetWindowsMessageHook() to set a function that is called for all Windows messages
[?] SDL_RenderGetD3D11Device() to get the D3D11 device used by the SDL renderer

### Linux

[x] SDL_HINT_AUDIO_INCLUDE_MONITORS to control whether PulseAudio recording should include monitor devices
[x] SDL_HINT_AUDIO_DEVICE_STREAM_ROLE to describe the role of your application for audio control panels

### Android:

[?] SDL_AndroidShowToast() to show a lightweight notification

## 2.0.14

[x] SDL_GameControllerGetSerial()
[x] SDL_GameControllerHasAxis()
[x] SDL_GameControllerHasButton()
[x] SDL_GameControllerGetNumTouchpads()
[x] SDL_GameControllerGetNumTouchpadFingers()
[x] SDL_GameControllerGetTouchpadFinger()
[x] SDL_GameControllerHasSensor()
[x] SDL_GameControllerSetSensorEnabled()
[x] SDL_GameControllerIsSensorEnabled()
[x] SDL_GameControllerGetSensorData
[x] SDL_GameControllerRumbleTriggers()
[x] SDL_GameControllerHasLED()
[x] SDL_GameControllerSetLED()
[x] SDL_JoystickGetSerial()
[x] SDL_JoystickRumbleTriggers()
[x] SDL_JoystickHasLED()
[x] SDL_JoystickSetLED()
[x] SDL_JoystickAttachVirtual()
[x] SDL_JoystickDetachVirtual()
[x] SDL_JoystickIsVirtual()
[x] SDL_JoystickSetVirtualAxis()
[x] SDL_JoystickSetVirtualButton()
[x] SDL_JoystickSetVirtualHat()
[x] SDL_LockSensors()
[x] SDL_UnlockSensors()
[x] SDL_HAPTIC_STEERING_AXIS
[x] SDL_GetPreferredLocales()
[x] SDL_OpenURL()
[x] SDL_HasSurfaceRLE()
[x] SDL_SIMDRealloc()
[x] SDL_GetErrorMsg()
[x] SDL_Metal_GetLayer()
[x] SDL_Metal_GetDrawableSize()
[x] SDL_HINT_JOYSTICK_HIDAPI_PS5
[x] SDL_HINT_MOUSE_RELATIVE_SCALING
[x] SDL_HINT_PREFERRED_LOCALES
[x] SDL_HINT_JOYSTICK_RAWINPUT
[x] SDL_HINT_JOYSTICK_HIDAPI_CORRELATE_XINPUT
[x] SDL_HINT_AUDIO_DEVICE_APP_NAME
[x] SDL_HINT_AUDIO_DEVICE_STREAM_NAME
[x] SDL_HINT_LINUX_JOYSTICK_DEADZONES
[x] SDL_HINT_THREAD_PRIORITY_POLICY
[x] SDL_HINT_THREAD_FORCE_REALTIME_TIME_CRITICAL
[x] SDL_HINT_ANDROID_BLOCK_ON_PAUSE_PAUSEAUDIO
[x] SDL_HINT_EMSCRIPTEN_ASYNCIFY
[x] SDL_PIXELFORMAT_XRGB4444
[x] SDL_PIXELFORMAT_XBGR4444
[x] SDL_PIXELFORMAT_XRGB1555
[x] SDL_PIXELFORMAT_XBGR1555
[x] SDL_PIXELFORMAT_XRGB8888
[x] SDL_PIXELFORMAT_XBGR8888
[x] SDL_WINDOW_METAL
[x] SDL_AndroidRequestPermission()

## 2.0.12

[x] SDL_GameControllerTypeForIndex
[x] SDL_GameControllerGetType
[x] SDL_GameControllerFromPlayerIndex
[x] SDL_GameControllerSetPlayerIndex
[x] SDL_JoystickFromPlayerIndex
[x] SDL_JoystickSetPlayerIndex
[x] SDL_GetTextureScaleMode
[x] SDL_SetTextureScaleMode
[x] SDL_LockTextureToSurface
[x] SDL_BLENDMODE_MUL
[x] SDL_TouchFingerEvent update
[x] SDL_HasARMSIMD
[x] SDL_HINT_DISPLAY_USABLE_BOUNDS
[x] SDL_HINT_GAMECONTROLLERTYPE
[x] SDL_HINT_GAMECONTROLLER_USE_BUTTON_LABELS
[x] SDL_HINT_JOYSTICK_HIDAPI_GAMECUBE
[x] SDL_HINT_VIDEO_X11_WINDOW_VISUALID
[x] SDL_HINT_VIDEO_X11_FORCE_EGL
[x] SDL_Metal_CreateView
[x] SDL_Metal_DestroyView
[x] SDL_GetAndroidSDKVersion

## 2.0.10

[x] SDL_SIMDGetAlignment
[x] SDL_SIMDAlloc
[x] SDL_SIMDFree
[x] SDL_RenderDrawPointF
[x] SDL_RenderDrawPointsF
[x] SDL_RenderDrawLineF
[x] SDL_RenderDrawLinesF
[x] SDL_RenderDrawRectF
[x] SDL_RenderDrawRectsF
[x] SDL_RenderFillRectF
[x] SDL_RenderFillRectsF
[x] SDL_RenderCopyF
[x] SDL_RenderCopyExF
[x] SDL_GetTouchDeviceType
[x] SDL_RenderFlush
[x] SDL_HINT_RENDER_BATCHING
[x] SDL_HINT_EVENT_LOGGING
[x] SDL_HINT_GAMECONTROLLERCONFIG_FILE
[x] SDL_HINT_MOUSE_TOUCH_EVENTS

## 2.0.9

[x] SDL_SENSORUPDATE
[x] SDL_DISPLAYEVENT
[x] SDL_JoystickGetDevicePlayerIndex
[x] SDL_JoystickGetPlayerIndex
[x] SDL_GameControllerGetPlayerIndex
[x] SDL_GameControllerRumble
[x] SDL_JoystickRumble
[x] SDL_GameControllerMappingForDeviceIndex
[x] SDL_HINT_MOUSE_DOUBLE_CLICK_TIME
[x] SDL_HINT_MOUSE_DOUBLE_CLICK_RADIUS
[x] SDL_HasColorKey
[x] SDL_HasAVX512F
[x] SDL_IsTablet
[?] SDL_THREAD_PRIORITY_TIME_CRITICAL

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
