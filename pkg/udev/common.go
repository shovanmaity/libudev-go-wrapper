// +build linux,cgo
package udev

/*
  #cgo LDFLAGS: -ludev
  #include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

const (
	UDEV_SUBSYSTEM     = "block"           // udev to filter this device type
	UDEV_SERIAL        = "ID_SERIAL_SHORT" // udev attribute to get device serial number
	UDEV_MODEL         = "ID_MODEL"        // udev attribute to get device model number
	UDEV_VENDOR        = "ID_VENDOR"       // udev attribute to get device vendor details
	UDEV_TYPE          = "ID_TYPE"         // udev attribute to get device type
	UDEV_ACTION        = "UDEV_ACTION"     // udev attribute to get monitor device action
	UDEV_ACTION_ADD    = "add"             // udev attribute constant for add action
	UDEV_ACTION_REMOVE = "remove"          // udev attribute constant for remove action
	UDEV_SOURCE        = "udev"            // udev source constant
	UDEV_DEVTYPE       = "DEVTYPE"         // udev attribute to get device device type ie - disk or part
)

//
func freeCharPtr(s *C.char) {
	C.free(unsafe.Pointer(s))
}
