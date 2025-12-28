package main

import (
	"fmt"
	"openlamp-cli/core"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: openlamp-cli <command> [value]")
		return
	}

	var err error

	command := os.Args[1]
	if command == "turnon" || command == "turnoff" {
		if len(os.Args) != 2 {
			fmt.Println("Incorrect number of arguments")
			return
		}
	}
	if command == "brightness" {
		if len(os.Args) != 3 {
			fmt.Println("Please specify brightness value (0-255)")
			return
		}
	}

	err = core.ConnectToLamp()
	if err != nil {
		fmt.Println("connect error:", err)
		os.Exit(1)
	}

	switch command {
	case "turnon":
		err = core.TurnOn()
	case "turnoff":
		err = core.TurnOff()
	case "brightness":
		value, err := strconv.Atoi(os.Args[2])
		if err != nil || value < 0 || value > 255 {
			fmt.Println("Please specify brightness value (0-255)")
			return
		}
		err = core.SetBrightness(byte(value))
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
