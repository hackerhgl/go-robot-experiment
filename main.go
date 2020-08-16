package main

import (
	"bytes"
	"image/jpeg"
	"io/ioutil"
	"log"

	"github.com/go-vgo/robotgo"
	"golang.org/x/image/bmp"
)

func main() {
	print("TESTING")
	robotgo.ActiveName("wowo")

	takeScreenshot()

	// log.Print(text)

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
