package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type AppDimensions struct {
	Width  int `json:width`
	Height int `json:height`
	X      int `json:x`
	Y      int `json:y`
}

func main() {
	start := time.Now()

	var action string = strings.TrimSpace(os.Args[1])

	if action == "maximize" {
		if runtime.GOOS != "darwin" {
			println("OS not supported")
			return
		}
		var window string = strings.TrimSpace(os.Args[2])
		var appID string = strings.TrimSpace(os.Args[3])
		robotgo.ActiveName(window)

		pid := robotgo.GetPID()
		robotgo.MaxWindow(pid)

		myself, _ := user.Current()
		path := "/Users/" + myself.Username + "/Library/Containers/" + appID + "/Data/Documents/dimensions.json"
		println(myself.Username)
		println(path)

		rawDimensions, err := ioutil.ReadFile(path)

		println(rawDimensions)

		var dimensions AppDimensions

		err = json.Unmarshal(rawDimensions, &dimensions)

		if err != nil {
			println("UNMARSAHLL ERROR")
			println(err.Error())
		}
		println("dimensions [WHXY]", dimensions.Width, dimensions.Height, dimensions.X, dimensions.Y)
		mousex, mousey := robotgo.GetMousePos()
		println("PRE mousex", mousex, "mousey", mousey)
		time.Sleep(time.Second)
		robotgo.MoveMouseSmooth(dimensions.X+100, dimensions.Y+10)
		print("10")
		time.Sleep(time.Second)
		robotgo.MoveMouseSmooth(dimensions.X+100, dimensions.Y+20)
		print("20")
		time.Sleep(time.Second)
		robotgo.MoveMouseSmooth(dimensions.X+100, dimensions.Y+30)
		print("30")
		time.Sleep(time.Second)
		robotgo.MoveMouseSmooth(dimensions.X+100, dimensions.Y+40)
		print("40")
		time.Sleep(time.Second)
		robotgo.MoveMouseSmooth(dimensions.X+100, dimensions.Y+50)
		print("50")
		mouse1x, mouse1y := robotgo.GetMousePos()
		println("post mousex", mouse1x, "mousey", mouse1y)

		time.Sleep(time.Second)
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
