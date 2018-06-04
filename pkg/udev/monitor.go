// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdeviceMon wraps a libudev monitor device object
type UdevMonitor struct {
	monitor *C.struct_udev_monitor
}

// newUdeviceMon is a helper function and returns a pointer to a new monitor device.
func newUdevMonitor(ptr *C.struct_udev_monitor) (um *UdevMonitor) {
	if ptr == nil {
		return nil
	}
	um = &UdevMonitor{
		monitor: ptr,
	}
	return
}

// UdevMonitorRef ....
func (um *UdevMonitor) UdevMonitorRef() {
	C.udev_monitor_ref(um.monitor)
}

// UdevMonitorUnref ....
func (um *UdevMonitor) UdevMonitorUnref() {
	C.udev_monitor_unref(um.monitor)
}

// UdevMonitorFilterAddMatchSubsystemDevtype adds filter in UdeviceMon struct.
func (um *UdevMonitor) UdevMonitorFilterAddMatchSubsystemDevtype(key string) int {
	subsystem := C.CString(key)
	if subsystem == nil {
		return -1
	}
	defer freeCharPtr(subsystem)
	ret := C.udev_monitor_filter_add_match_subsystem_devtype(um.monitor, subsystem, nil)
	return int(ret)
}

// UdevMonitorEnableReceiving binds udev_monitor socket to event source.
func (um *UdevMonitor) UdevMonitorEnableReceiving() int {
	return int(C.udev_monitor_enable_receiving(um.monitor))
}

// UdevMonitorGetFd retrieves socket file descriptor associated with monitor.
func (um *UdevMonitor) UdevMonitorGetFd() int {
	return int(C.udev_monitor_get_fd(um.monitor))
}

// UdevMonitorReceiveDevice receives data from udev monitor socket.
func (um *UdevMonitor) UdevMonitorReceiveDevice() *UdevDevice {
	return newUdevDevice(C.udev_monitor_receive_device(um.monitor))
}

// UdevMonitorSetReceiveBufferSize ...
func (um *UdevMonitor) UdevMonitorSetReceiveBufferSize() {

}

// UdevMonitorNewFromNetlink ..
func (um *UdevMonitor) UdevMonitorNewFromNetlink() {

}

//udev_monitor_filter_add_match_tag ()
//udev_monitor_filter_update ()
//udev_monitor_filter_remove ()
