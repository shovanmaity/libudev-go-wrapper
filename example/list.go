package main

import (
	"fmt"

	"github.com/shovanmaity/libudev-go-wrapper/pkg/udev"
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
	ret := newUdevEnumerate.AddMatchSubsystemFilter(udev.UDEV_SUBSYSTEM)
	if ret < 0 {
		fmt.Println("Unable to apply subsystem filter")
	}
	err := newUdevEnumerate.ScanDevices()
	if err < 0 {
		fmt.Println("Unable to scan device list")
	}
	for l := newUdevEnumerate.GetListEntry(); l != nil; l = l.GetNext() {
		s := l.GetName()
		dev := newUdev.GetDeviceFromSysPath(s)
		if dev == nil {
			continue
		}
		if dev.GetDevtype() == "disk" && dev.GetPropertyValue(udev.UDEV_TYPE) == "disk" {
			fmt.Println("-------------- Disk Details ------------------")
			fmt.Println("Vendor : ", dev.GetPropertyValue(udev.UDEV_VENDOR))
			fmt.Println("Model : ", dev.GetPropertyValue(udev.UDEV_MODEL))
			fmt.Println("Serial : ", dev.GetPropertyValue(udev.UDEV_SERIAL))
			fmt.Println("----------------------------------------------")
		}
		dev.UnrefDeviceUdev()
	}
}
