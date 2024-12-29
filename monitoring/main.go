package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	sites := readWebsiteFiles()

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
		fmt.Println(err)
		fmt.Println("Error when making request for ", site)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, " online")
	} else {
		fmt.Println(site, " down -", resp.StatusCode)
	}
}

func readWebsiteFiles() []string {
	var sites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	file.Close()
	fmt.Println(sites)

	return sites
}

func viewLog() {
	fmt.Println("Viewing log")
}
