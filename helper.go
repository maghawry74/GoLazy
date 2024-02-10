package main

import (
	"fmt"
	"os/exec"
)

func GetConfiguration() []*exec.Cmd {

	return []*exec.Cmd{
		// Path to your app goes here
		exec.Command("C:\\Program Files\\JetBrains\\JetBrains Rider 2023.2.3\\bin\\rider64.exe"),
		// The Command gets an array of arguments in case your command requires some
		exec.Command("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe", "--profile-directory=Profile 1"),
	}
}

func CmdExecuter(cmd *exec.Cmd) {
	if err := cmd.Start(); err != nil {
		fmt.Printf("%v", err)
	}
}
