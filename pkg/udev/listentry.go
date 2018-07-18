// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"

// UdevListEntry wraps udev_list_entry struct
type UdevListEntry struct {
	listEntry *C.struct_udev_list_entry
}

// newUdevListEntry a private helper function returns a pointer UdevListEntry
func newUdevListEntry(ptr *C.struct_udev_list_entry) *UdevListEntry {
	if ptr == nil {
		return nil
	}
	le := &UdevListEntry{
		listEntry: ptr,
	}
	return le
}

// GetNext ...
func (le *UdevListEntry) GetNext() *UdevListEntry {
	return newUdevListEntry(C.udev_list_entry_get_next(le.listEntry))
}

// GetName ...
func (le *UdevListEntry) GetName() string {
	return C.GoString(C.udev_list_entry_get_name(le.listEntry))
}
