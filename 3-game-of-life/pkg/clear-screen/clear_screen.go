package clearscreen

import (
	"os"
	"os/exec"
	"runtime"
)

var clearFunc = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func Clear() {
	value, ok := clearFunc[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("unsupported platform: " + runtime.GOOS)
	}
}
