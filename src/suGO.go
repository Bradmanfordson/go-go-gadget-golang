package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("./suGO <arg>")
}

func execute_command(cmd string) string {
	fmt.Println(cmd)
	output, err := exec.Command("sudo", cmd).Output()

	if err != nil {
    return "error" // getting an error when the first ARG contains a flag... oh well :)
	} else {
		// fmt.Println(string(output[:]))
		return string(output)
	}
}

func main() {
	args := os.Args

	if len(args) > 1 {
		fmt.Println("All args: ", args)
		fmt.Println("Program name: ", args[0])
		fmt.Println("Command to run as Super user: ", args[1:])
		fmt.Println("Command output: ", execute_command(strings.Join(args[1:], " ")))
	} else {
		usage()
	}

}
