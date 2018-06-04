// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// Udev wraps a libudev udev object
type Udev struct {
	udev *C.struct_udev
}

// newDevis a private helper function and returns a pointer to a new udev
func newUdev(ptr *C.struct_udev) (u *Udev) {
	if ptr == nil {
		return nil
	}
	u = &Udev{
		udev: ptr,
	}
	return
}

// NewDev is a function and which returns a pointer to a new Udev
func NewUdev() (u *Udev) {
	udev := C.udev_new()
	if udev == nil {
		return nil
	}
	return newUdev(udev)
}

// UnrefUdev frees udev structure.
func (u *Udev) UnrefUdev() {
	C.udev_unref(u.udev)
}

// NewDeviceFromSysPath identify the block device currently attached to the system
func (u *Udev) NewDeviceFromSysPath(sysPath string) *UdevDevice {
	syspath := C.CString(sysPath)
	if syspath == nil {
		return nil
	}
	defer freeCharPtr(syspath)
	dev := newUdevDevice(C.udev_device_new_from_syspath(u.udev, syspath))
	return dev
}

// NewUdevEnumerate returns a pointer to a Udevenumerate
func NewUdevEnumerate(u *Udev) (ue *UdevEnumerate) {
	ue = &UdevEnumerate{
		enumerate: C.udev_enumerate_new(u.udev),
	}
	return
}

// NewUdeviceMonitor use newUdeviceMon() and use returns UdeviceMon pointer in success
func (u *Udev) NewUdeviceMonitor(source string) *UdevMonitor {
	monitorSources := C.CString(source)
	if monitorSources == nil {
		return nil
	}
	defer freeCharPtr(monitorSources)
	udevmon := newUdevMonitor(C.udev_monitor_new_from_netlink(u.udev, monitorSources))
	return udevmon
}
