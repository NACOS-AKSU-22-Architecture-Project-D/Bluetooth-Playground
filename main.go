package main

import (
	"bluetooth_playground/utils"
	"fmt"
	"log"
	"os"

	"tinygo.org/x/bluetooth"
)

func main() {
	envErr := os.Setenv("DEVICE_MAC", "E8:F7:91:EB:B1:31")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}

	utils.ConnectSpecific()
	return


	//UNREACHABLE CODE
	adapter := bluetooth.DefaultAdapter
	// Enable adapter
	err := adapter.Enable()
	if err != nil {
		panic("failed to enable BLE adapter")
	}

	// Start scanning and define callback for scan results
	fmt.Println("Scanning for nearby devices...")
	err = adapter.Scan(onScan)
	if err != nil {
		panic("failed to register scan callback")
	}
}

func onScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	if device.LocalName() != "" {
		log.Println("found device:", device.Address.String(), device.RSSI, device.LocalName())
	}
	// log.Println("found device:", device.Address.String(), device.RSSI, device.LocalName(), device.AdvertisementPayload)
}
