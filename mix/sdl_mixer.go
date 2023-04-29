// Package mix is an audio mixer library based on the SDL library.
package mix

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
//
//#if !(SDL_MIXER_VERSION_ATLEAST(2,0,2))
//
//#if defined(WARN_OUTDATED)
//#pragma message("Mix_OpenAudioDevice is not supported before SDL_mixer 2.0.2")
//#endif
//
//static inline int Mix_OpenAudioDevice(int frequency, Uint16 format, int channels, int chunksize, const char* device, int allowed_changes)
//{
//	return -1;
//}
//#endif
import "C"
import "unsafe"
import "reflect"
import "github.com/veandco/go-sdl2/sdl"

// Chunk is the internal format for an audio chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_85.html)
type Chunk struct {
	allocated int32  // a boolean indicating whether to free abuf when the chunk is freed
	buf       *uint8 // pointer to the sample data, which is in the output format and sample rate
	len_      uint32 // length of abuf in bytes
	volume    uint8  // 0 = silent, 128 = max volume. This takes effect when mixing
}

// The different supported fading types.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_88.html)
const (
	NO_FADING Fading = iota
	FADING_OUT
	FADING_IN
)

// Dynamic libraries init flags used in mix.Init().
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_9.html)
const (
	INIT_FLAC = C.MIX_INIT_FLAC
	INIT_MOD  = C.MIX_INIT_MOD
	INIT_MP3  = C.MIX_INIT_MP3
	INIT_OGG  = C.MIX_INIT_OGG
)

// These are types of music files (not libraries used to load them)
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_87.html)
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

// Good default values for a PC soundcard.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_91.html)
const (
	CHANNELS          = 8
	DEFAULT_FREQUENCY = C.MIX_DEFAULT_FREQUENCY
	DEFAULT_FORMAT    = C.MIX_DEFAULT_FORMAT
	DEFAULT_CHANNELS  = C.MIX_DEFAULT_CHANNELS
	MAX_VOLUME        = C.MIX_MAX_VOLUME
	CHANNEL_POST      = -2
	EFFECTSMAXSPEED   = "MIX_EFFECTSMAXSPEED"
)

// DEFAULT_CHUNKSIZE is the default size of a chunk.
const DEFAULT_CHUNKSIZE = 1024

// Music is a data type used for Music data.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_86.html)
type Music C.Mix_Music

// MusicType is a file format encoding of the music.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_87.html)
type MusicType int

// Fading ia a return value from mix.FadingMusic() and mix.FadingChannel(). If no fading is taking place on the queried channel or music, then mix.NO_FADING is returned.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_88.html)
type Fading int

func cint(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

// Init loads dynamic libraries and prepares them for use. Flags should be one or more flags from mix.InitFlags OR'd together.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_9.html)
func Init(flags int) error {
	initted := int(C.Mix_Init(C.int(flags)))
	if initted&flags != flags {
		return sdl.GetError()
	}
	return nil
}

// Quit unloads libraries loaded with mix.Init().
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_10.html)
func Quit() {
	C.Mix_Quit()
}

// OpenAudio opens the mixer with a certain audio format.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_11.html)
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

// OpenAudioDevice opens the mixer with a certain audio format and a device.
// (http://hg.libsdl.org/SDL_mixer/rev/fb0562cc1559)
// (https://wiki.libsdl.org/SDL_OpenAudioDevice)
func OpenAudioDevice(frequency int, format uint16, channels, chunksize int, device string, allowedChanges int) error {
	_frequency := (C.int)(frequency)
	_format := (C.Uint16)(format)
	_channels := (C.int)(channels)
	_chunksize := (C.int)(chunksize)
	_allowedChanges := (C.int)(allowedChanges)
	_device := C.CString(device)
	defer C.free(unsafe.Pointer(_device))
	if device == "" {
		// Passing in a device name of NULL requests the most reasonable default
		// (and is equivalent to what SDL_OpenAudio() does to choose a device)
		_device = nil
	}
	if C.Mix_OpenAudioDevice(_frequency, _format, _channels, _chunksize, _device, _allowedChanges) < 0 {
		return sdl.GetError()
	}
	return nil
}

// AllocateChannels dynamically changes the number of channels managed by the mixer. If decreasing the number of channels, the upper channels are stopped. This function returns the new number of allocated channels. Returns: The number of channels allocated.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_26.html)
func AllocateChannels(numchans int) int {
	_numchans := (C.int)(numchans)
	return int(C.Mix_AllocateChannels(_numchans))
}

// QuerySpec returns the actual audio device parameters.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_15.html)
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

// LoadWAVRW loads src for use as a sample. This can load WAVE, AIFF, RIFF, OGG, and VOC formats. Note: You must call mix.OpenAudio() before this. It must know the output characteristics so it can convert the sample for playback, it does this conversion at load time. Returns: a pointer to the sample as a mix.Chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_20.html)
func LoadWAVRW(src *sdl.RWops, freesrc bool) (chunk *Chunk, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := cint(freesrc)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_LoadWAV_RW(_src, _freesrc)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// LoadWAV loads file for use as a sample. This is actually mix.LoadWAVRW(sdl.RWFromFile(file, "rb"), 1). This can load WAVE, AIFF, RIFF, OGG, and VOC files. Note: You must call SDL_OpenAudio before this. It must know the output characteristics so it can convert the sample for playback, it does this conversion at load time. Returns: a pointer to the sample as a mix.Chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_19.html)
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

// LoadMUS loads music file to use. This can load WAVE, MOD, MIDI, OGG, MP3, FLAC, and any file that you use a command to play with. If you are using an external command to play the music, you must call mix.SetMusicCMD before this, otherwise the internal players will be used. Alternatively, if you have set an external command up and don't want to use it, you must call Mix_SetMusicCMD(nil) to use the built-in players again. Returns: A pointer to a mix.Music.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_55.html)
func LoadMUS(file string) (mus *Music, err error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS(_file)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

// LoadMUSRW loads a music file from an sdl.RWop object (Ogg and MikMod specific currently).
func LoadMUSRW(src *sdl.RWops, freesrc int) (mus *Music, err error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	mus = (*Music)(unsafe.Pointer(C.Mix_LoadMUS_RW(_src, _freesrc)))
	if mus == nil {
		err = sdl.GetError()
	}
	return
}

// LoadMUSTypeRW loads a music file from an sdl.RWop object assuming a specific format.
func LoadMUSTypeRW(src *sdl.RWops, type_ MusicType, freesrc int) (mus *Music, err error) {
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

// QuickLoadWAV loads mem as a WAVE/RIFF file into a new sample. The WAVE in mem must be already in the output format. It would be better to use mix.LoadWAVRW() if you aren't sure. Note: This function does very little checking. If the format mismatches the output format, or if the buffer is not a WAVE, it will not return an error. This is probably a dangerous function to use. Returns: a pointer to the sample as a mix.Chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_21.html)
func QuickLoadWAV(mem []byte) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(&mem[0])
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_WAV(_mem)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// QuickLoadRAW loads mem as a raw sample. The data in mem must be already in the output format. If you aren't sure what you are doing, this is not a good function for you! Note: This function does very little checking. If the format mismatches the output format it will not return an error. This is probably a dangerous function to use. Returns: a pointer to the sample as a mix.Chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_22.html)
func QuickLoadRAW(mem *uint8, len_ uint32) (chunk *Chunk, err error) {
	_mem := (*C.Uint8)(mem)
	_len := (C.Uint32)(len_)
	chunk = (*Chunk)(unsafe.Pointer(C.Mix_QuickLoad_RAW(_mem, _len)))
	if chunk == nil {
		err = sdl.GetError()
	}
	return
}

// Free frees the memory used in chunk, and frees chunk itself as well. Do not use chunk after this without loading a new sample to it. Note: It's a bad idea to free a chunk that is still being played...
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_24.html)
func (chunk *Chunk) Free() {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	C.Mix_FreeChunk(_chunk)
}

// Free frees the loaded music. If music is playing it will be halted. If music is fading out, then this function will wait (blocking) until the fade out is complete.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_56.html)
func (music *Music) Free() {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	C.Mix_FreeMusic(_music)
}

// Type tells you the file format encoding of the music. This may be handy when used with mix.SetMusicPosition(), and other music functions that vary based on the type of music being played. If you want to know the type of music currently being played, pass in nil to music. Returns: The type of music or if music is nil then the currently playing music type, otherwise NONE if no music is playing.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_70.html)
func (music *Music) Type() MusicType {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	return (MusicType)(C.Mix_GetMusicType(_music))
}

// SetPanning sets a panning effect on audio channel. This effect will only work on stereo audio. Meaning you called mix.OpenAudio() with 2 channels (mix.DEFAULT_CHANNELS). The easiest way to do true panning is to call Mix_SetPanning(channel, left, 254 - left); so that the total volume is correct, if you consider the maximum volume to be 127 per channel for center, or 254 max for left, this works, but about halves the effective volume. This Function registers the effect for you, so don't try to mix.RegisterEffect() it yourself. NOTE: Setting both left and right to 255 will unregister the effect from channel. You cannot unregister it any other way, unless you use mix.UnregisterAllEffects() on the channel. NOTE: Using this function on a mono audio device will not register the effect, nor will it return an error status.
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

// SetPosition emulates a simple 3D audio effect. It's not all that realistic, but it can help improve some level of realism. By giving it the angle and distance from the camera's point of view, the effect pans and attenuates volumes. If you are looking for better positional audio, using OpenAL is suggested. NOTE: Using angle and distance of 0, will cause the effect to unregister itself from channel. You cannot unregister it any other way, unless you use mix.UnregisterAllEffects() on the channel.
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

// SetDistance simulates a simple attenuation of volume due to distance. The volume never quite reaches silence, even at max distance. NOTE: Using a distance of 0 will cause the effect to unregister itself from channel. You cannot unregister it any other way, unless you use mix.UnregisterAllEffects() on the channel.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_81.html)
func SetDistance(channel int, distance uint8) error {
	_channel := (C.int)(channel)
	_distance := (C.Uint8)(distance)
	if (C.Mix_SetDistance(_channel, _distance)) == 0 {
		return sdl.GetError()
	}
	return nil
}

// SetReverseStereo swaps left and right channel sound. NOTE: Using a flip of 0, will cause the effect to unregister itself from channel. You cannot unregister it any other way, unless you use mix.UnregisterAllEffects() on the channel.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_83.html)
func SetReverseStereo(channel, flip int) error {
	_channel := (C.int)(channel)
	_flip := (C.int)(flip)
	if (C.Mix_SetReverseStereo(_channel, _flip)) == 0 {
		return sdl.GetError()
	}
	return nil
}

// ReserveChannels reserves num channels from being used when playing samples when passing in -1 as a channel number to playback functions. The channels are reserved starting from channel 0 to num-1. Passing in zero will unreserve all channels. Normally mix starts without any channels reserved. Returns: The number of channels reserved. Never fails, but may return less channels than you ask for, depending on the number of channels previously allocated.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_43.html)
func ReserveChannels(num int) int {
	_num := (C.int)(num)
	return (int)(C.Mix_ReserveChannels(_num))
}

// GroupChannel adds which channel to group tag, or reset it's group to the default group tag (-1). Returns: True on success. False is returned when the channel specified is invalid.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_44.html)
func GroupChannel(which, tag int) bool {
	_which := (C.int)(which)
	_tag := (C.int)(tag)
	return C.Mix_GroupChannel(_which, _tag) != 0
}

// GroupChannels adds channels starting at from up through to to group tag, or reset it's group to the default group tag (-1). Returns: The number of tagged channels on success. If that number is less than to-from+1 then some channels were no tagged because they didn't exist.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_45.html)
func GroupChannels(from, to, tag int) int {
	_from := (C.int)(from)
	_to := (C.int)(to)
	_tag := (C.int)(tag)
	return int(C.Mix_GroupChannels(_from, _to, _tag))
}

// GroupAvailable finds the first available (not playing) channel in group tag. Returns: The channel found on success. -1 is returned when no channels in the group are available.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_47.html)
func GroupAvailable(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupAvailable(_tag))
}

// GroupCount returns the number of channels in group tag. Returns: The number of channels found in the group.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_46.html)
func GroupCount(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupCount(_tag))
}

// GroupOldest finds the oldest actively playing channel in group tag. Returns: The channel found on success. -1 is returned when no channels in the group are playing or the group is empty.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_48.html)
func GroupOldest(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupOldest(_tag))
}

// GroupNewer finds the newest, most recently started, actively playing channel in group tag. Returns: The channel found on success. -1 is returned when no channels in the group are playing or the group is empty.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_49.html)
func GroupNewer(tag int) int {
	_tag := (C.int)(tag)
	return (int)(C.Mix_GroupNewer(_tag))
}

// PlayTimed plays chunk on channel, or if channel is -1, pick the first free unreserved channel. The sample will play for loops+1 number of times, unless stopped by halt, or fade out, or setting a new expiration time of less time than it would have originally taken to play the loops, or closing the mixer. If the sample is long enough and has enough loops then the sample will stop after ticks milliseconds. Otherwise this function is the same as mix.PlayChannel(). Returns: the channel the sample is played on.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_29.html)
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

// Play plays chunk on channel, or if channel is -1, pick the first free unreserved channel. The sample will play for loops+1 number of times, unless stopped by halt, or fade out, or setting a new expiration time of less time than it would have originally taken to play the loops, or closing the mixer. Note: this just calls mix.PlayChannelTimed() with ticks set to -1. Returns: the channel the sample is played on.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_28.html)
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

// Play plays the loaded music loop times through from start to finish. The previous music will be halted, or if fading out it waits (blocking) for that to finish.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_57.html)
func (music *Music) Play(loops int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	if C.Mix_PlayMusic(_music, _loops) == -1 {
		return sdl.GetError()
	}
	return nil
}

// FadeIn fades in over ms milliseconds of time, the loaded music, playing it loop times through from start to finish. The fade in effect only applies to the first loop. Any previous music will be halted, or if it is fading out it will wait (blocking) for the fade to complete. This function is the same as mix.*Music.FadeInPos(loops, ms, 0).
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_58.html)
func (music *Music) FadeIn(loops, ms int) error {
	_music := (*C.Mix_Music)(unsafe.Pointer(music))
	_loops := (C.int)(loops)
	_ms := (C.int)(ms)
	if C.Mix_FadeInMusic(_music, _loops, _ms) == -1 {
		return sdl.GetError()
	}
	return nil
}

// FadeInPos fades in over ms milliseconds of time, the loaded music, playing it loop times through from start to finish. The fade in effect only applies to the first loop. The first time the music is played, it posistion will be set to posistion, which means different things for different types of music files, see mix.SetMusicPosition() for more info on that. Any previous music will be halted, or if it is fading out it will wait (blocking) for the fade to complete.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_59.html)
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

// FadeIn plays chunk on channel, or if channel is -1, pick the first free unreserved channel. The channel volume starts at 0 and fades up to full volume over ms milliseconds of time. The sample may end before the fade-in is complete if it is too short or doesn't have enough loops. The sample will play for loops+1 number of times, unless stopped by halt, or fade out, or setting a new expiration time of less time than it would have originally taken to play the loops, or closing the mixer. Note: this just calls mix.FadeInTimed() with ticks set to -1. Returns: the channel the sample is played on.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_30.html)
func (chunk *Chunk) FadeIn(channel, loops, ms int) (channel_ int, err error) {
	return chunk.FadeInTimed(channel, loops, ms, -1)
}

// FadeInTimed plays chunk on channel, or if channel is -1, pick the first free unreserved channel. The channel volume starts at 0 and fades up to full volume over ms milliseconds of time. The sample may end before the fade-in is complete if it is too short or doesn't have enough loops. The sample will play for loops+1 number of times, unless stopped by halt, or fade out, or setting a new expiration time of less time than it would have originally taken to play the loops, or closing the mixer. If the sample is long enough and has enough loops then the sample will stop after ticks milliseconds. Otherwise this function is the same as mix.FadeIn. Returns: the channel the sample is played on.
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

// Volume sets the volume for any allocated channel. If channel is -1 then all channels at are set at once. The volume is applied during the final mix, along with the sample volume. So setting this volume to 64 will halve the output of all samples played on the specified channel. All channels default to a volume of 128, which is the max. Newly allocated channels will have the max volume set, so setting all channels volumes does not affect subsequent channel allocations. Returns: current volume of the channel. If channel is -1, the average volume is returned.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_27.html)
func Volume(channel, volume int) int {
	_channel := (C.int)(channel)
	_volume := (C.int)(volume)
	return (int)(C.Mix_Volume(_channel, _volume))
}

// Volume sets the chunks volume. The volume setting will take effect when the chunk is used on a channel, being mixed into the output. Returns: previous chunks volume setting. If you passed a negative value for volume then this volume is still the current volume for the chunk.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_23.html)
func (chunk *Chunk) Volume(volume int) int {
	_chunk := (*C.Mix_Chunk)(unsafe.Pointer(chunk))
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeChunk(_chunk, _volume))
}

// VolumeMusic sets the musics volume, if it is 0 or greater, and return the previous volume setting. Setting the volume during a fade will not work, the faders use this function to perform their effect! Setting volume while using an external music player set by mix.SetMusicCMD() will have no effect, and mix.GetError() will show the reason why not. Returns: The previous volume setting.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_61.html)
func VolumeMusic(volume int) int {
	_volume := (C.int)(volume)
	return (int)(C.Mix_VolumeMusic(_volume))
}

// HaltChannel halts the channels playback, or all channels if -1 is passed in. Any callback set by mix.ChannelFinished() will be called.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_34.html)
func HaltChannel(channel int) {
	_channel := (C.int)(channel)
	C.Mix_HaltChannel(_channel)
}

// HaltGroup halts playback on all channels in group tag. Any callback set by mix.ChannelFinished() will be called once for each channel that stops.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_51.html)
func HaltGroup(tag int) {
	_tag := (C.int)(tag)
	C.Mix_HaltGroup(_tag)
}

// HaltMusic halts playback of music. This interrupts music fader effects. Any callback set by mix.HookMusicFinished() will be called when the music stops
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_67.html)
func HaltMusic() {
	C.Mix_HaltMusic()
}

// ExpireChannel halts channel playback, or all channels if -1 is passed in, after ticks milliseconds. Any callback set by mix.ChannelFinished() will be called when the channel expires. Returns: Number of channels set to expire. Whether or not they are active.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_35.html)
func ExpireChannel(channel, ticks int) int {
	_channel := (C.int)(channel)
	_ticks := (C.int)(ticks)
	return int(C.Mix_ExpireChannel(_channel, _ticks))
}

// FadeOutChannel gradually fades out which channel over ms milliseconds starting from now. The channel will be halted after the fade out is completed. Only channels that are playing are set to fade out, including paused channels. Any callback set by mix.ChannelFinished() will be called when the channel finishes fading out. Returns: The number of channels set to fade out.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_36.html)
func FadeOutChannel(which, ms int) int {
	_which := (C.int)(which)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutChannel(_which, _ms))
}

// FadeOutGroup gradually fades out channels in group tag over ms milliseconds starting from now. The channels will be halted after the fade out is completed. Only channels that are playing are set to fade out, including paused channels. Any callback set by mix.ChannelFinished() will be called when each channel finishes fading out. Returns: The number of channels set to fade out.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_50.html)
func FadeOutGroup(tag, ms int) int {
	_tag := (C.int)(tag)
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutGroup(_tag, _ms))
}

// FadeOutMusic gGradually fades out the music over ms milliseconds starting from now. The music will be halted after the fade out is completed. Only when music is playing and not fading already are set to fade out, including paused channels. Any callback set by mix.HookMusicFinished() will be called when the music finishes fading out. Returns: TRUE on success, FALSE on failure.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_68.html)
func FadeOutMusic(ms int) bool {
	_ms := (C.int)(ms)
	return int(C.Mix_FadeOutMusic(_ms)) == 0
}

// FadingMusic reports whether music is fading in, out, or not at all. Does not tell you if the channel is playing anything, or paused, so you'd need to test that separately. Returns: the fading status.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_73.html)
func FadingMusic() Fading {
	return (Fading)(C.Mix_FadingMusic())
}

// FadingChannel reports whether which channel is fading in, out, or not. Does not tell you if the channel is playing anything, or paused, so you'd need to test that separately. Returns: the fading status.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_40.html)
func FadingChannel(which int) Fading {
	_which := (C.int)(which)
	return (Fading)(C.Mix_FadingChannel(_which))
}

// Pause pauses channel, or all playing channels if -1 is passed in. You may still halt a paused channel. Note: Only channels which are actively playing will be paused.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_32.html)
func Pause(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Pause(_channel)
}

// Resume unpauses channel, or all playing and paused channels if -1 is passed in.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_33.html)
func Resume(channel int) {
	_channel := (C.int)(channel)
	C.Mix_Resume(_channel)
}

// Paused reports whether the channel is paused, or not. Note: Does not check if the channel has been halted after it was paused, which may seem a little weird. Returns: Zero if the channel is not paused. Otherwise if you passed in -1, the number of paused channels is returned. If you passed in a specific channel, then 1 is returned if it is paused.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_39.html)
func Paused(channel int) int {
	_channel := (C.int)(channel)
	return int(C.Mix_Paused(_channel))
}

// PauseMusic pauses the music playback. You may halt paused music. Note: Music can only be paused if it is actively playing.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_62.html)
func PauseMusic() {
	C.Mix_PauseMusic()
}

// ResumeMusic unpauses the music. This is safe to use on halted, paused, and already playing music.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_63.html)
func ResumeMusic() {
	C.Mix_ResumeMusic()
}

// RewindMusic rewinds the music to the start. This is safe to use on halted, paused, and already playing music. It is not useful to rewind the music immediately after starting playback, because it starts at the beginning by default. This function only works for these streams: MOD, OGG, MP3, Native MIDI.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_64.html)
func RewindMusic() {
	C.Mix_RewindMusic()
}

// PausedMusic reports whether music is paused, or not. Note: Does not check if the music was been halted after it was paused, which may seem a little weird. Returns: FALSE if music is not paused. TRUE if it is paused.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_72.html)
func PausedMusic() bool {
	return int(C.Mix_PausedMusic()) > 0
}

// SetMusicPosition sets the position of the currently playing music. The position takes different meanings for different music sources.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_65.html)
func SetMusicPosition(position int64) error {
	_position := (C.double)(position)
	if C.Mix_SetMusicPosition(_position) == -1 {
		return sdl.GetError()
	}
	return nil
}

// Playing reports whether the channel is playing, or not. Note: Does not check if the channel has been paused. Returns: Zero if the channel is not playing. Otherwise if you passed in -1, the number of channels playing is returned. If you passed in a specific channel, then 1 is returned if it is playing.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_38.html)
func Playing(channel int) int {
	_channel := (C.int)(channel)
	return int(C.Mix_Playing(_channel))
}

// PlayingMusic reports whether music is actively playing, or not. Note: Does not check if the channel has been paused. Returns: Zero if the music is not playing, or 1 if it is playing.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_71.html)
func PlayingMusic() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

// SetMusicCMD sets up a command line music player to use to play music. Any music playing will be halted. The music file to play is set by calling mix.LoadMUS(filename), and the filename is appended as the last argument on the commandline. This allows you to reuse the music command to play multiple files. The command will be sent signals SIGTERM to halt, SIGSTOP to pause, and SIGCONT to resume. The command program should react correctly to those signals for it to function properly with SDL_Mixer. mix.VolumeMusic has no effect when using an external music player, and mix.GetError will have an error code set. You should set the music volume in the music player's command if the music player supports that. Looping music works, by calling the command again when the previous music player process has ended. Playing music through a command uses a forked process to execute the music command. To use the internal music players set the command to nil. NOTE: External music is not mixed by SDL_mixer, so no post-processing hooks will be for music. NOTE: Playing music through an external command may not work if the sound driver does not support multiple openings of the audio device, since SDL_Mixer already has the audio device open for playing samples through channels. NOTE: Commands are not totally portable, so be careful.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_66.html)
func SetMusicCMD(command string) error {
	_command := C.CString(command)
	defer C.free(unsafe.Pointer(_command))
	if C.Mix_SetMusicCMD(_command) == -1 {
		return sdl.GetError()
	}
	return nil
}

// SetSynchroValue sets the synchro value.
func SetSynchroValue(value int) bool {
	_value := (C.int)(value)
	return int(C.Mix_SetSynchroValue(_value)) == 0
}

// GetSynchroValue returns the synchro value.
func GetSynchroValue() int {
	return (int)(C.Mix_GetSynchroValue())
}

// GetChunk returns the most recent sample chunk pointer played on channel. This pointer may be currently playing, or just the last used. Note: The actual chunk may have been freed, so this pointer may not be valid anymore. Returns: Pointer to the mix.Chunk. nil is returned if the channel is not allocated, or if the channel has not played any samples yet.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_41.html)
func GetChunk(channel int) *Chunk {
	_channel := (C.int)(channel)
	return (*Chunk)(unsafe.Pointer(C.Mix_GetChunk(_channel)))
}

// CloseAudio shuts down and cleanup the mixer API. After calling this all audio is stopped, the device is closed, and the SDL_mixer functions should not be used. You may, of course, use mix.OpenAudio() to start the functionality again. Note: This function doesn't do anything until you have called it the same number of times that you called mix.OpenAudio(). You may use mix.QuerySpec() to find out how many times mix.CloseAudio() needs to be called before the device is actually closed.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_12.html)
func CloseAudio() {
	C.Mix_CloseAudio()
}

// GetNumChunkDecoders returns the number of sample chunk decoders available from the mix.GetChunkDecoder() function. This number can be different for each run of a program, due to the change in availability of shared libraries that support each format. Returns: The number of sample chunk decoders available.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_17.html)
func GetNumChunkDecoders() int {
	return int(C.Mix_GetNumChunkDecoders())
}

// GetChunkDecoder reutns the name of the indexed sample chunk decoder. You need to get the number of sample chunk decoders available using the mix.GetNumChunkDecoders() function. Returns: The name of the indexed sample chunk decoder. This string is owned by the SDL_mixer library, do not modify or free it. It is valid until you call mix.CloseAudio() the final time.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_18.html)
func GetChunkDecoder(index int) string {
	return C.GoString(C.Mix_GetChunkDecoder(C.int(index)))
}

// GetNumMusicDecoders returns the number of music decoders available from the mix.GetMusicDecoder() function. This number can be different for each run of a program, due to the change in availability of shared libraries that support each format. Returns: The number of music decoders available.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_53.html)
func GetNumMusicDecoders() int {
	return int(C.Mix_GetNumMusicDecoders())
}

// GetMusicDecoder returns the name of the indexed music decoder. You need to get the number of music decoders available using the mix.GetNumMusicDecoders() function. Returns: The name of the indexed music decoder. This string is owned by the SDL_mixer library, do not modify or free it. It is valid until you call mix.CloseAudio() the final time.
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

// SetPostMix hooks a processor function mixFunc to the postmix stream for post processing effects. You may just be reading the data and displaying it, or you may be altering the stream to add an echo. Most processors also have state data that they allocate as they are in use, this would be stored in the arg pointer data space. This processor is never really finished, until the audio device is closed, or you pass nil as the mixFunc. There can only be one postmix function used at a time through this method. Use mix.RegisterEffect(MIX_CHANNEL_POST, mixFunc, nil, arg) to use multiple postmix processors. This postmix processor is run AFTER all the registered postmixers set up by mix.RegisterEffect().
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_79.html)
func SetPostMix(mixFunc func([]uint8)) {
	postMixFunc = mixFunc
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

// HookMusic sets up a custom music player function. The function will be called with arg passed into the udata parameter when the musicFunc is called. The stream parameter passes in the audio stream buffer to be filled with len bytes of music. The music player will then be called automatically when the mixer needs it. Music playing will start as soon as this is called. All the music playing and stopping functions have no effect on music after this. Pause and resume will work. Using a custom music player and the internal music player is not possible, the custom music player takes priority. To stop the custom music player call mix.HookMusic(nil, nil). NOTE: NEVER call SDL_Mixer functions, nor sdl.LockAudio(), from a callback function.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_60.html)
func HookMusic(musicFunc func([]uint8)) {
	hookMusicFunc = musicFunc
	C.Mix_HookMusic((*[0]byte)(C.callHookMusic), nil)
}

//export callHookMusicFinished
func callHookMusicFinished() {
	if hookMusicFinishedFunc != nil {
		hookMusicFinishedFunc()
	}
}

var hookMusicFinishedFunc func()

// HookMusicFinished sets up a function to be called when music playback is halted. Any time music stops, the music_finished function will be called. Call with nil to remove the callback. NOTE: NEVER call SDL_Mixer functions, nor sdl.LockAudio, from a callback function.
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

// ChannelFinished sets the specified channel_finished function to called when channel playback is halted. The channel parameter will contain the channel number that has finished. NOTE: NEVER call SDL_Mixer functions, nor sdl.LockAudio(), from a callback function.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_37.html)
func ChannelFinished(channelFinished func(int)) {
	channelFinishedFunc = channelFinished
	C.Mix_ChannelFinished((*[0]byte)(C.callChannelFinished))
}

// EffectFuncT is the prototype for effect processing functions. These functions are used to apply effects processing on a sample chunk. As a channel plays a sample, the registered effect functions are called. Each effect would then read and perhaps alter the len bytes of stream. It may also be advantageous to keep the effect state in the udata, with would be setup when registering the effect function on a channel.
// (https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_89.html)
type EffectFuncT func(channel int, stream []byte)

// EffectDoneT is the prototype for effect processing functions. These functions are used to apply effects processing on a sample chunk. As a channel plays a sample, the registered effect functions are called. Each effect would then read and perhaps alter the len bytes of stream. It may also be advantageous to keep the effect state in the udata, with would be setup when registering the effect function on a channel.
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

// RegisterEffect hooks a processor function f into a channel for post processing effects. You may just be reading the data and displaying it, or you may be altering the stream to add an echo. Most processors also have state data that they allocate as they are in use, this would be stored in the arg pointer data space. When a processor is finished being used, any function passed into d will be called, which is when your processor should clean up the data in the arg data space. The effects are put into a linked list, and always appended to the end, meaning they always work on previously registered effects output. Effects may be added multiple times in a row. Effects are cumulative this way.
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

// UnregisterAllEffects removes all effects registered to channel. If the channel is active all the registered effects will have their mix.EffectDoneT functions called, if they were specified in mix.RegisterEffect().
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
