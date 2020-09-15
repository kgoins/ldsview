package internal

import "fmt"

var Version string
var GitCommit string
var BuildDate string

func PrintVersionInfo() {
	fmt.Printf("Version=%s\n", Version)
	fmt.Printf("Commit=%s\n", GitCommit)
	fmt.Printf("BuildDate=%s\n", BuildDate)
}
