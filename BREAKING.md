## BREAKING CHANGES

### v0.3..master

#### sdl

+ Changed `RWFromMem()` (0ee14f91)
+ Changed and renamed `RWops.FreeRW()` to `RWops.Free()` (0ee14f91)
+ Changed and renamed `RWops.RWsize()` to `RWops.Size()` (0ee14f91)
+ Changed and renamed `RWops.RWseek()` to `RWops.Seek()` (0ee14f91)
+ Changed and renamed `RWops.RWread()` to `RWops.Read()` and `RWops.Read2()` (0ee14f91)
+ Changed and renamed `RWops.RWtell()` to `RWops.Tell()` (0ee14f91)
+ Changed and renamed `RWops.RWwrite()` to `RWops.Write()` and `RWops.Write2()` (0ee14f91)
+ Changed and renamed `RWops.RWclose()` to `RWops.Close()` (0ee14f91)
+ Changed and renamed `RWops.RWclose()` to `RWops.Close()` (0ee14f91)
+ Changed `Window.VulkanCreateSurface` to return `unsafe.Pointer` instead of `uintptr`

#### img

+ Changed `Init()` to return `error` instead of `int`

### v0.2..v0.3

+ Renamed `ButtonId` to `ButtonID` in `MessageBoxButtonData` struct
+ Renamed `GL_CreateContext()` to `GLCreateContext()`
+ Renamed `GL_DeleteContext()` to `GLDeleteContext()`
+ Renamed `GL_ExtensionSupported()` to `GLExtensionSupported()`
+ Renamed `GL_GetAttribute()` to `GLGetAttribute()`
+ Renamed `GL_GetDrawableSize()` to `GLGetDrawableSize()`
+ Renamed `GL_GetProcAddress()` to `GLGetProcAddress()`
+ Renamed `GL_GetSwapInterval()` to `GLGetSwapInterval()`
+ Renamed `GL_LoadLibrary()` to `GLLoadLibrary()`
+ Renamed `GL_MakeCurrent()` to `GLMakeCurrent()`
+ Renamed `GL_SetAttribute()` to `GLSetAttribute()`
+ Renamed `GL_SetSwapInterval()` to `GLSetSwapInterval()`
+ Renamed `GL_SwapWindow()` to `GLSwap()`
+ Renamed `GL_UnloadLibrary()` to `GLUnloadLibrary()`
+ Renamed `GameController.GetAttached()` to `GameController.Attached()`
+ Renamed `GameController.GetAxis()` to `GameController.Axis()`
+ Renamed `GameController.GetBindForAxis()` to `GameController.BindForAxis()`
+ Renamed `GameController.GetBindForButton()` to `GameController.BindForButton()`
+ Renamed `GameController.GetButton()` to `GameController.Button()`
+ Renamed `GameController.GetJoystick()` to `GameController.Joystick()`
+ Renamed `Id` to `ID` in `Finger` struct
+ Renamed `Joystick.GetAttached()` to `Joystick.Attached()`
+ Renamed `Joystick.GetAxis()` to `Joystick.Axis()`
+ Renamed `Joystick.GetBall()` to `Joystick.Ball()`
+ Renamed `Joystick.GetButton()` to `Joystick.Button()`
+ Renamed `Joystick.GetGUID()` to `Joystick.GUID()`
+ Renamed `Joystick.GetHat()` to `Joystick.Hat()`
+ Renamed `LoadBMP_RW()` to `LoadBMPRW()`
+ Renamed `LoadBMP_RW()` to `LoadBMPRW()`
+ Renamed `LoadCUR_RW()` to `LoadCURRW()`
+ Renamed `LoadGIF_RW()` to `LoadGIFRW()`
+ Renamed `LoadICO_RW()` to `LoadICORW()`
+ Renamed `LoadJPG_RW()` to `LoadJPGRW()`
+ Renamed `LoadLBM_RW()` to `LoadLBMRW()`
+ Renamed `LoadMUSType_RW()` to `LoadMUSTypeRW()`
+ Renamed `LoadMUS_RW()` to `LoadMUSRW()`
+ Renamed `LoadPCX_RW()` to `LoadPCXRW()`
+ Renamed `LoadPNG_RW()` to `LoadPNGRW()`
+ Renamed `LoadPNM_RW()` to `LoadPNMRW()`
+ Renamed `LoadTGA_RW()` to `LoadTGARW()`
+ Renamed `LoadTIF_RW()` to `LoadTIFRW()`
+ Renamed `LoadTexture_RW` to `LoadTextureRW`
+ Renamed `LoadTyped_RW` to `LoadTypedRW`
+ Renamed `LoadWAV_RW()` to `LoadWAVRW()`
+ Renamed `LoadWAV_RW()` to `LoadWAVRW()`
+ Renamed `LoadWEBP_RW()` to `LoadWEBPRW()`
+ Renamed `LoadXCF_RW()` to `LoadXCFRW()`
+ Renamed `LoadXPM_RW()` to `LoadXPMRW()`
+ Renamed `LoadXV_RW()` to `LoadXVRW()`
+ Renamed `Load_RW` to `LoadRW`
+ Renamed `QuickLoad_RAW()` to `QuickLoadRAW()`
+ Renamed `QuickLoad_WAV()` to `QuickLoadWAV()`
+ Renamed `RenderUTF8_Blended()` to `RenderUTF8Blended()`
+ Renamed `RenderUTF8_Blended_Wrapped()` to `RenderUTF8BlendedWrapped()`
+ Renamed `RenderUTF8_Shaded()` to `RenderUTF8Shaded()`
+ Renamed `RenderUTF8_Solid()` to `RenderUTF8Solid()`
+ Renamed `SW_YUVTexture` to `SWYUVTexture`
+ Renamed `SaveBMP_RW()` to `SaveBMPRW()`
+ Renamed `SavePNG_RW()` to `SavePNGRW()`
+ Renamed `Texture.GL_BindTexture()` to `Texture.GLBind()`
+ Renamed `Texture.GL_UnbindTexture()` to `Texture.GLUnbind()`
+ Renamed `TouchId` to `TouchID` in `MultiGestureEvent` struct
+ Renamed `Unicode` to `unused` in `Keysym` struct (must have been a typo)
+ `GLCreateContext()`, `GLMakeCurrent()`, `GLGetDrawableSize()`, `GLSwapWindow()` are now methods of `Window`
+ `GetCurrentDisplayMode()` returns (DisplayMode, error) instead of error
+ `GetCurrentVideoDriver()` returns (string, error) instead of string
+ `GetDesktopDisplayMode()` returns (DisplayMode, error) instead of error
+ `GetDisplayBounds()` returns (Rect, error) instead of error
+ `GetDisplayMode()` returns (DisplayMode, error) instead of error
+ `GetDisplayName()` returns (string, error) instead of string
+ `GetDisplayUsableBounds` returns (Rect, error) instead of error
+ `GetNumRenderDrivers()` returns (int, error) instead of int
+ `GetRenderDriverInfo()` returns (int, error) instead of int
+ `Haptic.GetEffectStatus()` returns (int, error) instead of int
+ `Haptic.NumAxes()` returns (int, error) instead of int
+ `Haptic.NumEffects()` returns (int, error) instead of int
+ `Haptic.NumEffectsPlaying()` returns (int, error) instead of int
+ `Haptic.Query()` returns (uint32, error) instead of uint
+ `HapticIndex()` returns (int, error) instead of int
+ `HapticName()` returns (string, error) instead of string
+ `HapticOpenFromJoystick()` returns (\*Haptic, error) instead of \*Haptic
+ `LoadBMPRW()` requires (\*RWops, bool) instead of (\*RWops, int)
+ `LoadWAV()` requires string instead of (string ,\*AudioSpec)
+ `LoadWAVRW()` requires (\*RWops, bool) instead of (\*RWops, bool ,\*AudioSpec)
+ `NumHaptics()` returns (int, error) instead of int
+ `Renderer.Destroy()` returns error
+ `Renderer.GetViewport()` and `Renderer.GetClipRect()` now returns Rect instead of being passed a \*Rect
+ `ShowCursor()` returns (int, error) instead of int
+ `Surface.SaveBMPRW()` requires (\*RWops, bool) instead of (\*RWops, int)
+ `Surface.SetColorKey()` requires (bool, uint32) instead of (int, uint32)
+ `Surface.SetRLE()` requires bool instead of int
+ `Texture.Destroy()` returns error
+ `Window.Destroy()` returns error
+ `Window.GetDisplayMode()` returns (DisplayMode, error) instead of error
+ `Window.GetID()` returns (uint32, error) instead of uint32
+ Changed Mutex, Sem, Cond to have methods instead of functions
+ Changed `GameControllerMapping()` into `GameController.Mapping()`
+ Haptic functions now return bool and/or error instead of int
+ Merged `KeyUpEvent` and `KeyDownEvent` into `KeyboardEvent`
+ Moved error to the last return value in `ShowMessageBox()`
+ Split `JoyDeviceEvent` into `JoyDeviceAddedEvent` and `JoyDeviceRemovedEvent`, since Which field can contain different types of data (int or JoystickID).

- Unexported `Flags`, `Locked` and `LockData` in `Surface` struct
- Unexported `Padding` in `AudioSpec` struct
- Unexported `Rloss`, `Gloss`, `Bloss`, `Aloss`, `Rshift`, `Gshift`, `Bshift`, `Ashift`, `RefCount`, `Next` in `PixelFormat` struct
- Unexported `Version`, `RefCount` in `Palette` struct
- Unexported `goHintCallback`
