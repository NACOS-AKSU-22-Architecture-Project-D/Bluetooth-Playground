package utils

import (
	"fmt"
	"log"
	"os"

	"tinygo.org/x/bluetooth"
)

func ConnectSpecific() {
	log.Println("This Module Connects to a specific device")
	adapter := bluetooth.DefaultAdapter
	// Enable adapter
	err := adapter.Enable()
	if err != nil {
		panic("failed to enable BLE adapter")
	}

	// Start scanning and define callback for scan results
	fmt.Println("Scanning for nearby devices...")
	err = adapter.Scan(ConnectDevice)
	if err != nil {
		panic("failed to register scan callback")
	}
}

func ConnectDevice(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	//log.Println("found device:", device.Address.String(), device.RSSI, device.LocalName(), device.AdvertisementPayload)
	if device.Address.String() == os.Getenv("DEVICE_MAC") {
		log.Println("found target device:", device.Address.String(), device.RSSI, device.LocalName())
		// Start connecting in a goroutine to not block
		go func() {
			res, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{})
			if err != nil {
				println("error connecting:", err.Error())
				return
			}
			// Call connect callback
			onConnect(device, res)

		}()
	}
}

func onConnect(scanResult bluetooth.ScanResult, device any) {
	println("connected:", scanResult.LocalName(), scanResult.Address.String())
}
