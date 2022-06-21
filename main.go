package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func getUsername() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil

}

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	machinename := strings.Split(hostname, ".")
	return machinename[0], nil
}

func getCurrentDir() (string, error) {
	fullPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := strings.Split(fullPath, "/")

	return dir[len(dir)-1], nil
}

func main() {
	user, err := getUsername()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	hostname, err := getHostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	dir, err := getCurrentDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s@%s %s %% ", user, hostname, dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
