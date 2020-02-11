package main

import (
	"fmt"
	"leb.io/siginfo"
	"os"
)

func sig(s os.Signal) {
	fmt.Printf("\ncontrol C\ns=%#v, %s\n", s, s.String())
}

func main() {
	siginfo.SetHandler("SIGINT", sig)
	for {

	}
}
