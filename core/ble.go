package core

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var lampChar bluetooth.DeviceCharacteristic

func ConnectToLamp() error {
	if err := adapter.Enable(); err != nil {
		return err
	}

	var foundAddr bluetooth.Address
	found := false

	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if result.Address.String() == "17:24:1E:DE:85:B7" {
			foundAddr = result.Address
			found = true
			err := adapter.StopScan()
			if err != nil {
				return
			}
		}
	})
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("device not found")
	}

	dev, err := adapter.Connect(foundAddr, bluetooth.ConnectionParams{})
	if err != nil {
		return err
	}

	srvcs, err := dev.DiscoverServices(nil)
	if err != nil {
		return err
	}

	for _, srvc := range srvcs {
		if srvc.UUID().String() == "0000a032-0000-1000-8000-00805f9b34fb" {
			chars, err := srvc.DiscoverCharacteristics(nil)
			if err != nil {
				return err
			}

			for _, char := range chars {
				if char.UUID().String() == "0000a040-0000-1000-8000-00805f9b34fb" {
					lampChar = char
					return nil
				}
			}
		}
	}
	return nil
}

func WriteToLamp(data []byte) error {
	if lampChar == (bluetooth.DeviceCharacteristic{}) {
		return fmt.Errorf("lamp characteristic not initialized")
	}

	_, err := lampChar.WriteWithoutResponse(data)
	return err
}
