package settings_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/connection"
	"github.com/hnnsgstfssn/nm/settings"
)

const profileID = "My Connection"

// Add a saved ethernet profile with a static address, unless a profile with
// the same id is already there. NetworkManager fills in the UUID.
func Example_addStaticProfile() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bus, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	exists, err := hasProfile(ctx, bus, profileID)
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		fmt.Println("already there")
		return
	}

	profile, err := settings.AddConnection(ctx, bus, nm.ConnectionSettings{
		"connection": {
			"id":             profileID,
			"type":           "802-3-ethernet",
			"interface-name": "eth1",
			"autoconnect":    true,
		},
		"802-3-ethernet": {
			"auto-negotiate": false,
		},
		"ipv4": {
			"method":  "manual",
			"gateway": "192.168.1.1",
			"address-data": []map[string]any{
				{"address": "192.168.1.1", "prefix": uint32(24)},
			},
			"never-default": true,
		},
		// "disabled" would be closer to the intent but NetworkManager does
		// not accept it here.
		"ipv6": {"method": "ignore"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(profile.Path)
}

func hasProfile(ctx context.Context, bus nm.Conn, id string) (bool, error) {
	profiles, err := settings.ListConnections(ctx, bus)
	if err != nil {
		return false, err
	}
	for _, p := range profiles {
		s, err := connection.GetSettings(ctx, bus, p)
		if err != nil {
			// Raced a deletion, or a profile this caller may not read.
			continue
		}
		if s["connection"]["id"] == id {
			return true, nil
		}
	}
	return false, nil
}
