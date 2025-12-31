package core

import (
	"fmt"
	"strconv"
)

var CmdTurnOn = []byte{0x55, 0xAA, 0x01, 0x08, 0x05, 0x01, 0xF1}
var CmdTurnOff = []byte{0x55, 0xAA, 0x01, 0x08, 0x05, 0x00, 0xF2}

var CmdBrightness = []byte{0x55, 0xAA, 0x01, 0x08, 0x01} // does not contain value and checksum
var CmdColor = []byte{0x55, 0xAA, 0x03, 0x08, 0x02}      // does not contain value and checksum

var Temperatures = map[string][]byte{
	"white":    {0x55, 0xAA, 0x01, 0x08, 0x09, 0x01, 0xED},
	"natural":  {0x55, 0xAA, 0x01, 0x08, 0x09, 0x02, 0xEC},
	"sunlight": {0x55, 0xAA, 0x01, 0x08, 0x09, 0x03, 0xEB},
	"sunset":   {0x55, 0xAA, 0x01, 0x08, 0x09, 0x04, 0xEA},
	"candle":   {0x55, 0xAA, 0x01, 0x08, 0x09, 0x05, 0xE9},
}

func TurnOn() error {
	return WriteToLamp(CmdTurnOn)
}

func TurnOff() error {
	return WriteToLamp(CmdTurnOff)
}

func SetBrightness(value byte) error {
	if value > 255 {
		value = 255
	}

	cmd := append([]byte{}, CmdBrightness...)
	cmd = append(cmd, value)
	cmd = append(cmd, ComputeChecksum(cmd))

	return WriteToLamp(cmd)
}

func SetColor(value string) error {
	r, _ := strconv.ParseUint(value[0:2], 16, 8) // input validated in main.go
	g, _ := strconv.ParseUint(value[2:4], 16, 8) // input validated in main.go
	b, _ := strconv.ParseUint(value[4:6], 16, 8) // input validated in main.go

	cmd := append([]byte{}, CmdColor...)
	cmd = append(cmd, byte(r), byte(g), byte(b))
	cmd = append(cmd, ComputeChecksum(cmd))

	return WriteToLamp(cmd)
}

func SetTemperature(value string) error {
	cmd, ok := Temperatures[value]
	if !ok {
		return fmt.Errorf("available temperatures: white, natural, sunlight, sunset, candle")
	}

	return WriteToLamp(cmd)
}

func ComputeChecksum(cmd []byte) byte {
	sum := 0
	for _, b := range cmd {
		sum += int(b)
	}
	return byte(0xFF - (sum & 0xFF))
}
