package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_QueueAudio is not supported before SDL 2.0.4")
#endif

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

#if defined(WARN_OUTDATED)
#pragma message("SDL_DequeueAudio is not supported before SDL 2.0.5")
#endif

static int SDL_DequeueAudio(SDL_AudioDeviceID dev, const void *data, Uint32 len)
{
	return -1;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,7))

struct _SDL_AudioStream;
typedef struct _SDL_AudioStream SDL_AudioStream;


#if defined(WARN_OUTDATED)
#pragma message("SDL_NewAudioStream is not supported before SDL 2.0.7")
#endif

static SDL_AudioStream * SDL_NewAudioStream(const SDL_AudioFormat src_format, const Uint8 src_channels, const int src_rate, const SDL_AudioFormat dst_format, const Uint8 dst_channels, const int dst_rate)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioStreamPut is not supported before SDL 2.0.7")
#endif

static int SDL_AudioStreamPut(SDL_AudioStream *stream, const void *buf, int len)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioStreamGet is not supported before SDL 2.0.7")
#endif

static int SDL_AudioStreamGet(SDL_AudioStream *stream, void *buf, int len)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioStreamAvailable is not supported before SDL 2.0.7")
#endif

static int SDL_AudioStreamAvailable(SDL_AudioStream *stream)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioStreamFlush is not supported before SDL 2.0.7")
#endif

static int SDL_AudioStreamFlush(SDL_AudioStream *stream)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioStreamClear is not supported before SDL 2.0.7")
#endif

static int SDL_AudioStreamClear(SDL_AudioStream *stream)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_FreeAudioStream is not supported before SDL 2.0.7")
#endif

static void SDL_FreeAudioStream(SDL_AudioStream *stream)
{
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,16))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetAudioDeviceSpec is not supported before SDL 2.0.16")
#endif

static int SDL_GetAudioDeviceSpec(int index, int iscapture, SDL_AudioSpec *spec)
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

// Audio format masks.
// (https://wiki.libsdl.org/SDL_AudioFormat)
const (
	AUDIO_MASK_BITSIZE  = C.SDL_AUDIO_MASK_BITSIZE  // (0xFF)
	AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE // (1<<8)
	AUDIO_MASK_ENDIAN   = C.SDL_AUDIO_MASK_ENDIAN   // (1<<12)
	AUDIO_MASK_SIGNED   = C.SDL_AUDIO_MASK_SIGNED   // (1<<15)
)

// Audio format values.
// (https://wiki.libsdl.org/SDL_AudioFormat)
const (
	AUDIO_S8 = C.AUDIO_S8 // signed 8-bit samples
	AUDIO_U8 = C.AUDIO_U8 // unsigned 8-bit samples

	AUDIO_S16LSB = C.AUDIO_S16LSB // signed 16-bit samples in little-endian byte order
	AUDIO_S16MSB = C.AUDIO_S16MSB // signed 16-bit samples in big-endian byte order
	AUDIO_S16SYS = C.AUDIO_S16SYS // signed 16-bit samples in native byte order
	AUDIO_S16    = C.AUDIO_S16    // AUDIO_S16LSB
	AUDIO_U16LSB = C.AUDIO_U16LSB // unsigned 16-bit samples in little-endian byte order
	AUDIO_U16MSB = C.AUDIO_U16MSB // unsigned 16-bit samples in big-endian byte order
	AUDIO_U16SYS = C.AUDIO_U16SYS // unsigned 16-bit samples in native byte order
	AUDIO_U16    = C.AUDIO_U16    // AUDIO_U16LSB

	AUDIO_S32LSB = C.AUDIO_S32LSB // 32-bit integer samples in little-endian byte order
	AUDIO_S32MSB = C.AUDIO_S32MSB // 32-bit integer samples in big-endian byte order
	AUDIO_S32SYS = C.AUDIO_S32SYS // 32-bit integer samples in native byte order
	AUDIO_S32    = C.AUDIO_S32    // AUDIO_S32LSB

	AUDIO_F32LSB = C.AUDIO_F32LSB // 32-bit floating point samples in little-endian byte order
	AUDIO_F32MSB = C.AUDIO_F32MSB // 32-bit floating point samples in big-endian byte order
	AUDIO_F32SYS = C.AUDIO_F32SYS // 32-bit floating point samples in native byte order
	AUDIO_F32    = C.AUDIO_F32    // AUDIO_F32LSB
)

// AllowedChanges flags specify how SDL should behave when a device cannot offer a specific feature. If the application requests a feature that the hardware doesn't offer, SDL will always try to get the closest equivalent. Used in OpenAudioDevice().
// (https://wiki.libsdl.org/SDL_OpenAudioDevice)
const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE    = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE  = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE       = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

// An enumeration of audio device states used in GetAudioDeviceStatus() and GetAudioStatus().
// (https://wiki.libsdl.org/SDL_AudioStatus)
const (
	AUDIO_STOPPED AudioStatus = C.SDL_AUDIO_STOPPED // audio device is stopped
	AUDIO_PLAYING             = C.SDL_AUDIO_PLAYING // audio device is playing
	AUDIO_PAUSED              = C.SDL_AUDIO_PAUSED  // audio device is paused
)

// MIX_MAXVOLUME is the full audio volume value used in MixAudioFormat() and AudioFormat().
// (https://wiki.libsdl.org/SDL_MixAudioFormat)
const MIX_MAXVOLUME = C.SDL_MIX_MAXVOLUME // full audio volume

// AudioFormat is an enumeration of audio formats.
// (https://wiki.libsdl.org/SDL_AudioFormat)
type AudioFormat uint16

// AudioCallback is a function to call when the audio device needs more data.`
// (https://wiki.libsdl.org/SDL_AudioSpec)
type AudioCallback C.SDL_AudioCallback

// AudioFilter is the filter list used in AudioCVT() (internal use)
// (https://wiki.libsdl.org/SDL_AudioCVT)
type AudioFilter C.SDL_AudioFilter

// AudioDeviceID is ID of an audio device previously opened with OpenAudioDevice().
// (https://wiki.libsdl.org/SDL_OpenAudioDevice)
type AudioDeviceID uint32

// AudioStatus is an enumeration of audio device states.
// (https://wiki.libsdl.org/SDL_AudioStatus)
type AudioStatus uint32
type cAudioStatus C.SDL_AudioStatus

// AudioSpec contains the audio output format. It also contains a callback that is called when the audio device needs more data.
// (https://wiki.libsdl.org/SDL_AudioSpec)
type AudioSpec struct {
	Freq     int32          // DSP frequency (samples per second)
	Format   AudioFormat    // audio data format
	Channels uint8          // number of separate sound channels
	Silence  uint8          // audio buffer silence value (calculated)
	Samples  uint16         // audio buffer size in samples (power of 2)
	_        uint16         // padding
	Size     uint32         // audio buffer size in bytes (calculated)
	Callback AudioCallback  // the function to call when the audio device needs more data
	UserData unsafe.Pointer // a pointer that is passed to callback (otherwise ignored by SDL)
}
type cAudioSpec C.SDL_AudioSpec

// AudioCVT contains audio data conversion information.
// (https://wiki.libsdl.org/SDL_AudioCVT)
type AudioCVT struct {
	Needed      int32           // set to 1 if conversion possible
	SrcFormat   AudioFormat     // source audio format
	DstFormat   AudioFormat     // target audio format
	RateIncr    float64         // rate conversion increment
	Buf         unsafe.Pointer  // the buffer to hold entire audio data. Use AudioCVT.BufAsSlice() for access via a Go slice
	Len         int32           // length of original audio buffer
	LenCVT      int32           // length of converted audio buffer
	LenMult     int32           // buf must be len*len_mult big
	LenRatio    float64         // given len, final size is len*len_ratio
	filters     [10]AudioFilter // filter list (internal use)
	filterIndex int32           // current audio conversion function (internal use)
}
type cAudioCVT C.SDL_AudioCVT

// AudioStream is a new audio conversion interface.
// (https://wiki.libsdl.org/SDL_AudioStream)
type AudioStream C.SDL_AudioStream

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

func (stream *AudioStream) cptr() *C.SDL_AudioStream {
	return (*C.SDL_AudioStream)(unsafe.Pointer(stream))
}

// BitSize returns audio formats bit size.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) BitSize() uint8 {
	return uint8(fmt & AUDIO_MASK_BITSIZE)
}

// IsFloat reports whether audio format is float.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsFloat() bool {
	return (fmt & AUDIO_MASK_DATATYPE) > 0
}

// IsBigEndian reports whether audio format is big-endian.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsBigEndian() bool {
	return (fmt & AUDIO_MASK_ENDIAN) > 0
}

// IsSigned reports whether audio format is signed.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsSigned() bool {
	return (fmt & AUDIO_MASK_SIGNED) > 0
}

// IsInt reports whether audio format is integer.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsInt() bool {
	return !fmt.IsFloat()
}

// IsLittleEndian reports whether audio format is little-endian.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsLittleEndian() bool {
	return !fmt.IsBigEndian()
}

// IsUnsigned reports whether audio format is unsigned.
// (https://wiki.libsdl.org/SDL_AudioFormat)
func (fmt AudioFormat) IsUnsigned() bool {
	return !fmt.IsSigned()
}

// AllocBuf allocates the requested memory for AudioCVT buffer.
func (cvt *AudioCVT) AllocBuf(size uintptr) {
	cvt.Buf = C.malloc(C.size_t(size))
}

// FreeBuf deallocates the memory previously allocated from AudioCVT buffer.
func (cvt *AudioCVT) FreeBuf() {
	C.free(cvt.Buf)
}

// BufAsSlice returns AudioCVT.buf as byte slice.
// NOTE: Must be used after ConvertAudio() because it uses LenCVT as slice length.
func (cvt AudioCVT) BufAsSlice() []byte {
	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = int(cvt.LenCVT)
	sliceHeader.Cap = int(cvt.Len * cvt.LenMult)
	sliceHeader.Data = uintptr(unsafe.Pointer(cvt.Buf))
	return b
}

// GetNumAudioDrivers returns the number of built-in audio drivers.
// (https://wiki.libsdl.org/SDL_GetNumAudioDrivers)
func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

// GetAudioDriver returns the name of a built in audio driver.
// (https://wiki.libsdl.org/SDL_GetAudioDriver)
func GetAudioDriver(index int) string {
	return string(C.GoString(C.SDL_GetAudioDriver(C.int(index))))
}

// AudioInit initializes a particular audio driver.
// (https://wiki.libsdl.org/SDL_AudioInit)
func AudioInit(driverName string) error {
	_driverName := C.CString(driverName)
	defer C.free(unsafe.Pointer(_driverName))
	if C.SDL_AudioInit(_driverName) != 0 {
		return GetError()
	}
	return nil
}

// AudioQuit shuts down audio if you initialized it with AudioInit().
// (https://wiki.libsdl.org/SDL_AudioQuit)
func AudioQuit() {
	C.SDL_AudioQuit()
}

// GetCurrentAudioDriver returns the name of the current audio driver.
// (https://wiki.libsdl.org/SDL_GetCurrentAudioDriver)
func GetCurrentAudioDriver() string {
	return string(C.GoString(C.SDL_GetCurrentAudioDriver()))
}

// OpenAudio opens the audio device. New programs might want to use OpenAudioDevice() instead.
// (https://wiki.libsdl.org/SDL_OpenAudio)
func OpenAudio(desired, obtained *AudioSpec) error {
	if C.SDL_OpenAudio(desired.cptr(), obtained.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetNumAudioDevices returns the number of built-in audio devices.
// (https://wiki.libsdl.org/SDL_GetNumAudioDevices)
func GetNumAudioDevices(isCapture bool) int {
	return int(C.SDL_GetNumAudioDevices(C.int(Btoi(isCapture))))
}

// GetAudioDeviceName returns the name of a specific audio device.
// (https://wiki.libsdl.org/SDL_GetAudioDeviceName)
func GetAudioDeviceName(index int, isCapture bool) string {
	return string(C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(Btoi(isCapture)))))
}

// OpenAudioDevice opens a specific audio device.
// (https://wiki.libsdl.org/SDL_OpenAudioDevice)
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

// GetAudioStatus returns the current audio state of the audio device. New programs might want to use GetAudioDeviceStatus() instead.
// (https://wiki.libsdl.org/SDL_GetAudioStatus)
func GetAudioStatus() AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioStatus())
}

// GetAudioDeviceStatus returns the current audio state of an audio device.
// (https://wiki.libsdl.org/SDL_GetAudioDeviceStatus)
func GetAudioDeviceStatus(dev AudioDeviceID) AudioStatus {
	return (AudioStatus)(C.SDL_GetAudioDeviceStatus(dev.c()))
}

// PauseAudio pauses and unpauses the audio device. New programs might want to use SDL_PauseAudioDevice() instead.
// (https://wiki.libsdl.org/SDL_PauseAudio)
func PauseAudio(pauseOn bool) {
	C.SDL_PauseAudio(C.int(Btoi(pauseOn)))
}

// PauseAudioDevice pauses and unpauses audio playback on a specified device.
// (https://wiki.libsdl.org/SDL_PauseAudioDevice)
func PauseAudioDevice(dev AudioDeviceID, pauseOn bool) {
	C.SDL_PauseAudioDevice(dev.c(), C.int(Btoi(pauseOn)))
}

// LoadWAVRW loads a WAVE from the data source, automatically freeing that source if freeSrc is true.
// (https://wiki.libsdl.org/SDL_LoadWAV_RW)
func LoadWAVRW(src *RWops, freeSrc bool) ([]byte, *AudioSpec) {
	var _audioBuf *C.Uint8
	var _audioLen C.Uint32
	audioSpec := (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(src.cptr(), C.int(Btoi(freeSrc)), (&AudioSpec{}).cptr(), &_audioBuf, &_audioLen)))

	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = (int)(_audioLen)
	sliceHeader.Cap = (int)(_audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(_audioBuf))
	return b, audioSpec
}

// LoadWAV loads a WAVE from a file.
// (https://wiki.libsdl.org/SDL_LoadWAV)
func LoadWAV(file string) ([]byte, *AudioSpec) {
	_file := C.CString(file)
	_rb := C.CString("rb")
	defer C.free(unsafe.Pointer(_file))
	defer C.free(unsafe.Pointer(_rb))

	var _audioBuf *C.Uint8
	var _audioLen C.Uint32
	audioSpec := (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(C.SDL_RWFromFile(_file, _rb), 1, (&AudioSpec{}).cptr(), &_audioBuf, &_audioLen)))

	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = (int)(_audioLen)
	sliceHeader.Cap = (int)(_audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(_audioBuf))
	return b, audioSpec
}

// FreeWAV frees data previously allocated with LoadWAV() or LoadWAVRW().
// (https://wiki.libsdl.org/SDL_FreeWAV)
func FreeWAV(audioBuf []uint8) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&audioBuf))
	_audioBuf := (*C.Uint8)(unsafe.Pointer(sliceHeader.Data))
	C.SDL_FreeWAV(_audioBuf)
}

// BuildAudioCVT initializes an AudioCVT structure for conversion.
// (https://wiki.libsdl.org/SDL_BuildAudioCVT)
func BuildAudioCVT(cvt *AudioCVT, srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) (converted bool, err error) {
	switch int(C.SDL_BuildAudioCVT(cvt.cptr(), srcFormat.c(), C.Uint8(srcChannels), C.int(srcRate), dstFormat.c(), C.Uint8(dstChannels), C.int(dstRate))) {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, GetError()
}

// ConvertAudio converts audio data to a desired audio format.
// (https://wiki.libsdl.org/SDL_ConvertAudio)
func ConvertAudio(cvt *AudioCVT) error {
	_cvt := (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
	if C.SDL_ConvertAudio(_cvt) != 0 {
		return GetError()
	}
	return nil
}

// QueueAudio queues more audio on non-callback devices.
// (https://wiki.libsdl.org/SDL_QueueAudio)
func QueueAudio(dev AudioDeviceID, data []byte) error {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	_data := unsafe.Pointer(sliceHeader.Data)
	_len := (C.Uint32)(sliceHeader.Len)
	if C.SDL_QueueAudio(dev.c(), _data, _len) != 0 {
		return GetError()
	}
	return nil
}

// DequeueAudio dequeues more audio on non-callback devices. Returns the number of bytes dequeued, which could be less than requested
// (https://wiki.libsdl.org/SDL_DequeueAudio)
func DequeueAudio(dev AudioDeviceID, data []byte) (n int, err error) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	_data := unsafe.Pointer(sliceHeader.Data)
	_len := (C.Uint32)(sliceHeader.Len)
	if dequeued := int(C.SDL_DequeueAudio(dev.c(), _data, _len)); dequeued > 0 {
		return dequeued, nil
	}
	return 0, GetError()
}

// GetQueuedAudioSize returns the number of bytes of still-queued audio.
// (https://wiki.libsdl.org/SDL_GetQueuedAudioSize)
func GetQueuedAudioSize(dev AudioDeviceID) uint32 {
	return uint32(C.SDL_GetQueuedAudioSize(dev.c()))
}

// ClearQueuedAudio drops any queued audio data waiting to be sent to the hardware.
// (https://wiki.libsdl.org/SDL_ClearQueuedAudio)
func ClearQueuedAudio(dev AudioDeviceID) {
	C.SDL_ClearQueuedAudio(dev.c())
}

// MixAudio mixes audio data. New programs might want to use MixAudioFormat() instead.
// (https://wiki.libsdl.org/SDL_MixAudio)
func MixAudio(dst, src *uint8, len uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudio(_dst, _src, C.Uint32(len), C.int(volume))
}

// MixAudioFormat mixes audio data in a specified format.
// (https://wiki.libsdl.org/SDL_MixAudioFormat)
func MixAudioFormat(dst, src *uint8, format AudioFormat, len uint32, volume int) {
	_dst := (*C.Uint8)(unsafe.Pointer(dst))
	_src := (*C.Uint8)(unsafe.Pointer(src))
	C.SDL_MixAudioFormat(_dst, _src, format.c(), C.Uint32(len), C.int(volume))
}

// LockAudio locks the audio device. New programs might want to use LockAudioDevice() instead.
// (https://wiki.libsdl.org/SDL_LockAudio)
func LockAudio() {
	C.SDL_LockAudio()
}

// LockAudioDevice locks out the audio callback function for a specified device.
// (https://wiki.libsdl.org/SDL_LockAudioDevice)
func LockAudioDevice(dev AudioDeviceID) {
	C.SDL_LockAudioDevice(dev.c())
}

// UnlockAudio unlocks the audio device. New programs might want to use UnlockAudioDevice() instead.
// (https://wiki.libsdl.org/SDL_UnlockAudio)
func UnlockAudio() {
	C.SDL_UnlockAudio()
}

// UnlockAudioDevice unlocks the audio callback function for a specified device.
// (https://wiki.libsdl.org/SDL_UnlockAudioDevice)
func UnlockAudioDevice(dev AudioDeviceID) {
	C.SDL_UnlockAudioDevice(dev.c())
}

// CloseAudio closes the audio device. New programs might want to use CloseAudioDevice() instead.
// (https://wiki.libsdl.org/SDL_CloseAudio)
func CloseAudio() {
	C.SDL_CloseAudio()
}

// CloseAudioDevice shuts down audio processing and closes the audio device.
// (https://wiki.libsdl.org/SDL_CloseAudioDevice)
func CloseAudioDevice(dev AudioDeviceID) {
	C.SDL_CloseAudioDevice(dev.c())
}

// NewAudioStream creates a new audio stream
// TODO: (https://wiki.libsdl.org/SDL_NewAudioStream)
func NewAudioStream(srcFormat AudioFormat, srcChannels uint8, srcRate int, dstFormat AudioFormat, dstChannels uint8, dstRate int) (stream *AudioStream, err error) {
	_srcFormat := C.SDL_AudioFormat(srcFormat)
	_srcChannels := C.Uint8(srcChannels)
	_srcRate := C.int(srcRate)
	_dstFormat := C.SDL_AudioFormat(dstFormat)
	_dstChannels := C.Uint8(dstChannels)
	_dstRate := C.int(dstRate)

	stream = (*AudioStream)(C.SDL_NewAudioStream(_srcFormat, _srcChannels, _srcRate, _dstFormat, _dstChannels, _dstRate))
	if stream == nil {
		err = GetError()
	}
	return
}

// Put adds data to be converted/resampled to the stream
// TODO: (https://wiki.libsdl.org/SDL_AudioStreamPut)
func (stream *AudioStream) Put(buf []byte) (err error) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	_buf := unsafe.Pointer(sliceHeader.Data)
	_len := C.int(len(buf))
	ret := int(C.SDL_AudioStreamPut(stream.cptr(), _buf, _len))
	err = errorFromInt(ret)
	return
}

// Get gets converted/resampled data from the stream. Returns the number of bytes read from the stream.
// (https://wiki.libsdl.org/SDL_AudioStreamGet)
func (stream *AudioStream) Get(buf []byte) (n int, err error) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	_buf := unsafe.Pointer(sliceHeader.Data)
	_len := C.int(len(buf))
	if ret := int(C.SDL_AudioStreamGet(stream.cptr(), _buf, _len)); ret < 0 {
		return 0, errorFromInt(ret)
	} else {
		return ret, nil
	}
}

// Available gets the number of converted/resampled bytes available
// (https://wiki.libsdl.org/SDL_AudioStreamAvailable)
func (stream *AudioStream) Available() (n int) {
	n = int(C.SDL_AudioStreamAvailable(stream.cptr()))
	return
}

// Flush tells the stream that you're done sending data, and anything being buffered
// should be converted/resampled and made available immediately.
// TODO: (https://wiki.libsdl.org/SDL_AudioStreamFlush)
func (stream *AudioStream) Flush() (err error) {
	ret := int(C.SDL_AudioStreamFlush(stream.cptr()))
	err = errorFromInt(ret)
	return
}

// Clear clears any pending data in the stream without converting it
// TODO: (https://wiki.libsdl.org/SDL_AudioStreamClear)
func (stream *AudioStream) Clear() {
	C.SDL_AudioStreamClear(stream.cptr())
}

// Free frees the audio stream
// TODO: (https://wiki.libsdl.org/SDL_AudoiStreamFree)
func (stream *AudioStream) Free() {
	C.SDL_FreeAudioStream(stream.cptr())
}

// GetAudioDeviceSpec returns the preferred audio format of a specific audio device.
// (https://wiki.libsdl.org/SDL_GetAudioDeviceSpec)
func GetAudioDeviceSpec(index int, isCapture bool) (spec *AudioSpec, err error) {
	spec = &AudioSpec{}
	err = errorFromInt(int(C.SDL_GetAudioDeviceSpec(C.int(index), C.int(Btoi(isCapture)), spec.cptr())))
	return
}
