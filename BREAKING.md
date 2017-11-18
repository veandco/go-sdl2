## BREAKING CHANGES

### v0.2..master

+ Renamed `Texture.GL_BindTexture()` and `Texture.GL_UnbindTexture()` to `Texture.GLBind()` and `Texture.GLUnbind()` respectively
+ Renamed `LoadWAV_RW()` to `LoadWAVRW()`
+ Renamed `TouchId` to `TouchID` in `MultiGestureEvent` struct
+ Renamed `Unicode` to `unused` in `Keysym` struct (must have been a typo)
+ Renamed `ButtonId` to `ButtonID` in `MessageBoxButtonData` struct
+ Renamed `GL_LoadLibrary()` to `GLLoadLibrary()`
+ Renamed `GL_GetProcAddress()` to `GLGetProcAddress()`
+ Renamed `GL_UnloadLibrary()` to `GLUnloadLibrary()`
+ Renamed `GL_ExtensionSupported()` to `GLExtensionSupported()`
+ Renamed `GL_SetAttribute()` to `GLSetAttribute()`
+ Renamed `GL_GetAttribute()` to `GLGetAttribute()`
+ Renamed `GL_CreateContext()` to `GLCreateContext()`
+ Renamed `GL_MakeCurrent()` to `GLMakeCurrent()`
+ Renamed `GL_SetSwapInterval()` to `GLSetSwapInterval()`
+ Renamed `GL_GetSwapInterval()` to `GLGetSwapInterval()`
+ Renamed `GL_GetDrawableSize()` to `GLGetDrawableSize()`
+ Renamed `GL_SwapWindow()` to `GLSwapWindow()`
+ Renamed `GL_DeleteContext()` to `GLDeleteContext()`
+ Renamed `LoadBMP_RW()` to `LoadBMPRW()`
+ Renamed `SaveBMP_RW()` to `SaveBMPRW()`
+ Renamed `Id` to `ID` in `Finger` struct
+ Renamed `SW_YUVTexture` to `SWYUVTexture`
+ Moved error to the last return value in `ShowMessageBox()`
+ Change Mutex, Sem, Cond to have methods instead of functions
+ Merge `KeyUpEvent` and `KeyDownEvent` into `KeyboardEvent`

- Unexported `Padding` in `AudioSpec` struct
- Unexported `goHintCallback`
- Unexported `Flags`, `Locked` and `LockData` in `Surface` struct
- Unexported `Rloss`, `Gloss`, `Bloss`, `Aloss`, `Rshift`, `Gshift`, `Bshift`, `Ashift`, `RefCount`, `Next` in `PixelFormat` struct
- Unexported `Version`, `RefCount` in `Palette` struct