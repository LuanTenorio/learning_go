package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	startMonitoringCode = 0
	viewLogCode         = 1
	exitProgramCode     = 2
	monitoraments       = 3
	delay               = 2
	logFileName         = "log.txt"
	websitesFileName    = "websites.txt"
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

	recordsLog(site, resp.StatusCode == 200)
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

	return sites
}

func recordsLog(site string, status bool) {
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	date := time.Now().Format("02/01/2006 15:04:05")

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(date + " " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func viewLog() {
	file, err := os.ReadFile(logFileName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
