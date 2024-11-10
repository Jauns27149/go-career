package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	command := exec.Command("echo", "Hello World")
	output, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	cmd := exec.Command("bash", "base/shell/script.sh")
	bytes, err := cmd.Output()
	fmt.Println(string(bytes))
}
