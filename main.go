package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	start := time.Now()

	var action string = strings.TrimSpace(os.Args[1])

	if action == "maximize" {
		var window string = strings.TrimSpace(os.Args[2])
		robotgo.ActiveName(window)

		pid := robotgo.GetPID()
		x, y, width, height := robotgo.GetBounds(pid)
		println("x:", x)
		println("y:", y)
		println("width:", width)
		println("height:", height)

		offX := (width / 2) + x
		offY := 15 + x

		robotgo.Move(offX, offY)

		robotgo.MouseClick("left", true)

		// println("PID", pid)
		// active, _ := robotgo.FindIds(window)
		// for i, id := range active {
		// 	println(i, id)
		// }
		// robotgo.MaxWindow(pid)
	}

	if action == "activate" {
		var window string = strings.TrimSpace(os.Args[2])
		robotgo.ActiveName(window)
	}
	if action == "screenshot" {
		if runtime.GOOS != "darwin" {
			println("OS not supported")
			return
		}
		robotgo.KeyTap("3", "shift", "command")
		time.Sleep(1 * time.Second)

		myself, _ := user.Current()
		desktop := myself.HomeDir + "/Desktop/"
		files, _ := ioutil.ReadDir(desktop)

		fileName := strings.TrimSpace(os.Args[2])

		filtered := []string{}

		for _, file := range files {
			name := file.Name()
			if file.IsDir() {
				continue
			}
			if !strings.Contains(name, ".png") {
				continue
			}
			filtered = append(filtered, file.Name())
		}

		filepath := desktop + filtered[len(filtered)-1]

		file, _ := ioutil.ReadFile(filepath)

		err := ioutil.WriteFile(fileName+".png", file, 0777)

		if err != nil {
			fmt.Println(err)
		}

		os.Remove(filepath)

	}

	fmt.Println(time.Since(start))

}
