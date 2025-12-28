package main

import (
	"fmt"
	"openlamp-cli/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: openlamp-cli <command>")
		return
	}

	var err error

	err = core.ConnectToLamp()
	if err != nil {
		fmt.Println("connect error:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "turnon":
		err = core.TurnOn()
	case "turnoff":
		err = core.TurnOff()
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
