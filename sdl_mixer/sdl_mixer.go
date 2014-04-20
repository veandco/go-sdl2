package mix

//#cgo !linux,freebsd LDFLAGS: -lSDL2 -lSDL2_mixer
//#cgo !windows pkg-config: sdl2
//#cgo !windows LDFLAGS: -lSDL2_mixer
//#include <SDL2/SDL_mixer.h>
import "C"
import "unsafe"
import "github.com/jackyb/go-sdl2/sdl"

type Chunk struct {
	Allocated int32
	Buf *uint8
	Len uint32
	Volume uint8
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

func OpenAudio(frequency int, format uint16, channels, chunksize int) bool {
	_frequency := (C.int) (frequency)
	_format := (C.Uint16) (format)
	_channels := (C.int) (channels)
	_chunksize := (C.int) (chunksize)
	return int(C.Mix_OpenAudio(_frequency, _format, _channels, _chunksize)) == 0
}

func AllocateChannels(numchans int) bool {
	_numchans := (C.int) (numchans)
	return int(C.Mix_AllocateChannels(_numchans)) == 0
}

func QuerySpec(frequency *int, format *uint16, channels *int) bool {
	_frequency := (*C.int) (unsafe.Pointer(frequency))
	_format := (*C.Uint16) (unsafe.Pointer(format))
	_channels := (*C.int) (unsafe.Pointer(channels))
	return int(C.Mix_QuerySpec(_frequency, _format, _channels)) > 0
}

func LoadWAV_RW(src *sdl.RWops, freesrc int) *Chunk {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	_freesrc := (C.int) (freesrc)
	return (*Chunk) (unsafe.Pointer(C.Mix_LoadWAV_RW(_src, _freesrc)))
}

func LoadWAV(file string) *Chunk {
	_file := (C.CString) (file)
	return (*Chunk) (unsafe.Pointer(C.Mix_LoadWAV_RW(C.SDL_RWFromFile(_file,
			C.CString("rb")), 1)))
}

func LoadMUS(file string) *Music {
	_file := (C.CString) (file)
	return (*Music) (unsafe.Pointer(C.Mix_LoadMUS(_file)))
}

func LoadMUS_RW(src *sdl.RWops, freesrc int) *Music {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	_freesrc := (C.int) (freesrc)
	return (*Music) (unsafe.Pointer(C.Mix_LoadMUS_RW(_src, _freesrc)))
}

func LoadMUSType_RW(src *sdl.RWops, type_ MusicType, freesrc int) *Music {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	_type := (C.Mix_MusicType) (type_)
	_freesrc := (C.int) (freesrc)
	return (*Music) (unsafe.Pointer(C.Mix_LoadMUSType_RW(_src, _type,
			_freesrc)))
}

func QuickLoad_WAV(mem *uint8) *Chunk {
	_mem := (*C.Uint8) (mem)
	return (*Chunk) (unsafe.Pointer(C.Mix_QuickLoad_WAV(_mem)))
}

func QuickLoad_RAW(mem *uint8, len_ uint32) *Chunk {
	_mem := (*C.Uint8) (mem)
	_len := (C.Uint32) (len_)
	return (*Chunk) (unsafe.Pointer(C.Mix_QuickLoad_RAW(_mem, _len)))
}

func (chunk *Chunk) Free() {
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	C.Mix_FreeChunk(_chunk)
}

func (music *Music) Free() {
	_music := (*C.Mix_Music) (unsafe.Pointer(music))
	C.Mix_FreeMusic(_music)
}

func (music *Music) Type() MusicType {
	_music := (*C.Mix_Music) (unsafe.Pointer(music))
	return (MusicType) (C.Mix_GetMusicType(_music))
}

func SetPanning(channel int, left, right uint8) bool {
	_channel := (C.int) (channel)
	_left := (C.Uint8) (left)
	_right := (C.Uint8) (right)
	return int(C.Mix_SetPanning(_channel, _left, _right)) == 0
}

func SetPosition(channel int, angle int16, distance uint8) bool {
	_channel := (C.int) (channel)
	_angle := (C.Sint16) (angle)
	_distance := (C.Uint8) (distance)
	return int(C.Mix_SetPosition(_channel, _angle, _distance)) == 0
}

func SetDistance(channel int, distance uint8) bool {
	_channel := (C.int) (channel)
	_distance := (C.Uint8) (distance)
	return int(C.Mix_SetDistance(_channel, _distance)) == 0
}

func SetReverseStereo(channel int, flip int) bool {
	_channel := (C.int) (channel)
	_flip := (C.int) (flip)
	return int(C.Mix_SetReverseStereo(_channel, _flip)) == 0
}

func ReserveChannels(num int) int {
	_num := (C.int) (num)
	return (int) (C.Mix_ReserveChannels(_num))
}

func GroupChannel(which, tag int) bool {
	_which := (C.int) (which)
	_tag := (C.int) (tag)
	return int(C.Mix_GroupChannel(_which, _tag)) > 0
}

func GroupChannels(from, to, tag int) bool {
	_from := (C.int) (from)
	_to := (C.int) (to)
	_tag := (C.int) (tag)
	return int(C.Mix_GroupChannels(_from, _to, _tag)) > 0
}

func GroupAvailable(tag int) int {
	_tag := (C.int) (tag)
	return (int) (C.Mix_GroupAvailable(_tag))
}

func GroupCount(tag int) int {
	_tag := (C.int) (tag)
	return (int) (C.Mix_GroupCount(_tag))
}

func GroupOldest(tag int) int {
	_tag := (C.int) (tag)
	return (int) (C.Mix_GroupOldest(_tag))
}

func GroupNewer(tag int) int {
	_tag := (C.int) (tag)
	return (int) (C.Mix_GroupNewer(_tag))
}

func (chunk *Chunk) PlayTimed(channel, loops, ticks int) bool {
	_channel := (C.int) (channel)
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	_loops := (C.int) (loops)
	_ticks := (C.int) (ticks)
	return int(C.Mix_PlayChannelTimed(_channel, _chunk, _loops, _ticks)) == 0
}

func (chunk *Chunk) PlayChannel(channel, loops int) bool {
	_channel := (C.int) (channel)
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	_loops := (C.int) (loops)
	return int(C.Mix_PlayChannelTimed(_channel, _chunk, _loops, -1)) == 0
}

func (music *Music) Play(loops int) bool {
	_music := (*C.Mix_Music) (unsafe.Pointer(music))
	_loops := (C.int) (loops)
	return int(C.Mix_PlayMusic(_music, _loops)) == 0
}

func (music *Music) FadeIn(loops, ms int) bool {
	_music := (*C.Mix_Music) (unsafe.Pointer(music))
	_loops := (C.int) (loops)
	_ms := (C.int) (ms)
	return int(C.Mix_FadeInMusic(_music, _loops, _ms)) == 0
}

func (music *Music) FadeInPos(loops, ms int, position float64) bool {
	_music := (*C.Mix_Music) (unsafe.Pointer(music))
	_loops := (C.int) (loops)
	_ms := (C.int) (ms)
	_position := (C.double) (position)
	return int(C.Mix_FadeInMusicPos(_music, _loops, _ms, _position)) == 0
}

func (chunk *Chunk) FadeIn(channel, loops, ms, ticks int ) bool {
	_channel := (C.int) (channel)
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	_loops := (C.int) (loops)
	_ms := (C.int) (ms)
	return int(C.Mix_FadeInChannelTimed(_channel, _chunk, _loops, _ms, -1)) == 0
}

func (chunk *Chunk) FadeInTimed(channel, loops, ms, ticks int) bool {
	_channel := (C.int) (channel)
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	_loops := (C.int) (loops)
	_ms := (C.int) (ms)
	_ticks := (C.int) (ticks)
	return int(C.Mix_FadeInChannelTimed(_channel, _chunk, _loops, _ms,
			_ticks)) == 0
}

func SetVolume(channel, volume int) int {
	_channel := (C.int) (channel)
	_volume := (C.int) (volume)
	return (int) (C.Mix_Volume(_channel, _volume))
}

func (chunk *Chunk) SetVolume(volume int) int {
	_chunk := (*C.Mix_Chunk) (unsafe.Pointer(chunk))
	_volume := (C.int) (volume)
	return (int) (C.Mix_VolumeChunk(_chunk, _volume))
}

func SetMusicVolume(volume int) int {
	_volume := (C.int) (volume)
	return (int) (C.Mix_VolumeMusic(_volume))
}

func HaltChannel(channel int) bool {
	_channel := (C.int) (channel)
	return int(C.Mix_HaltChannel(_channel)) == 0
}

func HaltGroup(tag int) bool {
	_tag := (C.int) (tag)
	return int(C.Mix_HaltGroup(_tag)) == 0
}

func HaltMusic() bool {
	return int(C.Mix_HaltMusic()) == 0
}

func ExpireChannel(channel, ticks int) bool {
	_channel := (C.int) (channel)
	_ticks := (C.int) (ticks)
	return int(C.Mix_ExpireChannel(_channel, _ticks)) == 0
}

func FadeOutChannel(which, ms int) bool {
	_which := (C.int) (which)
	_ms := (C.int) (ms)
	return int(C.Mix_FadeOutChannel(_which, _ms)) == 0
}

func FadeOutGroup(tag, ms int) bool {
	_tag := (C.int) (tag)
	_ms := (C.int) (ms)
	return int(C.Mix_FadeOutGroup(_tag, _ms)) == 0
}

func FadeOutMusic(ms int) bool {
	_ms := (C.int) (ms)
	return int(C.Mix_FadeOutMusic(_ms)) == 0
}

func FadingMusic() Fading {
	return (Fading) (C.Mix_FadingMusic())
}

func FadingChannel(which int) Fading {
	_which := (C.int) (which)
	return (Fading) (C.Mix_FadingChannel(_which))
}

func Pause(channel int) {
	_channel := (C.int) (channel)
	C.Mix_Pause(_channel)
}

func Resume(channel int) {
	_channel := (C.int) (channel)
	C.Mix_Resume(_channel)
}

func Paused(channel int) bool {
	_channel := (C.int) (channel)
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

func SetMusicPosition(position int64) bool {
	_position := (C.double) (position)
	return int(C.Mix_SetMusicPosition(_position)) == 0
}

func Playing(channel int) bool {
	_channel := (C.int) (channel)
	return int(C.Mix_Playing(_channel)) > 0
}

func MusicPlaying() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

func SetMusicCMD(command string) bool {
	_command := (C.CString) (command)
	return int(C.Mix_SetMusicCMD(_command)) == 0
}

func SetSynchroValue(value int) bool {
	_value := (C.int) (value)
	return int(C.Mix_SetSynchroValue(_value)) == 0
}

func GetSynchroValue() int {
	return (int) (C.Mix_GetSynchroValue())
}

func SetSoundFonts(paths string) bool {
	_paths := (C.CString) (paths)
	return int(C.Mix_SetSoundFonts(_paths)) == 0
}

func GetSoundFonts() string {
	return (string) (C.GoString(C.Mix_GetSoundFonts()))
}

func GetChunk(channel int) *Chunk {
	_channel := (C.int) (channel)
	return (*Chunk) (unsafe.Pointer(C.Mix_GetChunk(_channel)))
}

func CloseAudio() {
	C.Mix_CloseAudio()
}
