package sdl

/*
#include <SDL2/SDL.h>
*/
import "C"
import "unsafe"

const (
	AUDIO_MASK_BITSIZE       = 0xFF
	AUDIO_MASK_DATATYPE      = 1 << 8
	AUDIO_MASK_ENDIAN        = 1 << 12
	AUDIO_MASK_SIGNED        = 1 << 15
)

const (
	AUDIO_U8		 = 0x0008
	AUDIO_S8		 = 0x8008
	AUDIO_U16LSB		 = 0x0010
	AUDIO_S16LSB		 = 0x8010
	AUDIO_U16MSB		 = 0x1010
	AUDIO_S16MSB		 = 0x9010
	AUDIO_U16		 = AUDIO_U16LSB
	AUDIO_S16		 = AUDIO_S16LSB
)

const (
	AUDIO_S32LSB		 = 0x8020
	AUDIO_S32MSB		 = 0x9020
	AUDIO_S32		 = AUDIO_S32LSB
)

const (
	AUDIO_F32LSB		 = 0x8120
	AUDIO_F32MSB		 = 0x9120
	AUDIO_F32		 = AUDIO_F32LSB
)

const (
	AUDIO_U16SYS		= C.AUDIO_U16SYS
	AUDIO_S16SYS		= C.AUDIO_S16SYS
	AUDIO_S32SYS		= C.AUDIO_S32SYS
	AUDIO_F32SYS		= C.AUDIO_F32SYS
)

const (
	AUDIO_ALLOW_FREQUENCY_CHANGE	= 0x00000001
	AUDIO_ALLOW_FORMAT_CHANGE	= 0x00000002
	AUDIO_ALLOW_CHANNELS_CHANGE	= 0x00000004
	AUDIO_ALLOW_ANY_CHANGE		= (AUDIO_ALLOW_FREQUENCY_CHANGE |
					   AUDIO_ALLOW_FORMAT_CHANGE |
					   AUDIO_ALLOW_CHANNELS_CHANGE)
)

const (
	AUDIO_STOPPED			= iota
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
	Freq int
	Format AudioFormat
	Channels uint8
	Silence uint8
	Samples uint16
	Padding uint16
	Size uint32
	Callback AudioCallback
	UserData unsafe.Pointer
}

type AudioCVT struct {
	Needed int
	SrcFormat AudioFormat
	DstFormat AudioFormat
	RateIncr float64
	Buf	*uint8
	Len int
	LenCVT int
	LenMult int
	LenRatio float64
	filters [10]AudioFilter
	FilterIndex int
}

func (format AudioFormat) BitSize() uint8 {
	return (uint8) (format & AUDIO_MASK_BITSIZE)
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
	return (int) (C.SDL_GetNumAudioDrivers())
}

func GetAudioDriver(index int) string {
	_index := (C.int) (index)
	return (string) (C.GoString(C.SDL_GetAudioDriver(_index)))
}

func AudioInit(driver_name string) int {
	_driver_name := (C.CString) (driver_name)
	return (int) (C.SDL_AudioInit(_driver_name))
}

func AudioQuit() {
	C.SDL_AudioQuit()
}

func GetCurrentAudioDriver() string {
	return (string) (C.GoString(C.SDL_GetCurrentAudioDriver()))
}

func OpenAudio(desired, obtained *AudioSpec) int {
	_desired := (*C.SDL_AudioSpec) (unsafe.Pointer(desired))
	_obtained := (*C.SDL_AudioSpec) (unsafe.Pointer(obtained))
	return (int) (C.SDL_OpenAudio(_desired, _obtained))
}

func GetNumAudioDevices(iscapture int) int {
	_iscapture := (C.int) (iscapture)
	return (int) (C.SDL_GetNumAudioDevices(_iscapture))
}

func GetAudioDeviceName(index, iscapture int) string {
	_index := (C.int) (index)
	_iscapture := (C.int) (iscapture)
	return (string) (C.GoString(C.SDL_GetAudioDeviceName(_index, _iscapture)))
}

func OpenAudioDevice(device string, iscapture int, desired, obtained *AudioSpec, allowed_changes int) int {
	_device := (C.CString) (device)
	_iscapture := (C.int) (iscapture)
	_desired := (*C.SDL_AudioSpec) (unsafe.Pointer(desired))
	_obtained := (*C.SDL_AudioSpec) (unsafe.Pointer(obtained))
	_allowed_changes := (C.int) (allowed_changes)
	return (int) (C.SDL_OpenAudioDevice(_device, _iscapture, _desired, _obtained, _allowed_changes))
}

func GetAudioStatus() AudioStatus {
	return (AudioStatus) (C.SDL_GetAudioStatus())
}


func GetAudioDeviceStatus(dev AudioDeviceID) AudioStatus {
	_dev := (C.SDL_AudioDeviceID) (dev)
	return (AudioStatus) (C.SDL_GetAudioDeviceStatus(_dev))
}

func PauseAudio(pause_on int) {
	_pause_on := (C.int) (pause_on)
	C.SDL_PauseAudio(_pause_on)
}

func PauseAudioDevice(dev AudioDeviceID, pause_on int) {
	_dev := (C.SDL_AudioDeviceID) (dev)
	_pause_on := (C.int) (pause_on)
	C.SDL_PauseAudioDevice(_dev, _pause_on)
}

func LoadWAV_RW(src *RWops, freesrc int, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) *AudioSpec {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	_freesrc := (C.int) (freesrc)
	_spec := (*C.SDL_AudioSpec) (unsafe.Pointer(spec))
	_audio_buf := (**C.Uint8) (unsafe.Pointer(audio_buf))
	_audio_len := (*C.Uint32) (unsafe.Pointer(audio_len))
	return (*AudioSpec) (unsafe.Pointer(C.SDL_LoadWAV_RW(_src, _freesrc, _spec, _audio_buf, _audio_len)))
}

func LoadWAV(file string, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) *AudioSpec {
	_file := (C.CString) (file)
	_spec := (*C.SDL_AudioSpec) (unsafe.Pointer(spec))
	_audio_buf := (**C.Uint8) (unsafe.Pointer(audio_buf))
	_audio_len := (*C.Uint32) (unsafe.Pointer(audio_len))
	return (*AudioSpec) (unsafe.Pointer(C.SDL_LoadWAV_RW(C.SDL_RWFromFile(_file, C.CString("rb")),
					1, _spec, _audio_buf, _audio_len)))
}

func FreeWAV(audio_buf *uint8) {
	_audio_buf := (*C.Uint8) (unsafe.Pointer(audio_buf))
	C.SDL_FreeWAV(_audio_buf)
}

func BuildAudioCVT(cvt *AudioCVT, src_format AudioFormat, src_channels uint8, src_rate int,
		dst_format AudioFormat, dst_channels uint8, dst_rate int) int {
	_cvt := (*C.SDL_AudioCVT) (unsafe.Pointer(cvt))
	_src_format := (C.SDL_AudioFormat) (src_format)
	_src_channels := (C.Uint8) (src_channels)
	_src_rate := (C.int) (src_rate)
	_dst_format := (C.SDL_AudioFormat) (dst_format)
	_dst_channels := (C.Uint8) (dst_channels)
	_dst_rate := (C.int) (dst_rate)
	return (int) (C.SDL_BuildAudioCVT(_cvt, _src_format, _src_channels, _src_rate,
					_dst_format, _dst_channels, _dst_rate))
}

func ConvertAudio(cvt *AudioCVT) int {
	_cvt := (*C.SDL_AudioCVT) (unsafe.Pointer(cvt))
	return (int) (C.SDL_ConvertAudio(_cvt))
}

func MixAudio(dst, src *uint8, len_ uint32, volume int) {
	_dst := (*C.Uint8) (unsafe.Pointer(dst))
	_src := (*C.Uint8) (unsafe.Pointer(src))
	_len := (C.Uint32) (len_)
	_volume := (C.int) (volume)
	C.SDL_MixAudio(_dst, _src, _len, _volume)
}

func MixAudioFormat(dst, src *uint8, format AudioFormat, len_ uint32, volume int) {
	_dst := (*C.Uint8) (unsafe.Pointer(dst))
	_src := (*C.Uint8) (unsafe.Pointer(src))
	_format := (C.SDL_AudioFormat) (format)
	_len := (C.Uint32) (len_)
	_volume := (C.int) (volume)
	C.SDL_MixAudioFormat(_dst, _src, _format, _len, _volume)
}

func LockAudio() {
	C.SDL_LockAudio();
}

func LockAudioDevice(dev AudioDeviceID) {
	_dev := (C.SDL_AudioDeviceID) (dev)
	C.SDL_LockAudioDevice(_dev);
}

func UnlockAudio() {
	C.SDL_UnlockAudio();
}

func UnlockAudioDevice(dev AudioDeviceID) {
	_dev := (C.SDL_AudioDeviceID) (dev)
	C.SDL_UnlockAudioDevice(_dev);
}

func CloseAudio() {
	C.SDL_UnlockAudio();
}

func CloseAudioDevice(dev AudioDeviceID) {
	_dev := (C.SDL_AudioDeviceID) (dev)
	C.SDL_UnlockAudioDevice(_dev);
}

/*
func AudioDeviceConnected(dev AudioDeviceID) int {
	_dev := (C.SDL_AudioDeviceID) (dev)
	return (int) (C.SDL_AudioDeviceConnected(_dev))
}
*/
