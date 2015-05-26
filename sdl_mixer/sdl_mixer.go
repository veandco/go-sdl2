package mix

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_mixer
//#cgo darwin LDFLAGS: -framework SDL2 -framework SDL2_mixer
//#cgo linux freebsd pkg-config: sdl2
//#cgo linux freebsd LDFLAGS: -lSDL2_mixer
//#include <stdlib.h>
//#include "sdl_mixer_wrapper.h"
//
//extern void callPostMixFunction(void *udata, Uint8* stream, int length);
//extern void callHookMusic(void *udata, Uint8* stream, int length);
//extern void callHookMusicFinished();
//extern void callChannelFinished(int channel);
//extern void callEffectFunc(int channel, void *stream, int len, void *udata);
//extern void callEffectDone(int channel, void *udata);
//extern int callEachSoundFont(char* str, void* udata);
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

type Chunk struct {
	allocated int32
	buf       *uint8
	len_      uint32
	volume    uint8
}

const (
	NO_FADING = iota
	FADING_OUT
	FADING_IN
)

const (
	NONE = iota
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

const DEFAULT_FREQUENCY = 22050
const DEFAULT_FORMAT = sdl.AUDIO_S16LSB
const DEFAULT_CHANNELS = 2
const DEFAULT_CHUNKSIZE = 2
const MAX_VOLUME = 128

const CHANNEL_POST = -2
const EFFECTSMAXSPEED = "MIX_EFFECTSMAXSPEED"

type Music struct {
	unsafe.Pointer
}

type MusicType int
type Fading int

func cint(b bool) C.int {
	if b {
		return 1
	} else {
		return 0
	}
}

func OpenAudio(frequency int, format uint16, channels, chunksize int) error {
	_frequency := (C.int)(frequency)
	_format := (C.Uint16)(format)
	_channels := (C.int)(channels)
	_chunksize := (C.int)(chunksize)
	if C.Mix_OpenAudio(_frequency, _format, _channels, _chunksize) < 0 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func AllocateChannels(numchans int) int {
	_numchans := (C.int)(numchans)
	return int(C.Mix_AllocateChannels(_numchans))
}

// open is number of call to OpenAudio or 0 on error
func QuerySpec() (frequency int, format uint16, channels int, open int, err error) {
	var _frequency C.int
	var _format C.Uint16
	var _channels C.int
	open = int(C.Mix_QuerySpec(&_frequency, &_format, &_channels))
	if open == 0 {
		err = sdl.GetError()
	}
	frequency = int(_frequency)
	format = uint16(_format)
	channels = int(_channels)
	return
}

func LoadWAV_RW(src *sdl.RWops, freesrc bool) (chunk *Chunk, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := cint(freesrc)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_LoadWAV_RW(_src, _freesrc)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

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

func LoadMUS(file string) (mus *Music, err error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS(_file)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

func LoadMUS_RW(src *sdl.RWops, freesrc int) (mus *Music, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS_RW(_src, _freesrc)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

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

func QuickLoad_WAV(mem unsafe.Pointer) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(mem)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_WAV(_mem)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

func QuickLoad_RAW(mem *uint8, len_ uint32) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(mem)
	_len := (C.Uint32)(len_)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_RAW(_mem, _len)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

func (chunk *Chunk) Free() {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	C.Mix_FreeChunk(_chunk)
}

func (music *Music) Free() {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	C.Mix_FreeMusic(_music)
}

func (music *Music) Type() MusicType {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	return (MusicType)(C.Mix_GetMusicType(_music))
}

func SetPanning(channel int, left, right uint8) error {
	_channel := (C.int)(channel)
	_left := (C.Uint8)(left)
	_right := (C.Uint8)(right)
	if C.Mix_SetPanning(_channel, _left, _right) == 0 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func SetPosition(channel int, angle int16, distance uint8) error {
	_channel := (C.int)(channel)
	_angle := (C.Sint16)(angle)
	_distance := (C.Uint8)(distance)
	if (C.Mix_SetPosition(_channel, _angle, _distance)) == 0 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func SetDistance(channel int, distance uint8) error {
	_channel := (C.int)(channel)
	_distance := (C.Uint8)(distance)
	if (C.Mix_SetDistance(_channel, _distance)) == 0 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func SetReverseStereo(channel int, flip int) error {
	_channel := (C.int)(channel)
	_flip := (C.int)(flip)
	if (C.Mix_SetReverseStereo(_channel, _flip)) == 0 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func ReserveChannels(num int) int {
	_num := (C.int)(num)
	return (int)(C.Mix_ReserveChannels(_num))
}

func GroupChannel(which, tag int) bool {
	_which := (C.int)(which)
	_tag := (C.int)(tag)
	return C.Mix_GroupChannel(_which, _tag) != 0
}

func GroupChannels(from, to, tag int) int {
	_from := (C.int)(from)
	_to := (C.int)(to)
	_tag := (C.int)(tag)
	return int(C.Mix_GroupChannels(_from, _to, _tag))
}

func GroupAvailable(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupAvailable(_tag))
}

func GroupCount(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupCount(_tag))
}

func GroupOldest(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupOldest(_tag))
}

func GroupNewer(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupNewer(_tag))
}

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

func (chunk *Chunk) LengthInMs() int {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	return int(C.getChunkTimeMilliseconds(_chunk))
}

func (chunk *Chunk) PlayChannel(channel, loops int) (channel_ int, err error) {
	_channel := (C.int)(channel)
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_loops := (C.int)(loops)
	channel_ = int(C.Mix_PlayChannelTimed(_channel, _chunk, _loops, -1))
	if channel_ == -1 {
		err = sdl.GetError()
	}
	return
}

func (music *Music) Play(loops int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	if C.Mix_PlayMusic(_music, _loops) == -1 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func (music *Music) FadeIn(loops, ms int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	if C.Mix_FadeInMusic(_music, _loops, _ms) == -1 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func (music *Music) FadeInPos(loops, ms int, position float64) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	_position := (C.double)(position)
	if C.Mix_FadeInMusicPos(_music, _loops, _ms, _position) == -1 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func (chunk *Chunk) FadeIn(channel, loops, ms int) (channel_ int, err error) {
	return chunk.FadeInTimed(channel, loops, ms, -1)
}

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

func Volume(channel, volume int) int {
	_channel := (C.int)(channel)
	_volume := (C.int)(volume)
	return (int)(C.Mix_Volume(_channel, _volume))
}

func (chunk *Chunk) Volume(volume int) int {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeChunk(_chunk, _volume))
}

func VolumeMusic(volume int) int {
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeMusic(_volume))
}

func HaltChannel(channel int) {
	_channel := (C.int)(channel)
	C.Mix_HaltChannel(_channel)
}

func HaltGroup(tag int) {
	_tag := (C.int)(tag)
	C.Mix_HaltGroup(_tag)
}

func HaltMusic() {
	C.Mix_HaltMusic()
}

func ExpireChannel(channel, ticks int) int {
	_channel := (C.int)(channel)
	_ticks := (C.int)(ticks)
	return int(C.Mix_ExpireChannel(_channel, _ticks))
}

func FadeOutChannel(which, ms int) int {
	_which := (C.int)(which)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutChannel(_which, _ms))
}

func FadeOutGroup(tag, ms int) int {
	_tag := (C.int)(tag)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutGroup(_tag, _ms))
}

func FadeOutMusic(ms int) bool {
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutMusic(_ms)) == 0
}

func FadingMusic() Fading {
	return (Fading)(C.Mix_FadingMusic())
}

func FadingChannel(which int) Fading {
	_which := (C.int)(which)
	return (Fading)(C.Mix_FadingChannel(_which))
}

func Pause(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Pause(_channel)
}

func Resume(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Resume(_channel)
}

func Paused(channel int) bool {
	_channel := (C.int)(channel)
	return int(C.Mix_Paused(_channel)) > 0
}

func PauseMusic() {
	C.Mix_PauseMusic()
}

func ResumeMusic() {
	C.Mix_ResumeMusic()
}

func RewindMusic() {
	C.Mix_RewindMusic()
}

func PausedMusic() bool {
	return int(C.Mix_PausedMusic()) > 0
}

func SetMusicPosition(position int64) error {
	_position := (C.double)(position)
	if C.Mix_SetMusicPosition(_position) == -1 {
		return sdl.GetError()
	} else {
		return nil
	}
}

func Playing(channel int) bool {
	_channel := (C.int)(channel)
	return int(C.Mix_Playing(_channel)) > 0
}

func MusicPlaying() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

func SetMusicCMD(command string) error {
	_command := C.CString(command)
	defer C.free(unsafe.Pointer(_command))
	if C.Mix_SetMusicCMD(_command) == -1 {
		return sdl.GetError()
	} else {
		return nil
	}
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

// undocumented
func SetSoundFonts(paths string) bool {
	_paths := C.CString(paths)
	defer C.free(unsafe.Pointer(_paths))
	return int(C.Mix_SetSoundFonts(_paths)) == 0
}

// undocumented
func GetSoundFonts() string {
	return (string)(C.GoString(C.Mix_GetSoundFonts()))
}

func GetChunk(channel int) *Chunk {
	_channel := (C.int)(channel)
	return (*Chunk)(unsafe.Pointer(C.Mix_GetChunk(_channel)))
}

func CloseAudio() {
	C.Mix_CloseAudio()
}

func GetNumChunkDecoders() int {
	return int(C.Mix_GetNumChunkDecoders())
}

func GetChunkDecoder(index int) string {
	return C.GoString(C.Mix_GetChunkDecoder(C.int(index)))
}

func GetNumMusicDecoders() int {
	return int(C.Mix_GetNumMusicDecoders())
}

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

func HookMusic(music_func func([]uint8)) {
	hookMusicFunc = music_func
	C.Mix_HookMusic((*[0]byte)(C.callHookMusic), nil)
}

//export callHookMusicFinished
func callHookMusicFinished() {
	hookMusicFinishedFunc()
}

var hookMusicFinishedFunc func()

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

func ChannelFinished(channelFinished func(int)) {
	channelFinishedFunc = channelFinished
	C.Mix_ChannelFinished((*[0]byte)(C.callChannelFinished))
}

type EffectFuncT func(channel int, stream []byte)
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

func RegisterEffect(channel int, f EffectFuncT, d EffectDoneT) int {
	//the user data pointer is not required, because go has proper closures
	index := len(allEffectFunc)
	//since go functions can't be cast to C function pointers, we have a workaround here.
	allEffectFunc = append(allEffectFunc, f)
	allEffectDone = append(allEffectDone, d)
	return int(C.Mix_RegisterEffect(C.int(channel), (*[0]byte)(C.callEffectFunc), (*[0]byte)(C.callEffectDone), unsafe.Pointer(uintptr(index))))
}

// that we use the same function pointer for all functions definitely makes a problem when we want to remove it again
/*
func UnregisterEffect(channel int, f EffectFuncT) int {
	panic("TODO implement this function")
}
*/

func UnregisterAllEffects(channel int) int {
	// release all effect functions
	allEffectFunc = nil
	allEffectDone = nil
	return int(C.Mix_UnregisterAllEffects(C.int(channel)))
}

//export callEachSoundFont
func callEachSoundFont(str *C.char, udata unsafe.Pointer) C.int {
	return C.int(eachSoundFontFunc(C.GoString(str)))
}

var eachSoundFontFunc func(string) int

func EachSoundFont(function func(string) int) int {
	eachSoundFontFunc = function
	return int(C.Mix_EachSoundFont((*[0]byte)(C.callEachSoundFont), nil))
}
