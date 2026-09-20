package manager_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/device"
	"github.com/hnnsgstfssn/nm/manager"
)

// List every device NetworkManager knows about, realized or not, with its
// interface name and state.
func Example_listDevices() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bus, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	devices, err := manager.GetPropertyAllDevices(ctx, bus)
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devices {
		iface, err := device.GetPropertyInterface(ctx, bus, dev)
		if err != nil {
			log.Fatal(err)
		}
		state, err := device.GetPropertyState(ctx, bus, dev)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s %s\n", dev.Path, iface, state)
	}
}
