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
+ Renamed `LoadTyped_RW` to `LoadTypedRW`
+ Renamed `Load_RW` to `LoadRW`
+ Renamed `LoadTexture_RW` to `LoadTextureRW`
+ Renamed `LoadICO_RW()` to `LoadICORW()`
+ Renamed `LoadCUR_RW()` to `LoadCURRW()`
+ Renamed `LoadBMP_RW()` to `LoadBMPRW()`
+ Renamed `LoadGIF_RW()` to `LoadGIFRW()`
+ Renamed `LoadJPG_RW()` to `LoadJPGRW()`
+ Renamed `LoadLBM_RW()` to `LoadLBMRW()`
+ Renamed `LoadPCX_RW()` to `LoadPCXRW()`
+ Renamed `LoadPNG_RW()` to `LoadPNGRW()`
+ Renamed `LoadPNM_RW()` to `LoadPNMRW()`
+ Renamed `LoadTGA_RW()` to `LoadTGARW()`
+ Renamed `LoadTIF_RW()` to `LoadTIFRW()`
+ Renamed `LoadXCF_RW()` to `LoadXCFRW()`
+ Renamed `LoadXPM_RW()` to `LoadXPMRW()`
+ Renamed `LoadXV_RW()` to `LoadXVRW()`
+ Renamed `LoadWEBP_RW()` to `LoadWEBPRW()`
+ Renamed `SavePNG_RW()` to `SavePNGRW()`
+ Moved error to the last return value in `ShowMessageBox()`
+ Change Mutex, Sem, Cond to have methods instead of functions

- Unexported `Padding` in `AudioSpec` struct
- Unexported `goHintCallback`
- Unexported `Flags`, `Locked` and `LockData` in `Surface` struct
- Unexported `Rloss`, `Gloss`, `Bloss`, `Aloss`, `Rshift`, `Gshift`, `Bshift`, `Ashift`, `RefCount`, `Next` in `PixelFormat` struct
- Unexported `Version`, `RefCount` in `Palette` struct