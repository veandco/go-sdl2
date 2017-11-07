## BREAKING CHANGES

### v0.2..master

+ Renamed `Texture.GL_BindTexture()` and `Texture.GL_UnbindTexture()` to `Texture.GLBind()` and `Texture.GLUnbind()` respectively
+ Renamed `LoadWAV_RW()` to `LoadWAVRW()`
+ Renamed `TouchId` to `TouchID` in `MultiGestureEvent` struct
+ Renamed `Unicode` to `unused` in `Keysym` struct (must have been a typo)

- Unexported `Padding` in  `AudioSpec` struct
- Unexported `goHintCallback`