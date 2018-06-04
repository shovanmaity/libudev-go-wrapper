// +build linux,cgo
package udev

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
func newUdevDevice(ptr *C.struct_udev_device) (ud *UdevDevice) {
	if ptr == nil {
		return nil
	}
	ud = &UdevDevice{
		device: ptr,
	}
	return
}

// UnrefDeviceUdev frees udev structure.
func (ud *UdevDevice) UnrefDeviceUdev() {
	C.udev_device_unref(ud.device)
}

// PropertyValue retrieves the value of a device property
func (ud *UdevDevice) PropertyValue(key string) string {
	k := C.CString(key)
	defer freeCharPtr(k)
	return C.GoString(C.udev_device_get_property_value(ud.device, k))
}

// Devtype returns the devtype string of the udev device.
func (ud *UdevDevice) UdevDeviceGetDevtype() string {
	return C.GoString(C.udev_device_get_devtype(ud.device))
}

// UdevDeviceGetDevnode returns the device node file name belonging to the udev device.
func (ud *UdevDevice) UdevDeviceGetDevnode() string {
	return C.GoString(C.udev_device_get_devnode(ud.device))
}

// UdevDeviceGetAction returns device action when it is monitored.
func (ud *UdevDevice) UdevDeviceGetAction() string {
	return C.GoString(C.udev_device_get_action(ud.device))
}
