package logger

import (
	"fmt"
)

var Version string = "0.2.0"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}

func Warn(msg string) {
	fmt.Println("[WARN] ", msg)
}
