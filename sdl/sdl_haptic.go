package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	HAPTIC_CONSTANT	= 1 << iota
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
	HAPTIC_POLAR		= iota
	HAPTIC_CARTESION
	HAPTIC_SPHERICAL
)

const HAPTIC_INFINITY = 4294967295

type Haptic C.SDL_Haptic

type HapticDirection struct {
	Type byte
	dir [3]int32
}

type HapticConstant struct {
	Type uint16
	Direction HapticDirection
	Length uint32
	Delay uint16
	Button uint16
	Interval uint16
	Level int16
	AttackLength uint16
	AttackLevel uint16
	FadeLength uint16
	FadeLevel uint16
}

type HapticPeriodic struct {
	Type uint16
	Direction HapticDirection
	Length uint32
	Delay uint16
	Button uint16
	Interval uint16
	Period uint16
	Magnitude int16
	Offset int16
	Phase uint16
	AttackLength uint16
	AttackLevel uint16
	FadeLength uint16
	FadeLevel uint16
}

type HapticCondition struct {
	Type uint16
	Direction HapticDirection
	Length uint32
	Delay uint16
	Button uint16
	Interval uint16
	RightSat [3]uint16
	LeftSat [3]uint16
	RightCoeff [3]int16
	LeftCoeff [3]int16
	Deadband [3]uint16
	Center [3]int16
}

type HapticRamp struct {
	Type uint16
	Direction HapticDirection
	Length uint32
	Delay uint16
	Button uint16
	Interval uint16
	Start int16
	End int16
	AttackLength uint16
	AttackLevel uint16
	FadeLength uint16
	FadeLevel uint16
}

type HapticCustom struct {
	Type uint16
	Direction HapticDirection
	Length uint32
	Delay uint16
	Button uint16
	Interval uint16
	Channels uint8
	Period uint16
	Samples uint16
	Data *uint16
	AttackLength uint16
	AttackLevel uint16
	FadeLength uint16
	FadeLevel uint16
}

type HapticEffect interface {}

func NumHaptics() int {
	return (int) (C.SDL_NumHaptics())
}

func HapticName(device_index int) string {
	_device_index := (C.int) (device_index)
	return (C.GoString) (C.SDL_HapticName(_device_index))
}

func HapticOpen(device_index int) *Haptic {
	_device_index := (C.int) (device_index)
	return (*Haptic) (unsafe.Pointer(C.SDL_HapticOpen(_device_index)))
}

func HapticOpened(device_index int) int {
	_device_index := (C.int) (device_index)
	return (int) (C.SDL_HapticOpened(_device_index))
}

func HapticIndex(haptic *Haptic) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticIndex(_haptic))
}

func MouseIsHaptic() int {
	return (int) (C.SDL_MouseIsHaptic())
}

func HapticOpenFromMouse() *Haptic {
	return (*Haptic) (unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
}

func JoystickIsHaptic(joystick *Joystick) int {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (int) (C.SDL_JoystickIsHaptic(_joystick))
}

func HapticOpenFromJoystick(joystick *Joystick) *Haptic {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (*Haptic) (unsafe.Pointer(C.SDL_HapticOpenFromJoystick(_joystick)))
}

func (haptic *Haptic) Close() {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	C.SDL_HapticClose(_haptic)
}

func (haptic *Haptic) NumEffects() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticNumEffects(_haptic))
}

func (haptic *Haptic) NumEffectsPlaying() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticNumEffectsPlaying(_haptic))
}

func (haptic *Haptic) Query() uint {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (uint) (C.SDL_HapticQuery(_haptic))
}

func (haptic *Haptic) EffectSupported(effect *HapticEffect) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (*C.SDL_HapticEffect) (unsafe.Pointer(effect))
	return (int) (C.SDL_HapticEffectSupported(_haptic, _effect))
}

func (haptic *Haptic) NewEffect(effect *HapticEffect) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (*C.SDL_HapticEffect) (unsafe.Pointer(effect))
	return (int) (C.SDL_HapticNewEffect(_haptic, _effect))
}

func (haptic *Haptic) UpdateEffect(effect int, data *HapticEffect) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (C.int) (effect)
	_data := (*C.SDL_HapticEffect) (unsafe.Pointer(data))
	return (int) (C.SDL_HapticUpdateEffect(_haptic, _effect, _data))
}

func (haptic *Haptic) RunEffect(effect int, iterations uint32) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (C.int) (effect)
	_iterations := (C.Uint32) (iterations)
	return (int) (C.SDL_HapticRunEffect(_haptic, _effect, _iterations))
}

func (haptic *Haptic) StopEffect(effect int) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (C.int) (effect)
	return (int) (C.SDL_HapticStopEffect(_haptic, _effect))
}

func (haptic *Haptic) DestroyEffect(effect int) {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (C.int) (effect)
	C.SDL_HapticDestroyEffect(_haptic, _effect)
}

func (haptic *Haptic) GetEffectStatus(effect int) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_effect := (C.int) (effect)
	return (int) (C.SDL_HapticGetEffectStatus(_haptic, _effect))
}

func (haptic *Haptic) SetGain(gain int) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_gain := (C.int) (gain)
	return (int) (C.SDL_HapticSetGain(_haptic, _gain))
}

func (haptic *Haptic) SetAutocenter(autocenter int) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_autocenter := (C.int) (autocenter)
	return (int) (C.SDL_HapticSetAutocenter(_haptic, _autocenter))
}

func (haptic *Haptic) Pause() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticPause(_haptic))
}

func (haptic *Haptic) Unpause() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticUnpause(_haptic))
}

func (haptic *Haptic) StopAll() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticStopAll(_haptic))
}

func (haptic *Haptic) RumbleSupported() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticRumbleSupported(_haptic))
}

func (haptic *Haptic) RumbleInit() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticRumbleInit(_haptic))
}

func (haptic *Haptic) RumblePlay(strength float32, length uint32) int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	_strength := (C.float) (strength)
	_length := (C.Uint32) (length)
	return (int) (C.SDL_HapticRumblePlay(_haptic, _strength, _length))
}

func (haptic *Haptic) RumbleStop() int {
	_haptic := (*C.SDL_Haptic) (unsafe.Pointer(haptic))
	return (int) (C.SDL_HapticRumbleStop(_haptic))
}
