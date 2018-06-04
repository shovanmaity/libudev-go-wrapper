package main

import (
	"fmt"

	"github.com/ShovanMaity/libuev-go-wrapper/pkg/udev"
)

func main() {
	newUdev := udev.NewUdev()
	if newUdev == nil {
		fmt.Println("Unable to create Udev object")
	}
	defer newUdev.UnrefUdev()
	newUdevEnumerate := udev.NewUdevEnumerate(newUdev)
	if newUdevEnumerate == nil {
		fmt.Println("Unable to create UdevEnumerate object")
	}
	defer newUdevEnumerate.UnrefUdevEnumerate()
	ret := newUdevEnumerate.UdevEnumerateAddMatchSubsystem(udev.UDEV_SUBSYSTEM)
	if ret < 0 {
		fmt.Println("Unable to apply subsystem filter")
	}
	err := newUdevEnumerate.UdevEnumerateScanDevices()
	if err < 0 {
		fmt.Println("Unable to scan device list")
	}
	for l := newUdevEnumerate.UdevEnumerateGetListEntry(); l != nil; l = l.UdevListEntryGetNext() {
		s := l.UdevListEntryGetName()
		dev := newUdev.NewDeviceFromSysPath(s)
		if dev == nil {
			continue
		}
		if dev.UdevDeviceGetDevtype() == "disk" && dev.PropertyValue(udev.UDEV_TYPE) == "disk" {
			fmt.Println("-------------- Disk Details ------------------")
			fmt.Println("Vendor : ", dev.PropertyValue(udev.UDEV_VENDOR))
			fmt.Println("Model : ", dev.PropertyValue(udev.UDEV_MODEL))
			fmt.Println("Serial : ", dev.PropertyValue(udev.UDEV_SERIAL))
			fmt.Println("----------------------------------------------")
		}
		dev.UnrefDeviceUdev()
	}
}
