package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage %v <Problem No.>\n", os.Args[0])
		return
	}

	url := fmt.Sprintf("https://projecteuler.net/problem=%v", os.Args[1])
	cmd := exec.Command("xdg-open", url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
