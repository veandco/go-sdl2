package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))
#pragma message("SDL_QueueAudio is not supported before SDL 2.0.4")
static int SDL_QueueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len)
{
	return -1;
}
static Uint32 SDL_GetQueuedAudioSize(SDL_AudioDeviceID dev_id)
{
	return 0;
}
static void SDL_ClearQueuedAudio(SDL_AudioDeviceID dev)
{
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))
#pragma message("SDL_DequeueAudio is not supported before SDL 2.0.5")
static int SDL_DequeueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len)
{
	return -1;
}
#endif
*/
import "C"
import (
	"reflect"
	"unsafe"
)

const (
	AUDIO_MASK_BITSIZE  = C.SDL_AUDIO_MASK_BITSIZE
	AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE
	AUDIO_MASK_ENDIAN   = C.SDL_AUDIO_MASK_ENDIAN
	AUDIO_MASK_SIGNED   = C.SDL_AUDIO_MASK_SIGNED
)

const (
	AUDIO_U8     = C.AUDIO_U8
	AUDIO_S8     = C.AUDIO_S8
	AUDIO_U16LSB = C.AUDIO_U16LSB
	AUDIO_S16LSB = C.AUDIO_S16LSB
	AUDIO_U16MSB = C.AUDIO_U16MSB
	AUDIO_S16MSB = C.AUDIO_S16MSB
	AUDIO_U16    = C.AUDIO_U16
	AUDIO_S16    = C.AUDIO_S16
	AUDIO_S32LSB = C.AUDIO_S32LSB
	AUDIO_S32MSB = C.AUDIO_S32MSB
	AUDIO_S32    = C.AUDIO_S32
	AUDIO_F32LSB = C.AUDIO_F32LSB
	AUDIO_F32MSB = C.AUDIO_F32MSB
	AUDIO_F32    = C.AUDIO_F32
	AUDIO_U16SYS = C.AUDIO_U16SYS
	AUDIO_S16SYS = C.AUDIO_S16SYS
	AUDIO_S32SYS = C.AUDIO_S32SYS
	AUDIO_F32SYS = C.AUDIO_F32SYS
)

const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE    = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE  = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE       = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

const (
	AUDIO_STOPPED AudioStatus = C.SDL_AUDIO_STOPPED
	AUDIO_PLAYING             = C.SDL_AUDIO_PLAYING
	AUDIO_PAUSED              = C.SDL_AUDIO_PAUSED
)

const MIX_MAXVOLUME = C.SDL_MIX_MAXVOLUME

// AudioFormat (https://wiki.libsdl.org/SDL_AudioFormat)
type AudioFormat uint16
type AudioCallback C.SDL_AudioCallback
type AudioFilter C.SDL_AudioFilter
type AudioDeviceID uint32

// AudioStatus (https://wiki.libsdl.org/SDL_AudioStatus)
type AudioStatus uint32
type cAudioStatus C.SDL_AudioStatus

// AudioSpec (https://wiki.libsdl.org/SDL_AudioSpec)
type AudioSpec struct {
	Freq     int32
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
	Callback AudioCallback
	UserData unsafe.Pointer
}
type cAudioSpec C.SDL_AudioSpec

// AudioCVT (https://wiki.libsdl.org/SDL_AudioCVT)
type AudioCVT struct {
	Needed      int32
	SrcFormat   AudioFormat
	DstFormat   AudioFormat
	RateIncr    float64
	Buf         unsafe.Pointer // use AudioCVT.BufAsSlice() for access via Go slice
	Len         int32
	LenCVT      int32
	LenMult     int32
	LenRatio    float64
	filters     [10]AudioFilter // internal
	filterIndex int32           // internal
}
type cAudioCVT C.SDL_AudioCVT

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

func (cvt *AudioCVT) AllocBuf(size uintptr) {
	cvt.Buf = C.malloc(C.size_t(size))
}

func (cvt *AudioCVT) FreeBuf() {
	C.free(cvt.Buf)
}

// Access AudioCVT.buf as slice.
// NOTE: Must have used ConvertAudio() before usage because it uses LenCVT as slice length.
func (cvt AudioCVT) BufAsSlice() []byte {
	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = int(cvt.LenCVT)
	sliceHeader.Cap = int(cvt.Len * cvt.LenMult)
	sliceHeader.Data = uintptr(unsafe.Pointer(cvt.Buf))
	return b
}

// GetNumAudioDrivers (https://wiki.libsdl.org/SDL_GetNumAudioDrivers)
func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

// GetAudioDriver (https://wiki.libsdl.org/SDL_GetAudioDriver)
func GetAudioDriver(index int) string {
	return string(C.GoString(C.SDL_GetAudioDriver(C.int(index))))
}

// AudioInit (https://wiki.libsdl.org/SDL_AudioInit)
func AudioInit(driverName string) error {
	_driverName := C.CString(driverName)
	defer C.free(unsafe.Pointer(_driverName))
	if C.SDL_AudioInit(_driverName) != 0 {
		return GetError()
	}
	return nil
}

// AudioQuit (https://wiki.libsdl.org/SDL_AudioQuit)
func AudioQuit() {
	C.SDL_AudioQuit()
}

// GetCurrentAudioDriver (https://wiki.libsdl.org/SDL_GetCurrentAudioDriver)
func GetCurrentAudioDriver() string {
	return string(C.GoString(C.SDL_GetCurrentAudioDriver()))
}

// OpenAudio (https://wiki.libsdl.org/SDL_OpenAudio)
func OpenAudio(desired, obtained *AudioSpec) error {
	if C.SDL_OpenAudio(desired.cptr(), obtained.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetNumAudioDevices (https://wiki.libsdl.org/SDL_GetNumAudioDevices)
func GetNumAudioDevices(isCapture bool) int {
	return int(C.SDL_GetNumAudioDevices(C.int(Btoi(isCapture))))
}

// GetAudioDeviceName (https://wiki.libsdl.org/SDL_GetAudioDeviceName)
func GetAudioDeviceName(index int, isCapture bool) string {
	return string(C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(Btoi(isCapture)))))
}

// OpenAudioDevice (https://wiki.libsdl.org/SDL_OpenAudioDevice)
func OpenAudioDevice(device string, isCapture bool, desired, obtained *AudioSpec, allowedChanges int) (AudioDeviceID, error) {
	_device := C.CString(device)
	if device == "" {
		_device = nil
	}
	defer C.free(unsafe.Pointer(_device))
	if id := AudioDeviceID(C.SDL_OpenAudioDevice(_device, C.int(Btoi(isCapture)), desired.cptr(), obtained.cptr(), C.int(allowedChanges))); id > 0 {
		return id, nil
	}
	return 0, GetError()
}

// GetAudioStatus (https://wiki.libsdl.org/SDL_GetAudioStatus)
func GetAudioStatus() AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioStatus())
}

// GetAudioDeviceStatus (https://wiki.libsdl.org/SDL_GetAudioDeviceStatus)
func GetAudioDeviceStatus(dev AudioDeviceID) AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioDeviceStatus(dev.c()))
}

// PauseAudio (https://wiki.libsdl.org/SDL_PauseAudio)
func PauseAudio(pauseOn bool) {
	C.SDL_PauseAudio(C.int(Btoi(pauseOn)))
}

// PauseAudioDevice (https://wiki.libsdl.org/SDL_PauseAudioDevice)
func PauseAudioDevice(dev AudioDeviceID, pauseOn bool) {
	C.SDL_PauseAudioDevice(dev.c(), C.int(Btoi(pauseOn)))
}

// LoadWAV_RW (https://wiki.libsdl.org/SDL_LoadWAV_RW)
func LoadWAV_RW(src *RWops, freeSrc bool, spec *AudioSpec) ([]byte, *AudioSpec) {
	var _audioBuf *C.Uint8
	var _audioLen C.Uint32
	audioSpec := (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(src.cptr(), C.int(Btoi(freeSrc)), spec.cptr(), &_audioBuf, &_audioLen)))

	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = (int)(_audioLen)
	sliceHeader.Cap = (int)(_audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(_audioBuf))
	return b, audioSpec
}

// LoadWAV (https://wiki.libsdl.org/SDL_LoadWAV)
func LoadWAV(file string, spec *AudioSpec) ([]byte, *AudioSpec) {
	_file := C.CString(file)
	_rb := C.CString("rb")
	defer C.free(unsafe.Pointer(_file))
	defer C.free(unsafe.Pointer(_rb))

	var _audioBuf *C.Uint8
	var _audioLen C.Uint32
	audioSpec := (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(C.SDL_RWFromFile(_file, _rb), 1, spec.cptr(), &_audioBuf, &_audioLen)))

	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = (int)(_audioLen)
	sliceHeader.Cap = (int)(_audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(_audioBuf))
	return b, audioSpec
}

// FreeWAV (https://wiki.libsdl.org/SDL_FreeWAV)
func FreeWAV(audioBuf []uint8) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&audioBuf))
	_audioBuf := (*C.Uint8)(unsafe.Pointer(sliceHeader.Data))
	C.SDL_FreeWAV(_audioBuf)
}

// BuildAudioCVT (https://wiki.libsdl.org/SDL_BuildAudioCVT)
func BuildAudioCVT(cvt *AudioCVT, srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) (converted bool, err error) {
	switch int(C.SDL_BuildAudioCVT(cvt.cptr(), srcFormat.c(), C.Uint8(srcChannels), C.int(srcRate), dstFormat.c(), C.Uint8(dstChannels), C.int(dstRate))) {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, GetError()
}

// ConvertAudio (https://wiki.libsdl.org/SDL_ConvertAudio)
func ConvertAudio(cvt *AudioCVT) error {
	_cvt := (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
	if C.SDL_ConvertAudio(_cvt) != 0 {
		return GetError()
	}
	return nil
}

// QueueAudio (https://wiki.libsdl.org/SDL_QueueAudio)
func QueueAudio(dev AudioDeviceID, data []byte) error {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	_data := unsafe.Pointer(sliceHeader.Data)
	_len := (C.Uint32)(sliceHeader.Len)
	if C.SDL_QueueAudio(dev.c(), _data, _len) != 0 {
		return GetError()
	}
	return nil
}

// DequeueAudio (https://wiki.libsdl.org/SDL_DequeueAudio)
func DequeueAudio(dev AudioDeviceID, data []byte) error {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	_data := unsafe.Pointer(sliceHeader.Data)
	_len := (C.Uint32)(sliceHeader.Len)
	if C.SDL_DequeueAudio(dev.c(), _data, _len) != 0 {
		return GetError()
	}
	return nil
}

// GetQueuedAudioSize (https://wiki.libsdl.org/SDL_GetQueuedAudioSize)
func GetQueuedAudioSize(dev AudioDeviceID) uint32 {
	return uint32(C.SDL_GetQueuedAudioSize(dev.c()))
}

// ClearQueuedAudio (https://wiki.libsdl.org/SDL_ClearQueuedAudio)
func ClearQueuedAudio(dev AudioDeviceID) {
	C.SDL_ClearQueuedAudio(dev.c())
}

// MixAudio (https://wiki.libsdl.org/SDL_MixAudio)
func MixAudio(dst, src *uint8, len_ uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudio(_dst, _src, C.Uint32(len_), C.int(volume))
}

// MixAudioFormat (https://wiki.libsdl.org/SDL_MixAudioFormat)
func MixAudioFormat(dst, src *uint8, format AudioFormat, len_ uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudioFormat(_dst, _src, format.c(), C.Uint32(len_), C.int(volume))
}

// LockAudio (https://wiki.libsdl.org/SDL_LockAudio)
func LockAudio() {
	C.SDL_LockAudio()
}

// LockAudioDevice (https://wiki.libsdl.org/SDL_LockAudioDevice)
func LockAudioDevice(dev AudioDeviceID) {
	C.SDL_LockAudioDevice(dev.c())
}

// UnlockAudio (https://wiki.libsdl.org/SDL_UnlockAudio)
func UnlockAudio() {
	C.SDL_UnlockAudio()
}

// UnlockAudioDevice (https://wiki.libsdl.org/SDL_UnlockAudioDevice)
func UnlockAudioDevice(dev AudioDeviceID) {
	C.SDL_UnlockAudioDevice(dev.c())
}

// CloseAudio (https://wiki.libsdl.org/SDL_CloseAudio)
func CloseAudio() {
	C.SDL_CloseAudio()
}

// CloseAudioDevice (https://wiki.libsdl.org/SDL_CloseAudioDevice)
func CloseAudioDevice(dev AudioDeviceID) {
	C.SDL_CloseAudioDevice(dev.c())
}

/*
func AudioDeviceConnected(dev AudioDeviceID) int {
	_dev := (C.SDL_AudioDeviceID) (dev)
	return int (C.SDL_AudioDeviceConnected(_dev))
}
*/
