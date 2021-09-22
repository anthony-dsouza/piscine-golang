package main

import "github.com/01-edu/z01"

type Door struct {
	State string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.State = "CLOSE"
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.State = "OPEN"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.State == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.State == "ClOSE"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.State == "OPEN" {
		CloseDoor(door)
	}
}
