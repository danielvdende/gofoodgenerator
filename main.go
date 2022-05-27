package main

import (
	"fmt"
	webserver "gofoodgenerator/src"
)

func main() {
	fmt.Println("Starting webserver")
	webserver.Run()
}
