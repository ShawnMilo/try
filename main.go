package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	var success bool

	// must be piped-in data; not an interactive session
	if isTTY() {
		log.Fatal("please pipe in some data")
	}

	if len(os.Args) < 2 {
		log.Fatal("Please pass in a program name.")
		return
	}
	app := os.Args[1]

	raw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read stdin: %s\n", err)
	}
	input := string(raw)

	// if anything fails, write the original input
	defer func() {
		if success {
			return
		}
		fmt.Print(input)
	}()

	cmd := exec.Command(app)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	err = cmd.Start()
	if err != nil {
		return
	}
	go func() {
		_, err = fmt.Fprintln(stdin, input)
		if err != nil {
			log.Fatalf("Failed to write to cmd stdin: %s\n", err)
		}
		stdin.Close()
	}()

	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		return
	}
	fmt.Print(string(output))
	success = true
}

func isTTY() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return stat.Mode()&os.ModeCharDevice != 0
}
