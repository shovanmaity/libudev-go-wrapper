// +build linux,cgo

package udev // import "github.com/shovanmaity/libudev-go-wrapper/pkg/udev"
/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// Udev wraps udev c struct
type Udev struct {
	udev *C.struct_udev
}

// newUdev is a private helper function and returns a pointer of Udev
func newUdev(ptr *C.struct_udev) (u *Udev) {
	if ptr == nil {
		return nil
	}
	u = &Udev{
		udev: ptr,
	}
	return
}

// NewDev returns a pointer to Udev
func NewUdev() (u *Udev) {
	udev := C.udev_new()
	if udev == nil {
		return nil
	}
	return newUdev(udev)
}

// UnrefUdev frees udev pointer
func (u *Udev) UnrefUdev() {
	C.udev_unref(u.udev)
}

// GetDeviceFromSysPath identify the block device currently attached to the system
func (u *Udev) GetDeviceFromSysPath(sysPath string) *UdevDevice {
	syspath := C.CString(sysPath)
	if syspath == nil {
		return nil
	}
	defer freeCharPtr(syspath)
	device := newUdevDevice(C.udev_device_new_from_syspath(u.udev, syspath))
	return device
}

// NewUdevEnumerate returns a pointer to a Udevenumerate
func NewUdevEnumerate(u *Udev) *UdevEnumerate {
	ue := &UdevEnumerate{
		enumerate: C.udev_enumerate_new(u.udev),
	}
	return ue
}

// NewMonitor use newUdeviceMon() and use returns UdeviceMon pointer in success
// Monitor source can be kernel or udev.
func (u *Udev) NewMonitor(source string) *UdevMonitor {
	monitorSources := C.CString(source)
	if monitorSources == nil {
		return nil
	}
	defer freeCharPtr(monitorSources)
	udevmon := newUdevMonitor(C.udev_monitor_new_from_netlink(u.udev, monitorSources))
	return udevmon
}
