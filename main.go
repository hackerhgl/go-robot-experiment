package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"golang.org/x/image/bmp"
)

func main() {
	start := time.Now()

	var action string = strings.TrimSpace(os.Args[1])

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

		// files, err := filepath.Glob("*")
		// if err != nil {
		// 	log.Fatal(err)
		// }

	}

	// robotgo.KeyTap("tab", "alt")
	println("TESTING")

	fmt.Println(time.Since(start))

}

func takeScreenshot() {
	width, height := robotgo.GetScreenSize()

	bitMap := robotgo.CaptureScreen(0, 0, width, height)
	defer robotgo.FreeBitmap(bitMap)

	bs := robotgo.ToBitmapBytes(bitMap)
	img, err := bmp.Decode(bytes.NewReader(bs))
	if err != nil {
		log.Println("bmp.Decode err is: ", err)
		return
	}

	b := new(bytes.Buffer)
	err = jpeg.Encode(b, img, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Println("jpeg.Encode err is: ", err)
		return
	}
	bts := b.Bytes()
	ioutil.WriteFile("images/out.jpg", bts, 0666)
	println("FILE WRITTED")
	// return bts
}
