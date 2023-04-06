package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	commandoutput, err := exec.Command("ssh", "-p 2020", "root@sr1.packops.dev", "ls", "/root").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(commandoutput))
}
