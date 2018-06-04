// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdevDevice wraps a libudev device object
type UdevListEntry struct {
	listEntry *C.struct_udev_list_entry
}

// newUdevListEntry a private helper function and returns a pointer to a new udev
func newUdevListEntry(ptr *C.struct_udev_list_entry) (le *UdevListEntry) {
	if ptr == nil {
		return nil
	}
	le = &UdevListEntry{
		listEntry: ptr,
	}
	return
}

//UdevListEntryGetNext ...
func (le *UdevListEntry) UdevListEntryGetNext() *UdevListEntry {
	return newUdevListEntry(C.udev_list_entry_get_next(le.listEntry))
}

//UdevListEntryGetName ...
func (le *UdevListEntry) UdevListEntryGetName() string {
	return C.GoString(C.udev_list_entry_get_name(le.listEntry))
}
