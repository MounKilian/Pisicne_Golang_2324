package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = false
	CLOSE = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	ptrDoor.state = CLOSE
}

func OpenDoor(ptdrDoor *Door) {
	PrintStr("Door Open...")
	z01.PrintRune('\n')
	ptdrDoor.state = OPEN
}

func IsDoorOpen(ptdrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	return ptdrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	return ptrDoor.state
}

func main() {
	door := &Door{}

	door.state = OPEN

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
