package core

var CmdTurnOn = []byte{0x55, 0xAA, 0x01, 0x08, 0x05, 0x01, 0xF1}
var CmdTurnOff = []byte{0x55, 0xAA, 0x01, 0x08, 0x05, 0x00, 0xF2}
var CmdBrightness = []byte{0x55, 0xAA, 0x01, 0x08, 0x01} // This byte array does not contain brightness value and checksum

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

	cmd := make([]byte, len(CmdBrightness))
	copy(cmd, CmdBrightness)

	cmd = append(cmd, value)
	cmd = append(cmd, ComputeChecksum(cmd))

	return WriteToLamp(cmd)
}

func ComputeChecksum(cmd []byte) byte {
	sum := 0
	for _, b := range cmd {
		sum += int(b)
	}
	return byte(0xFF - (sum & 0xFF))
}
