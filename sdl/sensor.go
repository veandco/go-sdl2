package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,9))
typedef struct _SDL_Sensor SDL_Sensor;

typedef Sint32 SDL_SensorID;

typedef enum
{
    SDL_SENSOR_INVALID = -1,
    SDL_SENSOR_UNKNOWN,
    SDL_SENSOR_ACCEL,
    SDL_SENSOR_GYRO
} SDL_SensorType;


#if defined(WARN_OUTDATED)
#pragma message("SDL_NumSensors is not supported before SDL 2.0.9")
#endif

static inline int SDL_NumSensors()
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetDeviceName is not supported before SDL 2.0.9")
#endif

static const char * SDL_SensorGetDeviceName(int device_index)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetDeviceType is not supported before SDL 2.0.9")
#endif

static SDL_SensorType SDL_SensorGetDeviceType(int device_index)
{
	return SDL_SENSOR_INVALID;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetDeviceNonPortableType is not supported before SDL 2.0.9")
#endif

static int SDL_SensorGetDeviceNonPortableType(int device_index)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetDeviceInstanceID is not supported before SDL 2.0.9")
#endif

static SDL_SensorID SDL_SensorGetDeviceInstanceID(int device_index)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorOpen is not supported before SDL 2.0.9")
#endif

static SDL_Sensor * SDL_SensorOpen(int device_index)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorFromInstanceID is not supported before SDL 2.0.9")
#endif

static SDL_Sensor * SDL_SensorFromInstanceID(SDL_SensorID instance_id)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetName is not supported before SDL 2.0.9")
#endif

static const char * SDL_SensorGetName(SDL_Sensor *sensor)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetType is not supported before SDL 2.0.9")
#endif

static SDL_SensorType SDL_SensorGetType(SDL_Sensor *sensor)
{
	return SDL_SENSOR_INVALID;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetNonPortableType is not supported before SDL 2.0.9")
#endif

static int SDL_SensorGetNonPortableType(SDL_Sensor *sensor)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetInstanceID is not supported before SDL 2.0.9")
#endif

static SDL_SensorID SDL_SensorGetInstanceID(SDL_Sensor *sensor)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorGetData is not supported before SDL 2.0.9")
#endif

static int SDL_SensorGetData(SDL_Sensor *sensor, float *data, int num_values)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorClose is not supported before SDL 2.0.9")
#endif

static void SDL_SensorClose(SDL_Sensor *sensor)
{
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SensorUpdate is not supported before SDL 2.0.9")
#endif

static void SDL_SensorUpdate()
{
}
#endif
*/
import "C"
import "unsafe"

const (
	STANDARD_GRAVITY = 9.80665
)

// The different sensors defined by SDL
//
// Additional sensors may be available, using platform dependent semantics.
//
// Here are the additional Android sensors:
// https://developer.android.com/reference/android/hardware/SensorEvent.html#values
const (
	SENSOR_INVALID SensorType = C.SDL_SENSOR_INVALID // Returned for an invalid sensor
	SENSOR_UNKNOWN SensorType = C.SDL_SENSOR_UNKNOWN // Unknown sensor type
	SENSOR_ACCEL   SensorType = C.SDL_SENSOR_ACCEL   // Accelerometer
	SENSOR_GYRO    SensorType = C.SDL_SENSOR_GYRO    // Gyroscope
)

type Sensor C.SDL_Sensor
type SensorID int32
type SensorType int

// NumSensors counts the number of sensors attached to the system right now
// (https://wiki.libsdl.org/SDL_NumSensors)
func NumSensors() int {
	return int(C.SDL_NumSensors())
}

// SensorGetDeviceName gets the implementation dependent name of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor name, or empty string if deviceIndex is out of range.
// (https://wiki.libsdl.org/SDL_SensorGetDeviceName)
func SensorGetDeviceName(deviceIndex int) (name string) {
	name = C.GoString(C.SDL_SensorGetDeviceName(C.int(deviceIndex)))
	return
}

//  SensorGetDeviceType gets the type of a sensor.
//
//  This can be called before any sensors are opened.
//
//  Returns the sensor type, or SDL_SENSOR_INVALID if deviceIndex is out of range.
// (https://wiki.libsdl.org/SDL_SensorGetDeviceType)
func SensorGetDeviceType(deviceIndex int) (typ SensorType) {
	typ = SensorType(C.SDL_SensorGetDeviceType(C.int(deviceIndex)))
	return
}

// SensorGetDeviceNonPortableType gets the platform dependent type of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor platform dependent type, or -1 if deviceIndex is out of range.
// (https://wiki.libsdl.org/SDL_SensorGetDeviceNonPortableType)
func SensorGetDeviceNonPortableType(deviceIndex int) (typ int) {
	typ = int(C.SDL_SensorGetDeviceNonPortableType(C.int(deviceIndex)))
	return
}

// SensorGetDeviceInstanceID gets the instance ID of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor instance ID, or -1 if deviceIndex is out of range.
// (https://wiki.libsdl.org/SDL_SensorGetDeviceInstanceID)
func SensorGetDeviceInstanceID(deviceIndex int) (id SensorID) {
	id = SensorID(C.SDL_SensorGetDeviceInstanceID(C.int(deviceIndex)))
	return
}

// SensorOpen opens a sensor for use.
//
// The index passed as an argument refers to the N'th sensor on the system.
//
// Returns a sensor identifier, or nil if an error occurred.
// (https://wiki.libsdl.org/SDL_SensorOpen)
func SensorOpen(deviceIndex int) (sensor *Sensor) {
	sensor = (*Sensor)(unsafe.Pointer(C.SDL_SensorOpen(C.int(deviceIndex))))
	return
}

// SensorFromInstanceID returns the Sensor associated with an instance id.
// (https://wiki.libsdl.org/SDL_SensorFromInstanceID)
func SensorFromInstanceID(id SensorID) (sensor *Sensor) {
	sensor = (*Sensor)(unsafe.Pointer(C.SDL_SensorFromInstanceID(C.SDL_SensorID(id))))
	return
}

// GetName gets the implementation dependent name of a sensor.
//
// Returns the sensor name, or empty string if the sensor is nil.
// (https://wiki.libsdl.org/SDL_SensorGetName)
func (sensor *Sensor) GetName() (name string) {
	name = C.GoString(C.SDL_SensorGetName((*C.SDL_Sensor)(sensor)))
	return
}

// GetType gets the type of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor type, or SENSOR_INVALID if the sensor is nil.
// (https://wiki.libsdl.org/SDL_SensorGetType)
func (sensor *Sensor) GetType() (typ SensorType) {
	typ = SensorType(C.SDL_SensorGetType((*C.SDL_Sensor)(sensor)))
	return
}

// GetNonPortableType gets the platform dependent type of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor platform dependent type, or -1 if the sensor is nil.
// (https://wiki.libsdl.org/SDL_SensorGetNonPortableType)
func (sensor *Sensor) GetNonPortableType() (typ int) {
	typ = int(C.SDL_SensorGetNonPortableType((*C.SDL_Sensor)(sensor)))
	return
}

// GetInstanceID gets the instance ID of a sensor.
//
// This can be called before any sensors are opened.
//
// Returns the sensor instance ID, or -1 if the sensor is nil.
// (https://wiki.libsdl.org/SDL_SensorGetInstanceID)
func (sensor *Sensor) GetInstanceID() (id SensorID) {
	id = SensorID(C.SDL_SensorGetInstanceID((*C.SDL_Sensor)(sensor)))
	return
}

// GetData gets the current state of an opened sensor.
//
// The number of values and interpretation of the data is sensor dependent.
// (https://wiki.libsdl.org/SDL_SensorGetData)
func (sensor *Sensor) GetData(data []float32) (err error) {
	if data == nil {
		return nil
	}
	_data := (*C.float)(unsafe.Pointer(&data[0]))
	_numValues := C.int(len(data))
	err = errorFromInt(int(C.SDL_SensorGetData((*C.SDL_Sensor)(sensor), _data, _numValues)))
	return
}

// Close closes a sensor previously opened with SensorOpen()
// (https://wiki.libsdl.org/SDL_SensorClose)
func (sensor *Sensor) Close() {
	C.SDL_SensorClose((*C.SDL_Sensor)(sensor))
}

// SensorUpdate updates the current state of the open sensors.
//
// This is called automatically by the event loop if sensor events are enabled.
//
// This needs to be called from the thread that initialized the sensor subsystem.
// (https://wiki.libsdl.org/SDL_SensorUpdate)
func SensorUpdate() {
	C.SDL_SensorUpdate()
}
