// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdevEnumerate wraps a libudev udev_enumerate object
type UdevEnumerate struct {
	enumerate *C.struct_udev_enumerate
}

// UnrefUdevEnumerate frees udev_enumerate structure.
func (ue *UdevEnumerate) UnrefUdevEnumerate() {
	C.udev_enumerate_unref(ue.enumerate)
}

// UdevEnumerateAddMatchSubsystem adds filter in UdeviceMon struct.
func (ue *UdevEnumerate) UdevEnumerateAddMatchSubsystem(subSystem string) int {
	subsystem := C.CString(subSystem)
	if subsystem == nil {
		return -1
	}
	defer freeCharPtr(subsystem)
	ret := C.udev_enumerate_add_match_subsystem(ue.enumerate, subsystem)
	return int(ret)
}

// UdevEnumerateScanDevices ...
func (ue *UdevEnumerate) UdevEnumerateScanDevices() int {
	ret := C.udev_enumerate_scan_devices(ue.enumerate)
	return int(ret)
}

//UdevEnumerateGetListEntry ...
func (ue *UdevEnumerate) UdevEnumerateGetListEntry() *UdevListEntry {
	return newUdevListEntry(C.udev_enumerate_get_list_entry(ue.enumerate))
}
