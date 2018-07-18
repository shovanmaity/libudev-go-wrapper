// +build linux,cgo

package udev // import "github.com/shovanmaity/libudev-go-wrapper/pkg/udev"

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdevEnumerate wraps udev_enumerate struct
type UdevEnumerate struct {
	enumerate *C.struct_udev_enumerate
}

// UnrefUdevEnumerate frees udev_enumerate struct
func (ue *UdevEnumerate) UnrefUdevEnumerate() {
	C.udev_enumerate_unref(ue.enumerate)
}

// AddMatchSubsystemFilter adds filter in UdeviceMon struct.
func (ue *UdevEnumerate) AddMatchSubsystemFilter(subSystem string) int {
	subsystem := C.CString(subSystem)
	defer freeCharPtr(subsystem)
	ret := C.udev_enumerate_add_match_subsystem(ue.enumerate, subsystem)
	return int(ret)
}

// ScanDevices ...
func (ue *UdevEnumerate) ScanDevices() int {
	ret := C.udev_enumerate_scan_devices(ue.enumerate)
	return int(ret)
}

//GetListEntry ...
func (ue *UdevEnumerate) GetListEntry() *UdevListEntry {
	return newUdevListEntry(C.udev_enumerate_get_list_entry(ue.enumerate))
}
