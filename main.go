package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func askChromePath() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter chrome directory: ")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func createFile(chromePathValue string) error {
	return ioutil.WriteFile("./chrome-path.txt", []byte(chromePathValue), 0644)
}

func openKeep() error {
	chromePath, err := ioutil.ReadFile("./chrome-path.txt")
	if err != nil {
		return err
	}

	arg0 := "-app=https://keep.google.com"

	cmd := exec.Command(string(chromePath), arg0)
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	log.Println(string(stdout))

	return nil
}

func main() {
	if !fileExists("./chrome-path.txt") {
		path := askChromePath()

		if err := createFile(path); err != nil {
			panic("unhandled error")
		}
	}

	if err := openKeep(); err != nil {
		panic("unhandled error")
	}
}
