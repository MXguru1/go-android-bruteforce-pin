package main

import (
	accessory "github.com/Tryanks/go-accessoryhid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/hid"
	"log"
	"os"
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
		logger.Println("Found Android HID device:" + devices[0].Manufacturer)
	} else {
		logger.Println("Did not find any HID device")
		os.Exit(0)
	}

	phone := devices[0]
	defer phone.Close()

	touch, err := phone.Register(hid.TouchscreenReportDesc)
	if err != nil {
		logger.Fatalln(err)
	}

	touchscreen := hid.Touchscreen{
		Accessory: touch,
	}

	// Start X,Y coordinates
	x1 := 5000
	y1 := 7500

	// End X,Y coordinates
	x2 := 5000
	y2 := 4000

	logger.Println("Attempting to perform swipe")
	touchscreen.Swipe(int16(x1), int16(y1), int16(x2), int16(y2))
}
