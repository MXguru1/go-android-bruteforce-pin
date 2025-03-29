package hid

import (
	"fmt"
	accessory "github.com/Tryanks/go-accessoryhid"
	"github.com/jeanbritz/go-android-bruteforce-pin.git/pkg/utils"
)

type Touchscreen struct {
	Accessory *accessory.Accessory
}

type Actions interface {
	SetPosition(x int16, y int16)
	Press()
	Swipe(x1 int16, y1 int16, x2 int16, y2 int16)
}

func (t *Touchscreen) SetPosition(x int16, y int16) {
	xMsb := utils.GetMSB(x)
	xLsb := utils.GetLSB(x)
	yMsb := utils.GetMSB(y)
	yLsb := utils.GetLSB(y)

	// Use Pointer to set absolute coordinates
	err := t.Accessory.SendEvent([]byte{
		0x02, byte(xLsb), byte(xMsb), byte(yLsb), byte(yMsb),
	})

	if err != nil {
		err = fmt.Errorf("error occurred while setting pointer position %w", err)
		fmt.Println(err)
	}
}

func (t *Touchscreen) Press() {
	// Convert Pointer to Touch Accessory
	err := t.Accessory.SendEvent([]byte{
		0x01, 0, 0, 0, 0,
	})

	if err != nil {
		err = fmt.Errorf("error occurred while converting %w", err)
		fmt.Println(err)
	}
	// Press
	err = t.Accessory.SendEvent([]byte{
		0x00, 0, 0, 0, 0,
	})

	if err != nil {
		err = fmt.Errorf("error occurred while trying to press %w", err)
		fmt.Println(err)
	}
}

func (t *Touchscreen) Swipe(x1 int16, y1 int16, x2 int16, y2 int16) {
	x1Msb := utils.GetMSB(x1)
	x1Lsb := utils.GetLSB(x1)
	y1Msb := utils.GetMSB(y1)
	y1Lsb := utils.GetLSB(y1)

	// Start Swipe
	// 0x03 = 0x00000011 = (Tip Switch = 1, In Range = 1)
	err := t.Accessory.SendEvent([]byte{
		0x03, byte(x1Lsb), byte(x1Msb), byte(y1Lsb), byte(y1Msb),
	})

	if err != nil {
		err = fmt.Errorf("error occurred %w", err)
		fmt.Println(err)
	}

	x2Msb := utils.GetMSB(x2)
	x2Lsb := utils.GetLSB(x2)
	y2Msb := utils.GetMSB(y2)
	y2Lsb := utils.GetLSB(y2)

	// Move
	// 0x03 - Same flags as above
	err = t.Accessory.SendEvent([]byte{
		0x03, byte(x2Lsb), byte(x2Msb), byte(y2Lsb), byte(y2Msb),
	})

	if err != nil {
		err = fmt.Errorf("error occurred %w", err)
		fmt.Println(err)
	}

	// End Swipe
	// 0x01 = 0x00000001 = (Tip Switch = 0, In Range = 1)
	err = t.Accessory.SendEvent([]byte{
		0x01, byte(x2Lsb), byte(x2Msb), byte(y2Lsb), byte(y2Msb),
	})

	// Fully lifted
	// 0x00 = 0x00000000 = (Tip Switch = 0, In Range = 0)
	err = t.Accessory.SendEvent([]byte{
		0x00, byte(x2Lsb), byte(x2Msb), byte(y2Lsb), byte(y2Msb),
	})
}
