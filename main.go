package main

import (
	"fmt"
	"log"
	"main/util"
	"os"
	"os/exec"
	"strings"
	"time"
)

var args []string = os.Args

func HandleError(err error) {
	fmt.Println("Error: ")
	log.Fatal(err)
}

func GetExtension(name string) string {
	var SplitByDot []string = strings.Split(name, ".")
	return SplitByDot[len(SplitByDot)-1]
}

func Benchmark(c chan util.BenchmarkMessage, command string, fileArg string) {
	c <- *util.NewMessage("INFORMATION", false, true)
	c <- *util.NewMessage("Preparing Command", false, false)
	_cmd := exec.Command(command, fileArg)
	_cmd.Stdin = os.Stdin
	_cmd.Stdout = os.Stdout
	_cmd.Stderr = os.Stderr
	c <- *util.NewMessage("Starting", false, false)
	c <- *util.NewMessage("OUTPUT", false, true)
	_cmd.Start()
	startedAt := time.Now()
	_cmd.Wait()
	time_elapsed := time.Since(startedAt)
	c <- *util.NewMessage("RESULTS", false, true)
	c <- *util.NewMessage("Took "+time_elapsed.String()+" to Execute", false, false)
	c <- *util.NewMessage("Closing Program", true, false)
}

func main() {
	if len(args) > 0 {
		var commandName string
		switch strings.ToLower(GetExtension(args[1])) {
		case "js":
			commandName += "node"
		case "py":
			commandName += "py"
		}
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			HandleError(err)
			return
		}

		MessageChannel := make(chan util.BenchmarkMessage)

		go Benchmark(MessageChannel, commandName, args[1])

		for {
			recieved := <-MessageChannel
			if recieved.Closed {
				fmt.Println("[Benchmark]: " + "Done")
				break
			}
			if recieved.Title {
				fmt.Println("====== " + recieved.Message + " ======")
			} else {
				fmt.Println("[Benchmark]: " + recieved.Message)
			}
		}
	} else {
		HandleError(fmt.Errorf("e: No arguments Provided. (path to file)"))
	}
}
