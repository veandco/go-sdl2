package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_hid_init is not supported before SDL 2.0.18")
#pragma message("SDL_hid_exit is not supported before SDL 2.0.18")
#pragma message("SDL_hid_device_change_count is not supported before SDL 2.0.18")
#pragma message("SDL_hid_enumerate is not supported before SDL 2.0.18")
#pragma message("SDL_hid_free_enumeration is not supported before SDL 2.0.18")
#pragma message("SDL_hid_open is not supported before SDL 2.0.18")
#pragma message("SDL_hid_open_path is not supported before SDL 2.0.18")
#pragma message("SDL_hid_write is not supported before SDL 2.0.18")
#pragma message("SDL_hid_read_timeout is not supported before SDL 2.0.18")
#pragma message("SDL_hid_read is not supported before SDL 2.0.18")
#pragma message("SDL_hid_set_nonblocking is not supported before SDL 2.0.18")
#pragma message("SDL_hid_send_feature_report is not supported before SDL 2.0.18")
#pragma message("SDL_hid_get_feature_report is not supported before SDL 2.0.18")
#pragma message("SDL_hid_close is not supported before SDL 2.0.18")
#pragma message("SDL_hid_get_manufacturer_string is not supported before SDL 2.0.18")
#pragma message("SDL_hid_get_product_string is not supported before SDL 2.0.18")
#pragma message("SDL_hid_get_serial_number_string is not supported before SDL 2.0.18")
#pragma message("SDL_hid_get_indexed_string is not supported before SDL 2.0.18")
#pragma message("SDL_hid_ble_scan is not supported before SDL 2.0.18")
#endif

struct SDL_hid_device_;
typedef struct SDL_hid_device_ SDL_hid_device; // opaque hidapi structure

// hidapi info structure
// Information about a connected HID device
typedef struct SDL_hid_device_info
{
    // Platform-specific device path
    char *path;

    // Device Vendor ID
    unsigned short vendor_id;

    // Device Product ID
    unsigned short product_id;

    // Serial Number
    wchar_t *serial_number;

    // Device Release Number in binary-coded decimal, also known as Device Version Number
    unsigned short release_number;

    // Manufacturer String
    wchar_t *manufacturer_string;

    // Product string
    wchar_t *product_string;

    // Usage Page for this Device/Interface (Windows/Mac only).
    unsigned short usage_page;

    // Usage for this Device/Interface (Windows/Mac only).
    unsigned short usage;

    // The USB interface which this logical device represents.
    // Valid on both Linux implementations in all cases.
    // Valid on the Windows implementation only if the device
    // contains more than one interface.
    int interface_number;

    // Additional information about the USB interface.
    // Valid on libusb and Android implementations.
    int interface_class;
    int interface_subclass;
    int interface_protocol;

    // Pointer to the next device
    struct SDL_hid_device_info *next;
} SDL_hid_device_info;

static int SDL_hid_init(void)
{
	return -1;
}

static int SDL_hid_exit(void)
{
	return -1;
}

static Uint32 SDL_hid_device_change_count(void)
{
	return 0;
}

static SDL_hid_device_info * SDL_hid_enumerate(unsigned short vendor_id, unsigned short product_id)
{
	return NULL;
}

static void SDL_hid_free_enumeration(SDL_hid_device_info *devs)
{
}

static SDL_hid_device * SDL_hid_open(unsigned short vendor_id, unsigned short product_id, const wchar_t *serial_number)
{
	return NULL;
}

static SDL_hid_device * SDL_hid_open_path(const char *path, int bExclusive)
{
	return NULL;
}

static int SDL_hid_write(SDL_hid_device *dev, const unsigned char *data, size_t length)
{
	return -1;
}

static int SDL_hid_read_timeout(SDL_hid_device *dev, unsigned char *data, size_t length, int milliseconds)
{
	return -1;
}

static int SDL_hid_read(SDL_hid_device *dev, unsigned char *data, size_t length)
{
	return -1;
}

static int SDL_hid_set_nonblocking(SDL_hid_device *dev, int nonblock)
{
	return -1;
}

static int SDL_hid_send_feature_report(SDL_hid_device *dev, const unsigned char *data, size_t length)
{
	return -1;
}

static int SDL_hid_get_feature_report(SDL_hid_device *dev, unsigned char *data, size_t length)
{
	return -1;
}

static void SDL_hid_close(SDL_hid_device *dev)
{
}

static int SDL_hid_get_manufacturer_string(SDL_hid_device *dev, wchar_t *string, size_t maxlen)
{
	return -1;
}

static int SDL_hid_get_product_string(SDL_hid_device *dev, wchar_t *string, size_t maxlen)
{
	return -1;
}

static int SDL_hid_get_serial_number_string(SDL_hid_device *dev, wchar_t *string, size_t maxlen)
{
	return -1;
}

static int SDL_hid_get_indexed_string(SDL_hid_device *dev, int string_index, wchar_t *string, size_t maxlen)
{
	return -1;
}

static void SDL_hid_ble_scan(SDL_bool active)
{
}

#endif
*/
import "C"
import "unsafe"

type HIDDevice C.SDL_hid_device

type HIDDeviceInfo struct {
	path               *C.char
	VendorID           uint16
	ProductID          uint16
	SerialNumber       *C.wchar_t
	ReleaseNumber      uint16
	ManufacturerString *C.wchar_t
	ProductString      *C.wchar_t
	UsagePage          uint16
	Usage              uint16
	InterfaceNumber    int32
	InterfaceClass     int32
	InterfaceSubclass  int32
	InterfaceProtocol  int32
	next               *HIDDeviceInfo
}

func (info *HIDDeviceInfo) Path() string {
	return C.GoString(info.path)
}

// HIDInit initializes the HIDAPI library.
// (https://wiki.libsdl.org/SDL_hid_init)
func HIDInit() (err error) {
	ret := C.SDL_hid_init()
	return errorFromInt(int(ret))
}

// HIDExit finalizes the HIDAPI library.
// (https://wiki.libsdl.org/SDL_hid_exit)
func HIDExit() (err error) {
	ret := C.SDL_hid_exit()
	return errorFromInt(int(ret))
}

// HIDDeviceChangeCount checks to see if devices may have been added or removed.
// (https://wiki.libsdl.org/SDL_hid_device_change_count)
func HIDDeviceChangeCount() (n uint32) {
	return uint32(C.SDL_hid_device_change_count())
}

// HIDEnumerate enumerates the HID devices.
// (https://wiki.libsdl.org/SDL_hid_enumerate)
func HIDEnumerate(vendorID, productID uint16) (info *HIDDeviceInfo) {
	_vendorID := C.Uint16(vendorID)
	_productID := C.Uint16(productID)
	info = (*HIDDeviceInfo)(unsafe.Pointer(C.SDL_hid_enumerate(_vendorID, _productID)))
	return
}

// HIDFreeEnumeration frees an enumeration Linked List.
// (https://wiki.libsdl.org/SDL_hid_free_enumeration)
func HIDFreeEnumeration(info *HIDDeviceInfo) {
	_info := (*C.SDL_hid_device_info)(unsafe.Pointer(info))
	C.SDL_hid_free_enumeration(_info)
}

// HIDOpen opens a HID device using a Vendor ID (VID), Product ID (PID) and optionally a serial number.
// (https://wiki.libsdl.org/SDL_hid_open)
func HIDOpen(vendorID, productID uint16, _serialNumber *C.wchar_t) (device *HIDDevice) {
	_vendorID := C.Uint16(vendorID)
	_productID := C.Uint16(productID)
	_device := C.SDL_hid_open(_vendorID, _productID, _serialNumber)
	return (*HIDDevice)(_device)
}

// HIDOpenPath opens a HID device by its path name.
// (https://wiki.libsdl.org/SDL_hid_open_path)
func HIDOpenPath(path string, exclusive bool) (device *HIDDevice) {
	_path := C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	_exclusive := C.int(0)
	if exclusive {
		_exclusive = C.int(1)
	}
	_device := C.SDL_hid_open_path(_path, _exclusive)
	return (*HIDDevice)(_device)
}

// Write writes an Output report to a HID device.
// (https://wiki.libsdl.org/SDL_hid_write)
func (device *HIDDevice) Write(data []byte) (n int, err error) {
	_length := C.size_t(len(data))
	_data := (*C.uchar)(unsafe.Pointer(&data[0]))
	_device := (*C.SDL_hid_device)(device)
	n = int(C.SDL_hid_write(_device, _data, _length))
	return n, errorFromInt(n)
}

// ReadTimeout reads an Input report from a HID device with timeout.
// (https://wiki.libsdl.org/SDL_hid_read_timeout)
func (device *HIDDevice) ReadTimeout(data []byte, milliseconds int) (n int, err error) {
	_length := C.size_t(len(data))
	_data := (*C.uchar)(unsafe.Pointer(&data[0]))
	_device := (*C.SDL_hid_device)(device)
	_milliseconds := C.int(milliseconds)
	n = int(C.SDL_hid_read_timeout(_device, _data, _length, _milliseconds))
	return n, errorFromInt(n)
}

// Read an Input report from a HID device.
// (https://wiki.libsdl.org/SDL_hid_read)
func (device *HIDDevice) Read(data []byte) (n int, err error) {
	_length := C.size_t(len(data))
	_data := (*C.uchar)(unsafe.Pointer(&data[0]))
	_device := (*C.SDL_hid_device)(device)
	n = int(C.SDL_hid_read(_device, _data, _length))
	return n, errorFromInt(n)
}

// SetNonBlocking sets the device handle to be non-blocking.
// (https://wiki.libsdl.org/SDL_hid_set_nonblocking)
func (device *HIDDevice) SetNonBlocking(nonblock bool) (err error) {
	_device := (*C.SDL_hid_device)(device)
	_nonblock := C.int(Btoi(nonblock))
	return errorFromInt(int(C.SDL_hid_set_nonblocking(_device, _nonblock)))
}

// SendFeatureReport sends a Feature report to the device.
// (https://wiki.libsdl.org/SDL_hid_send_feature_report)
func (device *HIDDevice) SendFeatureReport(data []byte) (n int, err error) {
	_length := C.size_t(len(data))
	_data := (*C.uchar)(unsafe.Pointer(&data[0]))
	_device := (*C.SDL_hid_device)(device)
	n = int(C.SDL_hid_send_feature_report(_device, _data, _length))
	return n, errorFromInt(n)
}

// GetFeatureReport gets a feature report from a HID device.
// (https://wiki.libsdl.org/SDL_hid_get_feature_report)
func (device *HIDDevice) GetFeatureReport(data []byte) (n int, err error) {
	_length := C.size_t(len(data))
	_data := (*C.uchar)(unsafe.Pointer(&data[0]))
	_device := (*C.SDL_hid_device)(device)
	n = int(C.SDL_hid_get_feature_report(_device, _data, _length))
	return n, errorFromInt(n)
}

// Close closes a HID device.
// (https://wiki.libsdl.org/SDL_hid_close)
func (device *HIDDevice) Close() {
	_device := (*C.SDL_hid_device)(device)
	C.SDL_hid_close(_device)
}

// GetManufacturerString gets The Manufacturer String from a HID device.
// (https://wiki.libsdl.org/SDL_hid_get_manufacturer_string)
func (device *HIDDevice) GetManufacturerString(_str *C.wchar_t, _maxlen C.size_t) (err error) {
	_device := (*C.SDL_hid_device)(device)
	return errorFromInt(int(C.SDL_hid_get_manufacturer_string(_device, _str, _maxlen)))
}

// GetProductString gets The Product String from a HID device.
// (https://wiki.libsdl.org/SDL_hid_get_product_string)
func (device *HIDDevice) GetProductString(_str *C.wchar_t, _maxlen C.size_t) (err error) {
	_device := (*C.SDL_hid_device)(device)
	return errorFromInt(int(C.SDL_hid_get_product_string(_device, _str, _maxlen)))
}

// GetSerialNumberString gets The SerialNumber String from a HID device.
// (https://wiki.libsdl.org/SDL_hid_get_serial_number_string)
func (device *HIDDevice) GetSerialNumberString(_str *C.wchar_t, _maxlen C.size_t) (err error) {
	_device := (*C.SDL_hid_device)(device)
	return errorFromInt(int(C.SDL_hid_get_serial_number_string(_device, _str, _maxlen)))
}

// GetIndexedString gets a string from a HID device, based on its string index.
// (https://wiki.libsdl.org/SDL_hid_get_indexed_string)
func (device *HIDDevice) GetIndexedString(index int, _str *C.wchar_t, _maxlen C.size_t) (err error) {
	_device := (*C.SDL_hid_device)(device)
	_index := C.int(index)
	return errorFromInt(int(C.SDL_hid_get_indexed_string(_device, _index, _str, _maxlen)))
}

// HIDBLEScan starts or stops a BLE scan on iOS and tvOS to pair Steam Controllers.
// (https://wiki.libsdl.org/SDL_hid_ble_scan)
func (device *HIDDevice) HIDBLEScan(active bool) {
	_active := C.SDL_bool(Btoi(active))
	C.SDL_hid_ble_scan(_active)
}
