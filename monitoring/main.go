package main

import "fmt"
import "os"

const (
	StartMonitoring = 0
	ViewLog = 1
	ExitProgram = 2
)

func main(){
	for {
		showPanel()
		command := readCommand()
		executeCommand(command)
	}
}

func showPanel(){
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

func executeCommand(command int){
	switch command {
	case StartMonitoring: 
		startMonitoring()
	case ViewLog:
		viewLog()
	case ExitProgram:
		fmt.Println("Leaving the program")
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
	}
}

func startMonitoring(){
	fmt.Println("Starting monitoring")
	
}

func viewLog(){
	fmt.Println("Viewing log")
}
