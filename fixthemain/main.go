package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const OPEN = false
const CLOSE = true

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func OpenDoor(ptdrDoor *Door) {
	PrintStr("Door Open...")
	ptdrDoor.state = OPEN

}

func IsDoorOpen(ptdrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptdrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
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
