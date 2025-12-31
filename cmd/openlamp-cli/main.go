package main

import (
	"fmt"
	"openlamp-cli/core"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	var err error

	command := os.Args[1]
	if command == "help" {
		printHelp()
		return
	}
	if command == "turnon" || command == "turnoff" {
		if len(os.Args) != 2 {
			fmt.Println("Incorrect number of arguments")
			return
		}
	}
	if command == "brightness" && len(os.Args) != 3 {
		fmt.Println("correct usage: openlamp brightness <value> | (0-255)")
		return
	}
	if command == "color" && len(os.Args) != 3 {
		fmt.Println("correct usage: openlamp color <hex>")
		return
	}
	if command == "temp" && len(os.Args) != 3 {
		fmt.Println("correct usage: openlamp temp <temperature>")
		printAvailable("Available temperatures: ", core.Temperatures)
		return
	}
	if command == "scene" && len(os.Args) != 3 {
		fmt.Println("correct usage: openlamp scene <scene>")
		printAvailable("Available scenes: ", core.Scenes)
		return
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
	case "color":
		hexColor := os.Args[2]

		if len(hexColor) != 6 {
			fmt.Println("Please specify a 6-character hex color (without #)")
			return
		}

		_, err := strconv.ParseUint(hexColor, 16, 32)
		if err != nil {
			fmt.Println("Invalid hex color. Use only characters 0-9 and a-f")
			return
		}

		err = core.SetColor(hexColor)
	case "temp":
		err = core.SetTemperature(os.Args[2])
		if err != nil {
			fmt.Println("Temperature not found: ", os.Args[2])
			printAvailable("Available temperatures: ", core.Temperatures)
		}
	case "scene":
		err = core.SetScene(os.Args[2])
		if err != nil {
			fmt.Println("Scene not found: ", os.Args[2])
			printAvailable("Available scenes: ", core.Scenes)
		}
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}
}

func printAvailable(title string, items map[string][]byte) {
	fmt.Println(title)
	for name := range items {
		fmt.Println(" -", name)
	}
}

func printHelp() {
	fmt.Println("Openlamp CLI")
	fmt.Println("Usage: openlamp <command> [value]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Printf("%-30s %s\n", " turnon", "Turn the lamp on")
	fmt.Printf("%-30s %s\n", " turnoff", "Turn the lamp off")
	fmt.Printf("%-30s %s\n", " brightness <0-255>", "Set brightness level")
	fmt.Printf("%-30s %s\n", " color <RRGGBB>", "Set color in hex (without #)")
	fmt.Printf("%-30s %s\n", " temp <temperature>", "Set light temperature (list at: openlamp temp)")
	fmt.Printf("%-30s %s\n", " scene <scene>", "Set scene (list at: openlamp scene)")
	fmt.Printf("%-30s %s\n", " help", "Show this menu")
}
