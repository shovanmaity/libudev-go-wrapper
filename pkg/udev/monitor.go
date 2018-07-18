// +build linux,cgo
package udev // import "github.com/shovanmaity/libudev-go-wrapper/pkg/udev"

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdeviceMonitor wraps udev_monitor c struct
type UdevMonitor struct {
	monitor *C.struct_udev_monitor
}

// newUdeviceMon is private helper function and returns
// UdevMonitor pointer on success
func newUdevMonitor(ptr *C.struct_udev_monitor) (um *UdevMonitor) {
	if ptr == nil {
		return nil
	}
	um = &UdevMonitor{
		monitor: ptr,
	}
	return
}

// UdevMonitorRef
func (um *UdevMonitor) UdevMonitorRef() {
	C.udev_monitor_ref(um.monitor)
}

// UdevMonitorUnref ....
func (um *UdevMonitor) UdevMonitorUnref() {
	C.udev_monitor_unref(um.monitor)
}

// AddMatchSubsystemDevtypeFilter adds filter in UdeviceMon it starts monitoring
// only for filtered subsystem like block / usb ...
func (um *UdevMonitor) AddMatchSubsystemDevtypeFilter(key string) int {
	subsystem := C.CString(key)
	defer freeCharPtr(subsystem)
	ret := C.udev_monitor_filter_add_match_subsystem_devtype(um.monitor, subsystem, nil)
	return int(ret)
}

// EnableReceiving binds udev_monitor socket to event source.
func (um *UdevMonitor) EnableReceiving() int {
	return int(C.udev_monitor_enable_receiving(um.monitor))
}

// GetFdValue retrieves socket file descriptor associated with monitor.
func (um *UdevMonitor) GetFdValue() int {
	return int(C.udev_monitor_get_fd(um.monitor))
}

// ReceiveDevice receives data from udev monitor socket.
func (um *UdevMonitor) ReceiveDevice() *UdevDevice {
	return newUdevDevice(C.udev_monitor_receive_device(um.monitor))
}
