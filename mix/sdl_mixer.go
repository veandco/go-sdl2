package mix

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_mixer
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_mixer
//#include <stdlib.h>
//#include "sdl_mixer_wrapper.h"
//
//extern void callPostMixFunction(void *udata, Uint8* stream, int length);
//extern void callHookMusic(void *udata, Uint8* stream, int length);
//extern void callHookMusicFinished();
//extern void callChannelFinished(int channel);
//extern void callEffectFunc(int channel, void *stream, int len, void *udata);
//extern void callEffectDone(int channel, void *udata);
//
//static inline Uint32 getChunkTimeMilliseconds(Mix_Chunk *chunk)
//{
//	Uint32 points = 0;
//	Uint32 frames = 0;
//	int freq = 0;
//	Uint16 fmt = 0;
//	int chans = 0;
//	if (!Mix_QuerySpec(&freq, &fmt, &chans))
//		return 0;
//	points = (chunk->alen / ((fmt & 0xFF) / 8));
//	frames = (points / chans);
//	return (frames * 1000) / freq;
//}
import "C"
import "unsafe"
import "reflect"
import "github.com/veandco/go-sdl2/sdl"

// Chunk (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_85.html)
type Chunk struct {
	allocated int32
	buf       *uint8
	len_      uint32
	volume    uint8
}

const (
	NO_FADING Fading = iota
	FADING_OUT
	FADING_IN
)

const (
	INIT_FLAC = C.MIX_INIT_FLAC
	INIT_MOD  = C.MIX_INIT_MOD
	INIT_MP3  = C.MIX_INIT_MP3
	INIT_OGG  = C.MIX_INIT_OGG
)

const (
	NONE MusicType = iota
	CMD
	WAV
	MOD
	MID
	OGG
	MP3
	MP3_MAD
	FLAC
	MODPLUG
)

// Defines (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_91.html)
const (
	CHANNELS          = 8
	DEFAULT_FREQUENCY = C.MIX_DEFAULT_FREQUENCY
	DEFAULT_FORMAT    = C.MIX_DEFAULT_FORMAT
	DEFAULT_CHANNELS  = C.MIX_DEFAULT_CHANNELS
	MAX_VOLUME        = C.MIX_MAX_VOLUME
	CHANNEL_POST      = -2
	EFFECTSMAXSPEED   = "MIX_EFFECTSMAXSPEED"
)

const DEFAULT_CHUNKSIZE = 1024

// Music (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_86.html)
type Music C.Mix_Music

// MusicType (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_87.html)
type MusicType int

// Fading (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_88.html)
type Fading int

func cint(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func Init(flags int) error {
	initted := int(C.Mix_Init(C.int(flags)))
	if initted&flags != flags {
		return sdl.GetError()
	}
	return nil
}

func Quit() {
	C.Mix_Quit()
}

// OpenAudio (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_11.html)
func OpenAudio(frequency int, format uint16, channels, chunksize int) error {
	_frequency := (C.int)(frequency)
	_format := (C.Uint16)(format)
	_channels := (C.int)(channels)
	_chunksize := (C.int)(chunksize)
	if C.Mix_OpenAudio(_frequency, _format, _channels, _chunksize) < 0 {
		return sdl.GetError()
	}
	return nil
}

// AllocateChannels
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_26.html)
func AllocateChannels(numchans int) int {
	_numchans := (C.int)(numchans)
	return int(C.Mix_AllocateChannels(_numchans))
}

// QuerySpec (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_15.html)
// open is number of call to OpenAudio or 0 on error
func QuerySpec() (frequency int, format uint16, channels int, open int, err error) {
	var _frequency C.int
	var _format C.Uint16
	var _channels C.int
	if C.Mix_QuerySpec(&_frequency, &_format, &_channels) == 0 {
		err = sdl.GetError()
	}
	frequency = int(_frequency)
	format = uint16(_format)
	channels = int(_channels)
	return
}

// LoadWAV_RW
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_20.html)
func LoadWAV_RW(src *sdl.RWops, freesrc bool) (chunk *Chunk, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := cint(freesrc)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_LoadWAV_RW(_src, _freesrc)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// LoadWAV (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_19.html)
func LoadWAV(file string) (chunk *Chunk, err error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_rb := C.CString("rb")
	defer C.free(unsafe.Pointer(_rb))
	// why doesn't this call Mix_LoadWAV ?
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_LoadWAV_RW(C.SDL_RWFromFile(_file, _rb), 1)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// LoadMUS (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_55.html)
func LoadMUS(file string) (mus *Music, err error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS(_file)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

// undocumented
func LoadMUS_RW(src *sdl.RWops, freesrc int) (mus *Music, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS_RW(_src, _freesrc)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

// undocumented
func LoadMUSType_RW(src *sdl.RWops, type_ MusicType, freesrc int) (mus *Music, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_type := (C.Mix_MusicType)(type_)
	_freesrc := (C.int)(freesrc)
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUSType_RW(_src, _type,
		_freesrc)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

// QuickLoad_WAV
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_21.html)
func QuickLoad_WAV(mem unsafe.Pointer) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(mem)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_WAV(_mem)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// QuickLoad_RAW
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_22.html)
func QuickLoad_RAW(mem *uint8, len_ uint32) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(mem)
	_len := (C.Uint32)(len_)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_RAW(_mem, _len)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// Free (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_24.html)
func (chunk *Chunk) Free() {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	C.Mix_FreeChunk(_chunk)
}

// Free (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_56.html)
func (music *Music) Free() {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	C.Mix_FreeMusic(_music)
}

// Type (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_70.html)
func (music *Music) Type() MusicType {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	return (MusicType)(C.Mix_GetMusicType(_music))
}

// SetPanning
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_80.html)
func SetPanning(channel int, left, right uint8) error {
	_channel := (C.int)(channel)
	_left := (C.Uint8)(left)
	_right := (C.Uint8)(right)
	if C.Mix_SetPanning(_channel, _left, _right) == 0 {
		return sdl.GetError()
	}
	return nil
}

// SetPosition
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_82.html)
func SetPosition(channel int, angle int16, distance uint8) error {
	_channel := (C.int)(channel)
	_angle := (C.Sint16)(angle)
	_distance := (C.Uint8)(distance)
	if (C.Mix_SetPosition(_channel, _angle, _distance)) == 0 {
		return sdl.GetError()
	}
	return nil
}

// SetDistance
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_81.html)
func SetDistance(channel int, distance uint8) error {
	_channel := (C.int)(channel)
	_distance := (C.Uint8)(distance)
	if (C.Mix_SetDistance(_channel, _distance)) == 0 {
		return sdl.GetError()
	}
	return nil
}

// SetReverseStereo
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_83.html)
func SetReverseStereo(channel, flip int) error {
	_channel := (C.int)(channel)
	_flip := (C.int)(flip)
	if (C.Mix_SetReverseStereo(_channel, _flip)) == 0 {
		return sdl.GetError()
	}
	return nil
}

// ReserveChannels
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_43.html)
func ReserveChannels(num int) int {
	_num := (C.int)(num)
	return (int)(C.Mix_ReserveChannels(_num))
}

// GroupChannel
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_44.html)
func GroupChannel(which, tag int) bool {
	_which := (C.int)(which)
	_tag := (C.int)(tag)
	return C.Mix_GroupChannel(_which, _tag) != 0
}

// GroupChannels
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_45.html)
func GroupChannels(from, to, tag int) int {
	_from := (C.int)(from)
	_to := (C.int)(to)
	_tag := (C.int)(tag)
	return int(C.Mix_GroupChannels(_from, _to, _tag))
}

// GroupAvailable
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_47.html)
func GroupAvailable(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupAvailable(_tag))
}

// GroupCount
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_46.html)
func GroupCount(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupCount(_tag))
}

// GroupOldest
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_48.html)
func GroupOldest(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupOldest(_tag))
}

// GroupNewer
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_49.html)
func GroupNewer(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupNewer(_tag))
}

// PlayTimed (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_29.html)
func (chunk *Chunk) PlayTimed(channel, loops, ticks int) (channel_ int, err error) {
	_channel := (C.int)(channel)
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_loops := (C.int)(loops)
	_ticks := (C.int)(ticks)
	channel_ = int(C.Mix_PlayChannelTimed(_channel, _chunk, _loops, _ticks))
	if channel_ == -1 {
		err = sdl.GetError()
	}
	return
}

// LengthInMs returns the playing time of the chunk in milliseconds.
func (chunk *Chunk) LengthInMs() int {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	return int(C.getChunkTimeMilliseconds(_chunk))
}

// Play (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_28.html)
func (chunk *Chunk) Play(channel, loops int) (channel_ int, err error) {
	_channel := (C.int)(channel)
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_loops := (C.int)(loops)
	channel_ = int(C.Mix_PlayChannelTimed(_channel, _chunk, _loops, -1))
	if channel_ == -1 {
		err = sdl.GetError()
	}
	return
}

// Play (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_57.html)
func (music *Music) Play(loops int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	if C.Mix_PlayMusic(_music, _loops) == -1 {
		return sdl.GetError()
	}
	return nil
}

// FadeIn (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_58.html)
func (music *Music) FadeIn(loops, ms int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	if C.Mix_FadeInMusic(_music, _loops, _ms) == -1 {
		return sdl.GetError()
	}
	return nil
}

// FadeInPos (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_59.html)
func (music *Music) FadeInPos(loops, ms int, position float64) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	_position := (C.double)(position)
	if C.Mix_FadeInMusicPos(_music, _loops, _ms, _position) == -1 {
		return sdl.GetError()
	}
	return nil
}

// FadeIn (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_30.html)
func (chunk *Chunk) FadeIn(channel, loops, ms int) (channel_ int, err error) {
	return chunk.FadeInTimed(channel, loops, ms, -1)
}

// FadeInTimed
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_31.html)
func (chunk *Chunk) FadeInTimed(channel, loops, ms, ticks int) (channel_ int, err error) {
	_channel := (C.int)(channel)
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	_ticks := (C.int)(ticks)
	channel_ = int(C.Mix_FadeInChannelTimed(_channel, _chunk, _loops, _ms, _ticks))
	if channel_ == -1 {
		err = sdl.GetError()
	}
	return
}

// Volume (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_27.html)
func Volume(channel, volume int) int {
	_channel := (C.int)(channel)
	_volume := (C.int)(volume)
	return (int)(C.Mix_Volume(_channel, _volume))
}

// Volume (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_23.html)
func (chunk *Chunk) Volume(volume int) int {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeChunk(_chunk, _volume))
}

// VolumeMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_61.html)
func VolumeMusic(volume int) int {
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeMusic(_volume))
}

// HaltChannel
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_34.html)
func HaltChannel(channel int) {
	_channel := (C.int)(channel)
	C.Mix_HaltChannel(_channel)
}

// HaltGroup (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_51.html)
func HaltGroup(tag int) {
	_tag := (C.int)(tag)
	C.Mix_HaltGroup(_tag)
}

// HaltMusic (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_67.html)
func HaltMusic() {
	C.Mix_HaltMusic()
}

// ExpireChannel
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_35.html)
func ExpireChannel(channel, ticks int) int {
	_channel := (C.int)(channel)
	_ticks := (C.int)(ticks)
	return int(C.Mix_ExpireChannel(_channel, _ticks))
}

// FadeOutChannel
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_36.html)
func FadeOutChannel(which, ms int) int {
	_which := (C.int)(which)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutChannel(_which, _ms))
}

// FadeOutGroup
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_50.html)
func FadeOutGroup(tag, ms int) int {
	_tag := (C.int)(tag)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutGroup(_tag, _ms))
}

// FadeOutMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_68.html)
func FadeOutMusic(ms int) bool {
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutMusic(_ms)) == 0
}

// FadingMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_73.html)
func FadingMusic() Fading {
	return (Fading)(C.Mix_FadingMusic())
}

// FadingChannel
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_40.html)
func FadingChannel(which int) Fading {
	_which := (C.int)(which)
	return (Fading)(C.Mix_FadingChannel(_which))
}

// Pause (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_32.html)
func Pause(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Pause(_channel)
}

// Resume (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_33.html)
func Resume(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Resume(_channel)
}

// Paused (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_39.html)
func Paused(channel int) int {
	_channel := (C.int)(channel)
	return int(C.Mix_Paused(_channel))
}

// PauseMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_62.html)
func PauseMusic() {
	C.Mix_PauseMusic()
}

// ResumeMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_63.html)
func ResumeMusic() {
	C.Mix_ResumeMusic()
}

// RewindMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_64.html)
func RewindMusic() {
	C.Mix_RewindMusic()
}

// PausedMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_72.html)
func PausedMusic() bool {
	return int(C.Mix_PausedMusic()) > 0
}

// SetMusicPosition
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_65.html)
func SetMusicPosition(position int64) error {
	_position := (C.double)(position)
	if C.Mix_SetMusicPosition(_position) == -1 {
		return sdl.GetError()
	}
	return nil
}

// Playing (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_38.html)
func Playing(channel int) int {
	_channel := (C.int)(channel)
	return int(C.Mix_Playing(_channel))
}

// PlayingMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_71.html)
func PlayingMusic() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

// SetMusicCMD
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_66.html)
func SetMusicCMD(command string) error {
	_command := C.CString(command)
	defer C.free(unsafe.Pointer(_command))
	if C.Mix_SetMusicCMD(_command) == -1 {
		return sdl.GetError()
	}
	return nil
}

// undocumented
func SetSynchroValue(value int) bool {
	_value := (C.int)(value)
	return int(C.Mix_SetSynchroValue(_value)) == 0
}

// undocumented
func GetSynchroValue() int {
	return (int)(C.Mix_GetSynchroValue())
}

// GetChunk (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_41.html)
func GetChunk(channel int) *Chunk {
	_channel := (C.int)(channel)
	return (*Chunk)(unsafe.Pointer(C.Mix_GetChunk(_channel)))
}

// CloseAudio
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_12.html)
func CloseAudio() {
	C.Mix_CloseAudio()
}

// GetNumChunkDecoders
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_17.html)
func GetNumChunkDecoders() int {
	return int(C.Mix_GetNumChunkDecoders())
}

// GetChunkDecoder
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_18.html)
func GetChunkDecoder(index int) string {
	return C.GoString(C.Mix_GetChunkDecoder(C.int(index)))
}

// GetNumMusicDecoders
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_53.html)
func GetNumMusicDecoders() int {
	return int(C.Mix_GetNumMusicDecoders())
}

// GetMusicDecoder
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_54.html)
func GetMusicDecoder(index int) string {
	return C.GoString(C.Mix_GetMusicDecoder(C.int(index)))
}

//export callPostMixFunction
func callPostMixFunction(udata unsafe.Pointer, stream *C.Uint8, length C.int) {
	var slice []uint8
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(unsafe.Pointer(stream))
	header.Len = int(length)
	header.Cap = int(length)
	postMixFunc(slice)
}

var postMixFunc func([]uint8)

// SetPostMix
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_79.html)
func SetPostMix(mix_func func([]uint8)) {
	postMixFunc = mix_func
	C.Mix_SetPostMix((*[0]byte)(C.callPostMixFunction), nil)
}

//export callHookMusic
func callHookMusic(udata unsafe.Pointer, stream *C.Uint8, length C.int) {
	var slice []uint8
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(unsafe.Pointer(stream))
	header.Len = int(length)
	header.Cap = int(length)
	hookMusicFunc(slice)
}

var hookMusicFunc func([]uint8)

// HookMusic
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_60.html)
func HookMusic(music_func func([]uint8)) {
	hookMusicFunc = music_func
	C.Mix_HookMusic((*[0]byte)(C.callHookMusic), nil)
}

//export callHookMusicFinished
func callHookMusicFinished() {
	hookMusicFinishedFunc()
}

var hookMusicFinishedFunc func()

// HookMusicFinished
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_69.html)
func HookMusicFinished(musicFinished func()) {
	hookMusicFinishedFunc = musicFinished
	C.Mix_HookMusicFinished((*[0]byte)(C.callHookMusicFinished))
}

// GetMusicHookData
// data is not required, so never set and to need for this function

//export callChannelFinished
func callChannelFinished(channel C.int) {
	channelFinishedFunc(int(channel))
}

var channelFinishedFunc func(int)

// ChannelFinished
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_37.html)
func ChannelFinished(channelFinished func(int)) {
	channelFinishedFunc = channelFinished
	C.Mix_ChannelFinished((*[0]byte)(C.callChannelFinished))
}

// EffectFuncT
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_89.html)
type EffectFuncT func(channel int, stream []byte)

// EffectDoneT
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_90.html)
type EffectDoneT func(channel int)

// not perfect yet, since these slices are never freed of any data
var allEffectFunc []EffectFuncT
var allEffectDone []EffectDoneT

//export callEffectFunc
func callEffectFunc(channel C.int, stream unsafe.Pointer, length C.int, udata unsafe.Pointer) {
	index := int(uintptr(udata))
	var slice []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(stream)
	header.Len = int(length)
	header.Cap = int(length)
	allEffectFunc[index](int(channel), slice)
}

//export callEffectDone
func callEffectDone(channel C.int, udata unsafe.Pointer) {
	index := int(uintptr(udata))
	allEffectDone[index](int(channel))
}

// RegisterEffect
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_76.html)
func RegisterEffect(channel int, f EffectFuncT, d EffectDoneT) error {
	//the user data pointer is not required, because go has proper closures
	index := len(allEffectFunc)
	//since go functions can't be cast to C function pointers, we have a workaround here.
	allEffectFunc = append(allEffectFunc, f)
	allEffectDone = append(allEffectDone, d)
	if C.Mix_RegisterEffect(C.int(channel), (*[0]byte)(C.callEffectFunc), (*[0]byte)(C.callEffectDone), unsafe.Pointer(uintptr(index))) == 0 {
		return sdl.GetError()
	}
	return nil
}

// that we use the same function pointer for all functions definitely makes a problem when we want to remove it again
/*
func UnregisterEffect(channel int, f EffectFuncT) int {
	panic("TODO implement this function")
}
*/

// UnregisterAllEffects
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_78.html)
func UnregisterAllEffects(channel int) error {
	// release all effect functions
	allEffectFunc = nil
	allEffectDone = nil
	if C.Mix_UnregisterAllEffects(C.int(channel)) == 0 {
		return sdl.GetError()
	}
	return nil
}
