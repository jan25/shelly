package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("echo", "hello").Output()
	fmt.Printf("%s", out)
}
