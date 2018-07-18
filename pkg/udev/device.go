// +build linux,cgo

package udev // import "github.com/shovanmaity/libudev-go-wrapper/pkg/udev"

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdevDevice wraps a libudev device object
type UdevDevice struct {
	device *C.struct_udev_device
}

// newUdevDevice is a private helper function and returns a pointer to a new device.
func newUdevDevice(ptr *C.struct_udev_device) *UdevDevice {
	if ptr == nil {
		return nil
	}
	ud := &UdevDevice{
		device: ptr,
	}
	return ud
}

// UnrefDeviceUdev frees udev structure.
func (ud *UdevDevice) UnrefDeviceUdev() {
	C.udev_device_unref(ud.device)
}

// GetPropertyValue retrieves the value of a device property
func (ud *UdevDevice) GetPropertyValue(key string) string {
	k := C.CString(key)
	defer freeCharPtr(k)
	return C.GoString(C.udev_device_get_property_value(ud.device, k))
}

// GetDevtype returns type of the disk
func (ud *UdevDevice) GetDevtype() string {
	return C.GoString(C.udev_device_get_devtype(ud.device))
}

// GetDevnode returns the device node file name belonging to the udev device.
func (ud *UdevDevice) GetDevnode() string {
	return C.GoString(C.udev_device_get_devnode(ud.device))
}

// GetAction returns device action like add when it is attached remove
// when the device is removed
func (ud *UdevDevice) GetAction() string {
	return C.GoString(C.udev_device_get_action(ud.device))
}
