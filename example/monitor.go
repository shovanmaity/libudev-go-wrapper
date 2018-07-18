package main

import (
	"fmt"
	"syscall"

	"github.com/shovanmaity/libudev-go-wrapper/pkg/fdutil"
	"github.com/shovanmaity/libudev-go-wrapper/pkg/udev"
)

func main() {
	newUdev := udev.NewUdev()
	if newUdev == nil {
		fmt.Println("Unable to create udev new object")
	}
	defer newUdev.UnrefUdev()
	udevMonitor := newUdev.NewMonitor(udev.UDEV_SOURCE)
	if udevMonitor == nil {
		fmt.Println("Unable to create udevmonitor object")
	}
	defer udevMonitor.UdevMonitorUnref()
	ret := udevMonitor.AddMatchSubsystemDevtypeFilter(udev.UDEV_SUBSYSTEM)
	if ret < 0 {
		fmt.Println("Unable to apply monitor filter")
	}
	ret = udevMonitor.EnableReceiving()
	if ret < 0 {
		fmt.Println("Unable to enable monitor receiving")
	}
	fd := udevMonitor.GetFdValue()
	if fd < 0 {
		fmt.Println("Unable to get fd value")
	}
	for {
		fds := &syscall.FdSet{}
		fdutil.FD_ZERO(fds)
		fdutil.FD_SET(fds, int(fd))
		ret, _ := syscall.Select(int(fd)+1, fds, nil, nil, nil)
		if ret <= 0 {
			continue
		}
		if fdutil.FD_ISSET(fds, int(fd)) {
			newdev := udevMonitor.ReceiveDevice()
			if newdev.GetDevnode() == "" {
				continue
			}
			devType := newdev.GetPropertyValue(udev.UDEV_DEVTYPE)
			action := newdev.GetAction()
			if devType == "disk" && (action == udev.UDEV_ACTION_ADD || action == udev.UDEV_ACTION_REMOVE) {
				fmt.Println("-------------- Disk Details ------------------")
				fmt.Println("Action : ", action)
				fmt.Println("Vendor : ", newdev.GetPropertyValue(udev.UDEV_VENDOR))
				fmt.Println("Model : ", newdev.GetPropertyValue(udev.UDEV_MODEL))
				fmt.Println("Serial : ", newdev.GetPropertyValue(udev.UDEV_SERIAL))
				fmt.Println("----------------------------------------------")
			}
			newdev.UnrefDeviceUdev()
		}
	}
}
