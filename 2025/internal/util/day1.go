package util

import "fmt"

type Dial struct {
	Start, End, curPos, Range int
}

func NewDial() *Dial {
	d := &Dial{Start: 0, End: 99, curPos: 50}
	d.Range = (d.End - d.Start) + 1

	return d
}

// rotates the dial to the right and returns the current position
func (d *Dial) RotateRight(amount int) int {
	d.curPos = (d.curPos + amount) % d.Range

	return d.curPos
}

// rotates the dial to the left and returns the current position
func (d *Dial) RotateLeft(amount int) int {
	d.curPos -= amount

	if d.curPos < 0 {
		d.curPos = (d.curPos + d.Range) % d.Range
	}

	return d.curPos
}

func (d *Dial) GetRange() int {
	return d.Range
}

func (d *Dial) GetCurrentPosition() int {
	return d.curPos
}

// split the line so that we get the direction of the rotation (L or R) and the amount to rotate
func ParseRotation(line string) (string, int) {
	direction := string(line[0])
	amount := 0

	fmt.Sscanf(line[1:], "%d", &amount)

	return direction, amount
}
