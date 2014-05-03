package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	AUDIO_MASK_BITSIZE  = 0xFF
	AUDIO_MASK_DATATYPE = 1 << 8
	AUDIO_MASK_ENDIAN   = 1 << 12
	AUDIO_MASK_SIGNED   = 1 << 15
)

const (
	AUDIO_U8     = 0x0008
	AUDIO_S8     = 0x8008
	AUDIO_U16LSB = 0x0010
	AUDIO_S16LSB = 0x8010
	AUDIO_U16MSB = 0x1010
	AUDIO_S16MSB = 0x9010
	AUDIO_U16    = AUDIO_U16LSB
	AUDIO_S16    = AUDIO_S16LSB
)

const (
	AUDIO_S32LSB = 0x8020
	AUDIO_S32MSB = 0x9020
	AUDIO_S32    = AUDIO_S32LSB
)

const (
	AUDIO_F32LSB = 0x8120
	AUDIO_F32MSB = 0x9120
	AUDIO_F32    = AUDIO_F32LSB
)

const (
	AUDIO_U16SYS = C.AUDIO_U16SYS
	AUDIO_S16SYS = C.AUDIO_S16SYS
	AUDIO_S32SYS = C.AUDIO_S32SYS
	AUDIO_F32SYS = C.AUDIO_F32SYS
)

const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = 0x00000001
	AUDIO_ALLOW_FORMAT_CHANGE    = 0x00000002
	AUDIO_ALLOW_CHANNELS_CHANGE  = 0x00000004
	AUDIO_ALLOW_ANY_CHANGE       = (AUDIO_ALLOW_FREQUENCY_CHANGE |
		AUDIO_ALLOW_FORMAT_CHANGE |
		AUDIO_ALLOW_CHANNELS_CHANGE)
)

const (
	AUDIO_STOPPED = iota
	AUDIO_PLAYING
	AUDIO_PAUSED
)

const MIX_MAXVOLUME = 128

type AudioFormat uint16
type AudioCallback C.SDL_AudioCallback
type AudioFilter C.SDL_AudioFilter
type AudioDeviceID uint32
type AudioStatus uint

type AudioSpec struct {
	Freq     int
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
	Callback AudioCallback
	UserData unsafe.Pointer
}

type AudioCVT struct {
	Needed      int
	SrcFormat   AudioFormat
	DstFormat   AudioFormat
	RateIncr    float64
	Buf         *uint8
	Len         int
	LenCVT      int
	LenMult     int
	LenRatio    float64
	filters     [10]AudioFilter
	FilterIndex int
}

func (fmt AudioFormat) c() C.SDL_AudioFormat {
    return C.SDL_AudioFormat(fmt)
}

func (id AudioDeviceID) c() C.SDL_AudioDeviceID {
    return C.SDL_AudioDeviceID(id)
}

func (as *AudioSpec) cptr() *C.SDL_AudioSpec {
    return (*C.SDL_AudioSpec)(unsafe.Pointer(as))
}

func (cvt *AudioCVT) cptr() *C.SDL_AudioCVT {
    return (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
}

func (format AudioFormat) BitSize() uint8 {
	return uint8(format & AUDIO_MASK_BITSIZE)
}

func (format AudioFormat) IsFloat() bool {
	return (format & AUDIO_MASK_DATATYPE) > 0
}

func (format AudioFormat) IsBigEndian() bool {
	return (format & AUDIO_MASK_ENDIAN) > 0
}

func (format AudioFormat) IsSigned() bool {
	return (format & AUDIO_MASK_SIGNED) > 0
}

func (format AudioFormat) IsInt() bool {
	return !format.IsFloat()
}

func (format AudioFormat) IsLittleEndian() bool {
	return !format.IsBigEndian()
}

func (format AudioFormat) IsUnsigned() bool {
	return !format.IsSigned()
}

func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

func GetAudioDriver(index int) string {
	return string(C.GoString(C.SDL_GetAudioDriver(C.int(index))))
}

func AudioInit(driverName string) int {
	_driverName := C.CString(driverName)
	defer C.free(unsafe.Pointer(_driverName))
	return int(C.SDL_AudioInit(_driverName))
}

func AudioQuit() {
	C.SDL_AudioQuit()
}

func GetCurrentAudioDriver() string {
	return string(C.GoString(C.SDL_GetCurrentAudioDriver()))
}

func OpenAudio(desired, obtained *AudioSpec) int {
	return int(C.SDL_OpenAudio(desired.cptr(), obtained.cptr()))
}

func GetNumAudioDevices(isCapture int) int {
	return int(C.SDL_GetNumAudioDevices(C.int(isCapture)))
}

func GetAudioDeviceName(index, isCapture int) string {
	return string(C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(isCapture))))
}

func OpenAudioDevice(device string, isCapture int, desired, obtained *AudioSpec, allowedChanges int) int {
	_device := C.CString(device)
	defer C.free(unsafe.Pointer(_device))
	return int(C.SDL_OpenAudioDevice(_device, C.int(isCapture), desired.cptr(), obtained.cptr(), C.int(allowedChanges)))
}

func GetAudioStatus() AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioStatus())
}

func GetAudioDeviceStatus(dev AudioDeviceID) AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioDeviceStatus(dev.c()))
}

func PauseAudio(pauseOn int) {
	C.SDL_PauseAudio(C.int(pauseOn))
}

func PauseAudioDevice(dev AudioDeviceID, pauseOn int) {
	C.SDL_PauseAudioDevice(dev.c(), C.int(pauseOn))
}

func LoadWAV_RW(src *RWops, freeSrc int, spec *AudioSpec, audioBuf **uint8, audioLen *uint32) *AudioSpec {
	_audioBuf := (**C.Uint8)(unsafe.Pointer(audioBuf))
	_audioLen := (*C.Uint32)(unsafe.Pointer(audioLen))
	return (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(src.cptr(), C.int(freeSrc), spec.cptr(), _audioBuf, _audioLen)))
}

func LoadWAV(file string, spec *AudioSpec, audioBuf **uint8, audioLen *uint32) *AudioSpec {
	_file := C.CString(file)
	_rb := C.CString("rb")
	defer C.free(unsafe.Pointer(_file))
	defer C.free(unsafe.Pointer(_rb))
	_audioBuf := (**C.Uint8)(unsafe.Pointer(audioBuf))
	_audioLen := (*C.Uint32)(unsafe.Pointer(audioLen))
	return (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(C.SDL_RWFromFile(_file, _rb), 1, spec.cptr(), _audioBuf, _audioLen)))
}

func FreeWAV(audioBuf *uint8) {
	_audioBuf := (*C.Uint8)(unsafe.Pointer(audioBuf))
	C.SDL_FreeWAV(_audioBuf)
}

func BuildAudioCVT(cvt *AudioCVT, srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) int {
	return int(C.SDL_BuildAudioCVT(cvt.cptr(), srcFormat.c(), C.Uint8(srcChannels), C.int(srcRate), dstFormat.c(), C.Uint8(dstChannels), C.int(dstRate)))
}

func ConvertAudio(cvt *AudioCVT) int {
	_cvt := (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
	return int(C.SDL_ConvertAudio(_cvt))
}

func MixAudio(dst, src *uint8, len_ uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudio(_dst, _src, C.Uint32(len_), C.int(volume))
}

func MixAudioFormat(dst, src *uint8, format AudioFormat, len_ uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudioFormat(_dst, _src, format.c(), C.Uint32(len_), C.int(volume))
}

func LockAudio() {
	C.SDL_LockAudio()
}

func LockAudioDevice(dev AudioDeviceID) {
	C.SDL_LockAudioDevice(dev.c())
}

func UnlockAudio() {
	C.SDL_UnlockAudio()
}

func UnlockAudioDevice(dev AudioDeviceID) {
	C.SDL_UnlockAudioDevice(dev.c())
}

func CloseAudio() {
	C.SDL_UnlockAudio()
}

func CloseAudioDevice(dev AudioDeviceID) {
	C.SDL_UnlockAudioDevice(dev.c())
}

/*
func AudioDeviceConnected(dev AudioDeviceID) int {
	_dev := (C.SDL_AudioDeviceID) (dev)
	return int (C.SDL_AudioDeviceConnected(_dev))
}
*/
