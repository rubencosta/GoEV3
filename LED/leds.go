// Provides APIs for interacting with EV3's LEDs.
package LED

import (
	"fmt"
	"log"
	"os"

	"github.com/ldmberman/GoEV3/utilities"
)

// Constants for the LED positions (left and right).
type Position string

const (
	Left  Position = "left"
	Right          = "right"
)

// Constants for the LED colors.
type Color string

const (
	Green Color = "green"
	Red         = "red"
	Amber       = "amber"
)

func findFilename(color Color, position Position) string {
	if color == Amber {
		log.Fatal("Amber colors must be decomposed into green and red")
	}

	filename := fmt.Sprintf("/sys/class/leds/ev3:%s:%s:ev3dev", string(position), string(color))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("Cannot find the LED interface\n", filename)
	}

	return filename
}

// Set Brightness (0 - 255)
func SetBrightness(color Color, position Position, brightness uint8) {
	if color == Amber {
		utilities.WriteIntValue(findFilename(Green, position), "brightness", int64(brightness))
		utilities.WriteIntValue(findFilename(Red, position), "brightness", int64(brightness))
	} else {
		utilities.WriteIntValue(findFilename(color, position), "brightness", int64(brightness))
	}

}

// Turns on the given LED with the specified color.
func TurnOn(color Color, position Position) {
	SetBrightness(color, position, 255)
}

// Turns off the given LED with the specified color.
func TurnOff(color Color, position Position) {
	SetBrightness(color, position, 0)
}
