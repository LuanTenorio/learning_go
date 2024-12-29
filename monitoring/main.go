package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	startMonitoringCode = 0
	viewLogCode         = 1
	exitProgramCode     = 2
	monitoraments       = 3
	delay               = 2
)

func main() {
	for {
		showPanel()
		command := readCommand()
		executeCommand(command)
	}
}

func showPanel() {
	fmt.Println("0 - Start monitoring")
	fmt.Println("1 - View log")
	fmt.Println("2 - Exit the program")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("commando: ", command)

	return command
}

func executeCommand(command int) {
	switch command {
	case startMonitoringCode:
		startMonitoring()
	case viewLogCode:
		viewLog()
	case exitProgramCode:
		fmt.Println("Leaving the program")
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
	}
}

func startMonitoring() {
	fmt.Println("Starting monitoring")
	sites := []string{"https://github.com/LuanTenorio", "https://www.youtube.com/"}

	for i := 0; i < monitoraments; i++ {
		for _, site := range sites {
			checkSite(site)
		}

		time.Sleep(delay * time.Second)
	}
}

func checkSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error when making request for ", site)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, " online")
	} else {
		fmt.Println(site, " down -", resp.StatusCode)
	}
}

func viewLog() {
	fmt.Println("Viewing log")
}
