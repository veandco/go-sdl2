package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	HAPTIC_CONSTANT  = C.SDL_HAPTIC_CONSTANT
	HAPTIC_SINE      = C.SDL_HAPTIC_SINE
	HAPTIC_LEFTRIGHT = C.SDL_HAPTIC_LEFTRIGHT
	//HAPTIC_SQUARE = C.SDL_HAPTIC_SQUARE (back in SDL 2.1)
	HAPTIC_TRIANGLE     = C.SDL_HAPTIC_TRIANGLE
	HAPTIC_SAWTOOTHUP   = C.SDL_HAPTIC_SAWTOOTHUP
	HAPTIC_SAWTOOTHDOWN = C.SDL_HAPTIC_SAWTOOTHDOWN
	HAPTIC_RAMP         = C.SDL_HAPTIC_RAMP
	HAPTIC_SPRING       = C.SDL_HAPTIC_SPRING
	HAPTIC_DAMPER       = C.SDL_HAPTIC_DAMPER
	HAPTIC_INERTIA      = C.SDL_HAPTIC_INERTIA
	HAPTIC_FRICTION     = C.SDL_HAPTIC_FRICTION
	HAPTIC_CUSTOM       = C.SDL_HAPTIC_CUSTOM
	HAPTIC_GAIN         = C.SDL_HAPTIC_GAIN
	HAPTIC_AUTOCENTER   = C.SDL_HAPTIC_AUTOCENTER
	HAPTIC_STATUS       = C.SDL_HAPTIC_STATUS
	HAPTIC_PAUSE        = C.SDL_HAPTIC_PAUSE
)

const (
	HAPTIC_POLAR     = C.SDL_HAPTIC_POLAR
	HAPTIC_CARTESIAN = C.SDL_HAPTIC_CARTESIAN
	HAPTIC_SPHERICAL = C.SDL_HAPTIC_SPHERICAL
	HAPTIC_INFINITY  = C.SDL_HAPTIC_INFINITY
)

type Haptic C.SDL_Haptic

// HapticDirection (https://wiki.libsdl.org/SDL_HapticDirection)
type HapticDirection struct {
	Type byte
	dir  [3]int32
}

// HapticConstant (https://wiki.libsdl.org/SDL_HapticConstant)
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

// HapticPeriodic (https://wiki.libsdl.org/SDL_HapticPeriodic)
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

// HapticCondition (https://wiki.libsdl.org/SDL_HapticCondition)
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

// HapticRamp (https://wiki.libsdl.org/SDL_HapticRamp)
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

// HapticLeftRight (https://wiki.libsdl.org/SDL_HapticLeftRight)
type HapticLeftRight struct {
	Type           uint16
	Length         uint32
	LargeMagnitude uint16
	SmallMagnitude uint16
}

// HapticCustom (https://wiki.libsdl.org/SDL_HapticCustom)
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

// HapticEffect (https://wiki.libsdl.org/SDL_HapticEffect)
type HapticEffect C.SDL_HapticEffect

func (he HapticEffect) Type() uint16 {
	return *((*uint16)(unsafe.Pointer(&he[0])))
}

func (he HapticEffect) Constant() *HapticConstant {
	return (*HapticConstant)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) Periodic() *HapticPeriodic {
	return (*HapticPeriodic)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) Condition() *HapticCondition {
	return (*HapticCondition)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) Ramp() *HapticRamp {
	return (*HapticRamp)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) LeftRight() *HapticLeftRight {
	return (*HapticLeftRight)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) Custom() *HapticCustom {
	return (*HapticCustom)(unsafe.Pointer(&he[0]))
}

func (he HapticEffect) SetType(typ uint16) {
	*((*uint16)(unsafe.Pointer(&he[0]))) = typ
}

func (h *Haptic) cptr() *C.SDL_Haptic {
	return (*C.SDL_Haptic)(unsafe.Pointer(h))
}

// NumHaptics (https://wiki.libsdl.org/SDL_NumHaptics)
func NumHaptics() int {
	return int(C.SDL_NumHaptics())
}

// HapticName (https://wiki.libsdl.org/SDL_HapticName)
func HapticName(index int) string {
	return (C.GoString)(C.SDL_HapticName(C.int(index)))
}

// HapticOpen (https://wiki.libsdl.org/SDL_HapticOpen)
func HapticOpen(index int) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpen(C.int(index))))
}

// HapticOpened (https://wiki.libsdl.org/SDL_HapticOpened)
func HapticOpened(index int) int {
	return int(C.SDL_HapticOpened(C.int(index)))
}

// HapticIndex (https://wiki.libsdl.org/SDL_HapticIndex)
func HapticIndex(h *Haptic) int {
	return int(C.SDL_HapticIndex(h.cptr()))
}

// MouseIsHaptic (https://wiki.libsdl.org/SDL_MouseIsHaptic)
func MouseIsHaptic() int {
	return int(C.SDL_MouseIsHaptic())
}

// HapticOpenFromMouse (https://wiki.libsdl.org/SDL_HapticOpenFromMouse)
func HapticOpenFromMouse() *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
}

// JoystickIsHaptic (https://wiki.libsdl.org/SDL_JoystickIsHaptic)
func JoystickIsHaptic(joy *Joystick) int {
	return int(C.SDL_JoystickIsHaptic(joy.cptr()))
}

// HapticOpenFromJoystick (https://wiki.libsdl.org/SDL_HapticOpenFromJoystick)
func HapticOpenFromJoystick(joy *Joystick) *Haptic {
	return (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromJoystick(joy.cptr())))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticClose)
func (h *Haptic) Close() {
	C.SDL_HapticClose(h.cptr())
}

// Haptic (https://wiki.libsdl.org/SDL_HapticNumAxes)
func (h *Haptic) NumAxes() int {
	return int(C.SDL_HapticNumAxes(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticNumEffects)
func (h *Haptic) NumEffects() int {
	return int(C.SDL_HapticNumEffects(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticNumEffectsPlaying)
func (h *Haptic) NumEffectsPlaying() int {
	return int(C.SDL_HapticNumEffectsPlaying(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticQuery)
func (h *Haptic) Query() uint {
	return uint(C.SDL_HapticQuery(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticEffectSupported)
func (h *Haptic) EffectSupported(he *HapticEffect) int {
	_he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticEffectSupported(h.cptr(), _he))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticNewEffect)
func (h *Haptic) NewEffect(he *HapticEffect) int {
	_he := (*C.SDL_HapticEffect)(unsafe.Pointer(he))
	return int(C.SDL_HapticNewEffect(h.cptr(), _he))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticUpdateEffect)
func (h *Haptic) UpdateEffect(effect int, data *HapticEffect) int {
	_data := (*C.SDL_HapticEffect)(unsafe.Pointer(data))
	return int(C.SDL_HapticUpdateEffect(h.cptr(), C.int(effect), _data))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticRunEffect)
func (h *Haptic) RunEffect(effect int, iterations uint32) int {
	return int(C.SDL_HapticRunEffect(h.cptr(), C.int(effect), C.Uint32(iterations)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticStopEffect)
func (h *Haptic) StopEffect(effect int) int {
	return int(C.SDL_HapticStopEffect(h.cptr(), C.int(effect)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticDestroyEffect)
func (h *Haptic) DestroyEffect(effect int) {
	C.SDL_HapticDestroyEffect(h.cptr(), C.int(effect))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticGetEffectStatus)
func (h *Haptic) GetEffectStatus(effect int) int {
	return int(C.SDL_HapticGetEffectStatus(h.cptr(), C.int(effect)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticSetGain)
func (h *Haptic) SetGain(gain int) int {
	return int(C.SDL_HapticSetGain(h.cptr(), C.int(gain)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticSetAutocenter)
func (h *Haptic) SetAutocenter(autocenter int) int {
	return int(C.SDL_HapticSetAutocenter(h.cptr(), C.int(autocenter)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticPause)
func (h *Haptic) Pause() int {
	return int(C.SDL_HapticPause(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticUnpause)
func (h *Haptic) Unpause() int {
	return int(C.SDL_HapticUnpause(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticStopAll)
func (h *Haptic) StopAll() int {
	return int(C.SDL_HapticStopAll(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticRumbleSupported)
func (h *Haptic) RumbleSupported() int {
	return int(C.SDL_HapticRumbleSupported(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticRumbleInit)
func (h *Haptic) RumbleInit() int {
	return int(C.SDL_HapticRumbleInit(h.cptr()))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticRumblePlay)
func (h *Haptic) RumblePlay(strength float32, length uint32) int {
	return int(C.SDL_HapticRumblePlay(h.cptr(), C.float(strength), C.Uint32(length)))
}

// Haptic (https://wiki.libsdl.org/SDL_HapticRumbleStop)
func (h *Haptic) RumbleStop() int {
	return int(C.SDL_HapticRumbleStop(h.cptr()))
}
