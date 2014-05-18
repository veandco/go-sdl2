package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	HAPTIC_CONSTANT = 1 << iota
	HAPTIC_SINE
	HAPTIC_SQUARE
	HAPTIC_TRIANGLE
	HAPTIC_SAWTOOTHUP
	HAPTIC_SAWTOOTHDOWN
	HAPTIC_RAMP
	HAPTIC_SPRING
	HAPTIC_DAMPER
	HAPTIC_INERTIA
	HAPTIC_FRICTION
	HAPTIC_CUSTOM
	HAPTIC_GAIN
	HAPTIC_AUTOCENTER
	HAPTIC_STATUS
	HAPTIC_PAUSE
)

const (
	HAPTIC_POLAR = iota
	HAPTIC_CARTESION
	HAPTIC_SPHERICAL
)

const HAPTIC_INFINITY = 4294967295

type Haptic C.SDL_Haptic

type HapticDirection struct {
	Type byte
	dir  [3]int32
}

type HapticConstant struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Level        int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticPeriodic struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Period       uint16
	Magnitude    int16
	Offset       int16
	Phase        uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticCondition struct {
	Type       uint16
	Direction  HapticDirection
	Length     uint32
	Delay      uint16
	Button     uint16
	Interval   uint16
	RightSat   [3]uint16
	LeftSat    [3]uint16
	RightCoeff [3]int16
	LeftCoeff  [3]int16
	Deadband   [3]uint16
	Center     [3]int16
}

type HapticRamp struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Start        int16
	End          int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticCustom struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Channels     uint8
	Period       uint16
	Samples      uint16
	Data         *uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticEffect interface{}

func (h *Haptic) cptr() *C.SDL_Haptic {
    return (*C.SDL_Haptic)(unsafe.Pointer(h))
}

func NumHaptics() int {
	return int(C.SDL_NumHaptics())
}

func HapticName(index int) string {
	return (C.GoString)(C.SDL_HapticName(C.int(index)))
}

func HapticOpen(index int) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpen(C.int(index))))
}

func HapticOpened(index int) int {
	return int(C.SDL_HapticOpened(C.int(index)))
}

func HapticIndex(h *Haptic) int {
	return int(C.SDL_HapticIndex(h.cptr()))
}

func MouseIsHaptic() int {
	return int(C.SDL_MouseIsHaptic())
}

func HapticOpenFromMouse() *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
}

func JoystickIsHaptic(joy *Joystick) int {
	return int(C.SDL_JoystickIsHaptic(joy.cptr()))
}

func HapticOpenFromJoystick(joy *Joystick) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromJoystick(joy.cptr())))
}

func (h *Haptic) Close() {
	C.SDL_HapticClose(h.cptr())
}

func (h *Haptic) NumEffects() int {
	return int(C.SDL_HapticNumEffects(h.cptr()))
}

func (h *Haptic) NumEffectsPlaying() int {
	return int(C.SDL_HapticNumEffectsPlaying(h.cptr()))
}

func (h *Haptic) Query() uint {
	return uint(C.SDL_HapticQuery(h.cptr()))
}

func (h *Haptic) EffectSupported(he *HapticEffect) int {
    _he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticEffectSupported(h.cptr(), _he))
}

func (h *Haptic) NewEffect(he *HapticEffect) int {
    _he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticNewEffect(h.cptr(), _he))
}

func (h *Haptic) UpdateEffect(effect int, data *HapticEffect) int {
    _data := (*C.SDL_HapticEffect)(unsafe.Pointer(data))
	return int(C.SDL_HapticUpdateEffect(h.cptr(), C.int(effect), _data))
}

func (h *Haptic) RunEffect(effect int, iterations uint32) int {
	return int(C.SDL_HapticRunEffect(h.cptr(), C.int(effect), C.Uint32(iterations)))
}

func (h *Haptic) StopEffect(effect int) int {
	return int(C.SDL_HapticStopEffect(h.cptr(), C.int(effect)))
}

func (h *Haptic) DestroyEffect(effect int) {
	C.SDL_HapticDestroyEffect(h.cptr(), C.int(effect))
}

func (h *Haptic) GetEffectStatus(effect int) int {
	return int(C.SDL_HapticGetEffectStatus(h.cptr(), C.int(effect)))
}

func (h *Haptic) SetGain(gain int) int {
	return int(C.SDL_HapticSetGain(h.cptr(), C.int(gain)))
}

func (h *Haptic) SetAutocenter(autocenter int) int {
	return int(C.SDL_HapticSetAutocenter(h.cptr(), C.int(autocenter)))
}

func (h *Haptic) Pause() int {
	return int(C.SDL_HapticPause(h.cptr()))
}

func (h *Haptic) Unpause() int {
	return int(C.SDL_HapticUnpause(h.cptr()))
}

func (h *Haptic) StopAll() int {
	return int(C.SDL_HapticStopAll(h.cptr()))
}

func (h *Haptic) RumbleSupported() int {
	return int(C.SDL_HapticRumbleSupported(h.cptr()))
}

func (h *Haptic) RumbleInit() int {
	return int(C.SDL_HapticRumbleInit(h.cptr()))
}

func (h *Haptic) RumblePlay(strength float32, length uint32) int {
	return int(C.SDL_HapticRumblePlay(h.cptr(), C.float(strength), C.Uint32(length)))
}

func (h *Haptic) RumbleStop() int {
	return int(C.SDL_HapticRumbleStop(h.cptr()))
}
