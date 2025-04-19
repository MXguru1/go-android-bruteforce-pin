package hid

import (
	"fmt"
	accessory "github.com/Tryanks/go-accessoryhid"
	"time"
)

type Swiper struct {
	Accessory    *accessory.Accessory
	ScreenWidth  int
	ScreenHeight int
	StepDuration time.Duration
}

type SwipeActions interface {
	PerformSwipe(startX int, startY int, endX int, endY int, steps int)
	SendHidEvent(tipSwitch bool, contactID bool, x uint16, y uint16)
}

func (s *Swiper) PerformSwipe(startX int, startY int, endX int, endY int, steps int) {

	// Generate intermediate points
	points := interpolatePoints(startX, startY, endX, endY, steps)

	// Send reports for touch down and movement
	for _, point := range points {
		x, y := scaleCoordinates(point[0], point[1], s.ScreenWidth, s.ScreenHeight)
		s.SendHidEvent(true, false, x, y)
		time.Sleep(s.StepDuration)
	}

	// Send lift-off report (Tip Switch = false)
	x, y := scaleCoordinates(endX, endY, s.ScreenWidth, s.ScreenHeight)
	s.SendHidEvent(false, false, x, y)
}

func (s *Swiper) SendHidEvent(tipSwitch bool, contactID bool, x uint16, y uint16) {
	report := make([]byte, 6)

	if tipSwitch {
		report[0] |= 0x01
	}
	if contactID {
		report[0] |= 0x02
	}

	// Byte 1: Padding (always 0)
	report[1] = 0x00

	// Bytes 2-3: X-coordinate (big-endian)
	report[2] = byte(x >> 8)   // High byte
	report[3] = byte(x & 0xFF) // Low byte
	// Bytes 4-5: Y-coordinate (big-endian)
	report[4] = byte(y >> 8)
	report[5] = byte(y & 0xFF)

	fmt.Printf("Sending HID Event: { %x } { %x } { %x - %x - %x - %x }\n", report[0], report[1], report[2], report[3], report[4], report[5])
	err := s.Accessory.SendEvent(report[:])
	if err != nil {
		err = fmt.Errorf("error occurred while sending HID Report %w", err)
		fmt.Println(err)
	}
}

// scaleCoordinates converts pixel coordinates to the HID descriptor's range (0-32767)
func scaleCoordinates(px int, py int, screenWidth int, screenHeight int) (uint16, uint16) {
	const maxHID = 32767
	x := uint16((float64(px) / float64(screenWidth)) * float64(maxHID))
	y := uint16((float64(py) / float64(screenHeight)) * float64(maxHID))
	return x, y
}

func interpolatePoints(startX int, startY int, endX int, endY int, steps int) [][2]int {
	points := make([][2]int, steps)
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)
		points[i][0] = startX + int(float64(endX-startX)*t) // X
		points[i][1] = startY + int(float64(endY-startY)*t) // Y
	}
	return points
}
