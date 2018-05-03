package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Haptic effects.
// (https://wiki.libsdl.org/SDL_HapticEffect)
const (
	HAPTIC_CONSTANT     = C.SDL_HAPTIC_CONSTANT     // constant haptic effect
	HAPTIC_SINE         = C.SDL_HAPTIC_SINE         // periodic haptic effect that simulates sine waves
	HAPTIC_LEFTRIGHT    = C.SDL_HAPTIC_LEFTRIGHT    // haptic effect for direct control over high/low frequency motors
	HAPTIC_TRIANGLE     = C.SDL_HAPTIC_TRIANGLE     // periodic haptic effect that simulates triangular waves
	HAPTIC_SAWTOOTHUP   = C.SDL_HAPTIC_SAWTOOTHUP   // periodic haptic effect that simulates saw tooth up waves
	HAPTIC_SAWTOOTHDOWN = C.SDL_HAPTIC_SAWTOOTHDOWN // periodic haptic effect that simulates saw tooth down waves
	HAPTIC_RAMP         = C.SDL_HAPTIC_RAMP         // ramp haptic effect
	HAPTIC_SPRING       = C.SDL_HAPTIC_SPRING       // condition haptic effect that simulates a spring.  Effect is based on the axes position
	HAPTIC_DAMPER       = C.SDL_HAPTIC_DAMPER       // condition haptic effect that simulates dampening.  Effect is based on the axes velocity
	HAPTIC_INERTIA      = C.SDL_HAPTIC_INERTIA      // condition haptic effect that simulates inertia.  Effect is based on the axes acceleration
	HAPTIC_FRICTION     = C.SDL_HAPTIC_FRICTION     // condition haptic effect that simulates friction.  Effect is based on the axes movement
	HAPTIC_CUSTOM       = C.SDL_HAPTIC_CUSTOM       // user defined custom haptic effect
	HAPTIC_GAIN         = C.SDL_HAPTIC_GAIN         // device supports setting the global gain
	HAPTIC_AUTOCENTER   = C.SDL_HAPTIC_AUTOCENTER   // device supports setting autocenter
	HAPTIC_STATUS       = C.SDL_HAPTIC_STATUS       // device can be queried for effect status
	HAPTIC_PAUSE        = C.SDL_HAPTIC_PAUSE        // device can be paused
	//HAPTIC_SQUARE = C.SDL_HAPTIC_SQUARE (back in SDL 2.1)
)

// Direction encodings.
// (https://wiki.libsdl.org/SDL_HapticDirection)
const (
	HAPTIC_POLAR     = C.SDL_HAPTIC_POLAR     // uses polar coordinates for the direction
	HAPTIC_CARTESIAN = C.SDL_HAPTIC_CARTESIAN // uses cartesian coordinates for the direction
	HAPTIC_SPHERICAL = C.SDL_HAPTIC_SPHERICAL // uses spherical coordinates for the direction
	HAPTIC_INFINITY  = C.SDL_HAPTIC_INFINITY  // used to play a device an infinite number of times
)

// Haptic identifies an SDL haptic.
// (https://wiki.libsdl.org/CategoryForceFeedback)
type Haptic C.SDL_Haptic

// HapticDirection contains a haptic direction.
// (https://wiki.libsdl.org/SDL_HapticDirection)
type HapticDirection struct {
	Type byte     // the type of encoding
	Dir  [3]int32 // the encoded direction
}

// HapticConstant contains a template for a constant effect.
// (https://wiki.libsdl.org/SDL_HapticConstant)
type HapticConstant struct {
	Type         uint16          // HAPTIC_CONSTANT
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Level        int16           // strength of the constant effect
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

func (he *HapticConstant) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticPeriodic contains a template for a periodic effect.
// (https://wiki.libsdl.org/SDL_HapticPeriodic)
type HapticPeriodic struct {
	Type         uint16          // HAPTIC_SINE, HAPTIC_LEFTRIGHT, HAPTIC_TRIANGLE, HAPTIC_SAWTOOTHUP, HAPTIC_SAWTOOTHDOWN
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Period       uint16          // period of the wave
	Magnitude    int16           // peak value; if negative, equivalent to 180 degrees extra phase shift
	Offset       int16           // mean value of the wave
	Phase        uint16          // positive phase shift given by hundredth of a degree
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

func (he *HapticPeriodic) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticCondition contains a template for a condition effect.
// (https://wiki.libsdl.org/SDL_HapticCondition)
type HapticCondition struct {
	Type       uint16          // HAPTIC_SPRING, HAPTIC_DAMPER, HAPTIC_INERTIA, HAPTIC_FRICTION
	Direction  HapticDirection // direction of the effect - not used at the moment
	Length     uint32          // duration of the effect
	Delay      uint16          // delay before starting the effect
	Button     uint16          // button that triggers the effect
	Interval   uint16          // how soon it can be triggered again after button
	RightSat   [3]uint16       // level when joystick is to the positive side; max 0xFFFF
	LeftSat    [3]uint16       // level when joystick is to the negative side; max 0xFFFF
	RightCoeff [3]int16        // how fast to increase the force towards the positive side
	LeftCoeff  [3]int16        // how fast to increase the force towards the negative side
	Deadband   [3]uint16       // size of the dead zone; max 0xFFFF: whole axis-range when 0-centered
	Center     [3]int16        // position of the dead zone
}

func (he *HapticCondition) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticRamp contains a template for a ramp effect.
// (https://wiki.libsdl.org/SDL_HapticRamp)
type HapticRamp struct {
	Type         uint16          // HAPTIC_RAMP
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Start        int16           // beginning strength level
	End          int16           // ending strength level
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

func (he *HapticRamp) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticLeftRight contains a template for a left/right effect.
// (https://wiki.libsdl.org/SDL_HapticLeftRight)
type HapticLeftRight struct {
	Type           uint16 // HAPTIC_LEFTRIGHT
	Length         uint32 // duration of the effect
	LargeMagnitude uint16 // control of the large controller motor
	SmallMagnitude uint16 // control of the small controller motor
}

func (he *HapticLeftRight) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticCustom contains a template for a custom effect.
// (https://wiki.libsdl.org/SDL_HapticCustom)
type HapticCustom struct {
	Type         uint16          // SDL_HAPTIC_CUSTOM
	Direction    HapticDirection // direction of the effect
	Length       uint32          // duration of the effect
	Delay        uint16          // delay before starting the effect
	Button       uint16          // button that triggers the effect
	Interval     uint16          // how soon it can be triggered again after button
	Channels     uint8           // axes to use, minimum of 1
	Period       uint16          // sample periods
	Samples      uint16          // amount of samples
	Data         *uint16         // should contain channels*samples items
	AttackLength uint16          // duration of the attack
	AttackLevel  uint16          // level at the start of the attack
	FadeLength   uint16          // duration of the fade
	FadeLevel    uint16          // level at the end of the fade
}

func (he *HapticCustom) cHapticEffect() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

// HapticEffect union that contains the generic template for any haptic effect.
// (https://wiki.libsdl.org/SDL_HapticEffect)
type HapticEffect interface {
	cHapticEffect() *C.SDL_HapticEffect
}

func (h *Haptic) cptr() *C.SDL_Haptic {
	return (*C.SDL_Haptic)(unsafe.Pointer(h))
}

// NumHaptics returns the number of haptic devices attached to the system.
// (https://wiki.libsdl.org/SDL_NumHaptics)
func NumHaptics() (int, error) {
	i := int(C.SDL_NumHaptics())
	return i, errorFromInt(i)
}

// HapticName returns the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL_HapticName)
func HapticName(index int) (string, error) {
	name := C.SDL_HapticName(C.int(index))
	if name == nil {
		return "", GetError()
	}
	return C.GoString(name), nil
}

// HapticOpen opens a haptic device for use.
// (https://wiki.libsdl.org/SDL_HapticOpen)
func HapticOpen(index int) (*Haptic, error) {
	haptic := (*Haptic)(unsafe.Pointer(C.SDL_HapticOpen(C.int(index))))
	if haptic == nil {
		return nil, GetError()
	}
	return haptic, nil
}

// HapticOpened reports whether the haptic device at the designated index has been opened.
// (https://wiki.libsdl.org/SDL_HapticOpened)
func HapticOpened(index int) (bool, error) {
	ret := int(C.SDL_HapticOpened(C.int(index)))
	if ret == 0 {
		return false, GetError()
	}
	return ret == 1, nil
}

// HapticIndex returns the index of a haptic device.
// (https://wiki.libsdl.org/SDL_HapticIndex)
func HapticIndex(h *Haptic) (int, error) {
	i := int(C.SDL_HapticIndex(h.cptr()))
	return i, errorFromInt(i)
}

// MouseIsHaptic reports whether or not the current mouse has haptic capabilities.
// (https://wiki.libsdl.org/SDL_MouseIsHaptic)
func MouseIsHaptic() (bool, error) {
	ret := int(C.SDL_MouseIsHaptic())
	return ret == C.SDL_TRUE, errorFromInt(ret)
}

// HapticOpenFromMouse open a haptic device from the current mouse.
// (https://wiki.libsdl.org/SDL_HapticOpenFromMouse)
func HapticOpenFromMouse() (*Haptic, error) {
	haptic := (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
	if haptic == nil {
		return nil, GetError()
	}
	return haptic, nil
}

// JoystickIsHaptic reports whether a joystick has haptic features.
// (https://wiki.libsdl.org/SDL_JoystickIsHaptic)
func JoystickIsHaptic(joy *Joystick) (bool, error) {
	ret := int(C.SDL_JoystickIsHaptic(joy.cptr()))
	return ret == C.SDL_TRUE, errorFromInt(ret)
}

// HapticOpenFromJoystick opens a haptic device for use from a joystick device.
// (https://wiki.libsdl.org/SDL_HapticOpenFromJoystick)
func HapticOpenFromJoystick(joy *Joystick) (*Haptic, error) {
	haptic := (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromJoystick(joy.cptr())))
	if haptic == nil {
		return nil, GetError()
	}
	return haptic, nil
}

// Close closes a haptic device previously opened with HapticOpen().
// (https://wiki.libsdl.org/SDL_HapticClose)
func (h *Haptic) Close() {
	C.SDL_HapticClose(h.cptr())
}

// NumAxes returns the number of haptic axes the device has.
// (https://wiki.libsdl.org/SDL_HapticNumAxes)
func (h *Haptic) NumAxes() (int, error) {
	i := int(C.SDL_HapticNumAxes(h.cptr()))
	return i, errorFromInt(i)
}

// NumEffects returns the number of effects a haptic device can store.
// (https://wiki.libsdl.org/SDL_HapticNumEffects)
func (h *Haptic) NumEffects() (int, error) {
	i := int(C.SDL_HapticNumEffects(h.cptr()))
	return i, errorFromInt(i)
}

// NumEffectsPlaying returns the number of effects a haptic device can play at the same time.
// (https://wiki.libsdl.org/SDL_HapticNumEffectsPlaying)
func (h *Haptic) NumEffectsPlaying() (int, error) {
	i := int(C.SDL_HapticNumEffectsPlaying(h.cptr()))
	return i, errorFromInt(i)
}

// Query returns haptic device's supported features in bitwise manner.
// (https://wiki.libsdl.org/SDL_HapticQuery)
func (h *Haptic) Query() (uint32, error) {
	i := uint32(C.SDL_HapticQuery(h.cptr()))
	if i == 0 {
		return 0, GetError()
	}
	return i, nil
}

// EffectSupported reports whether an effect is supported by a haptic device.
// Pass pointer to a Haptic struct (Constant|Periodic|Condition|Ramp|LeftRight|Custom) instead of HapticEffect union.
// (https://wiki.libsdl.org/SDL_HapticEffectSupported)
func (h *Haptic) EffectSupported(he HapticEffect) (bool, error) {
	ret := int(C.SDL_HapticEffectSupported(
		h.cptr(),
		he.cHapticEffect()))
	return ret == C.SDL_TRUE, errorFromInt(ret)
}

// NewEffect creates a new haptic effect on a specified device.
// Pass pointer to a Haptic struct (Constant|Periodic|Condition|Ramp|LeftRight|Custom) instead of HapticEffect union.
// (https://wiki.libsdl.org/SDL_HapticNewEffect)
func (h *Haptic) NewEffect(he HapticEffect) (int, error) {
	ret := int(C.SDL_HapticNewEffect(
		h.cptr(),
		he.cHapticEffect()))
	return ret, errorFromInt(ret)
}

// UpdateEffect updates the properties of an effect.
// Pass pointer to a Haptic struct (Constant|Periodic|Condition|Ramp|LeftRight|Custom) instead of HapticEffect union.
// (https://wiki.libsdl.org/SDL_HapticUpdateEffect)
func (h *Haptic) UpdateEffect(effect int, data HapticEffect) error {
	return errorFromInt(int(
		C.SDL_HapticUpdateEffect(
			h.cptr(),
			C.int(effect),
			data.cHapticEffect())))
}

// RunEffect runs the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL_HapticRunEffect)
func (h *Haptic) RunEffect(effect int, iterations uint32) error {
	return errorFromInt(int(
		C.SDL_HapticRunEffect(
			h.cptr(),
			C.int(effect),
			C.Uint32(iterations))))
}

// StopEffect stops the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL_HapticStopEffect)
func (h *Haptic) StopEffect(effect int) error {
	return errorFromInt(int(
		C.SDL_HapticStopEffect(h.cptr(), C.int(effect))))
}

// DestroyEffect destroys a haptic effect on the device.
// (https://wiki.libsdl.org/SDL_HapticDestroyEffect)
func (h *Haptic) DestroyEffect(effect int) {
	C.SDL_HapticDestroyEffect(h.cptr(), C.int(effect))
}

// GetEffectStatus returns the status of the current effect on the specified haptic device.
// (https://wiki.libsdl.org/SDL_HapticGetEffectStatus)
func (h *Haptic) GetEffectStatus(effect int) (int, error) {
	i := int(C.SDL_HapticGetEffectStatus(h.cptr(), C.int(effect)))
	return i, errorFromInt(i)

}

// SetGain sets the global gain of the specified haptic device.
// (https://wiki.libsdl.org/SDL_HapticSetGain)
func (h *Haptic) SetGain(gain int) error {
	return errorFromInt(int(
		C.SDL_HapticSetGain(h.cptr(), C.int(gain))))
}

// SetAutocenter sets the global autocenter of the device.
// (https://wiki.libsdl.org/SDL_HapticSetAutocenter)
func (h *Haptic) SetAutocenter(autocenter int) error {
	return errorFromInt(int(
		C.SDL_HapticSetAutocenter(h.cptr(), C.int(autocenter))))
}

// Pause pauses a haptic device.
// (https://wiki.libsdl.org/SDL_HapticPause)
func (h *Haptic) Pause() error {
	return errorFromInt(int(
		C.SDL_HapticPause(h.cptr())))
}

// Unpause unpauses a haptic device.
// (https://wiki.libsdl.org/SDL_HapticUnpause)
func (h *Haptic) Unpause() error {
	return errorFromInt(int(
		C.SDL_HapticUnpause(h.cptr())))
}

// StopAll stops all the currently playing effects on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticStopAll)
func (h *Haptic) StopAll() error {
	return errorFromInt(int(
		C.SDL_HapticStopAll(h.cptr())))
}

// RumbleSupported reports whether rumble is supported on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumbleSupported)
func (h *Haptic) RumbleSupported() (bool, error) {
	ret := int(C.SDL_HapticRumbleSupported(h.cptr()))
	return ret == C.SDL_TRUE, errorFromInt(ret)
}

// RumbleInit initializes the haptic device for simple rumble playback.
// (https://wiki.libsdl.org/SDL_HapticRumbleInit)
func (h *Haptic) RumbleInit() error {
	return errorFromInt(int(
		C.SDL_HapticRumbleInit(h.cptr())))
}

// RumblePlay runs a simple rumble effect on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumblePlay)
func (h *Haptic) RumblePlay(strength float32, length uint32) error {
	return errorFromInt(int(
		C.SDL_HapticRumblePlay(h.cptr(), C.float(strength), C.Uint32(length))))
}

// RumbleStop stops the simple rumble on a haptic device.
// (https://wiki.libsdl.org/SDL_HapticRumbleStop)
func (h *Haptic) RumbleStop() error {
	return errorFromInt(int(
		C.SDL_HapticRumbleStop(h.cptr())))
}
