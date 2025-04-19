package main

import (
	accessory "github.com/Tryanks/go-accessoryhid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/hid"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(
		os.Stdout,
		"main: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	devices, err := accessory.GetDevices(2)
	if err != nil {
		logger.Fatalln(err)
	}
	if len(devices) > 0 {
		logger.Println("Found Android HID device: " + devices[0].Manufacturer)
	} else {
		logger.Println("Did not find any HID device")
		os.Exit(0)
	}

	phone := devices[0]
	defer phone.Close()

	touch, err := phone.Register(hid.TouchscreenReportDesc2)
	if err != nil {
		logger.Fatalln(err)
	}
	swiper := hid.Swiper{
		Accessory:    touch,
		ScreenHeight: 1280,
		ScreenWidth:  720,
		StepDuration: time.Second,
	}

	time.Sleep(2 * time.Second) // This necessary to avoid pipe errors

	// Start X,Y coordinates
	//x1 := (200 / 1920) * 32767
	//y1 := (300 / 1080) * 32767
	//
	//// End X,Y coordinates
	//x2 := (600 / 1920) * 32767
	//y2 := (700 / 1080) * 32767
	//
	//x3 := 5000
	//y3 := 1000

	logger.Println("Attempting to perform swipe")
	swiper.PerformSwipe(200, 300, 600, 700, 5)
}
