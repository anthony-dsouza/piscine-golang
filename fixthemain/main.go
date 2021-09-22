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
	PrintStr("Door Closing...")
	ptrDoor.State = "CLOSE"
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.State = "OPEN"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	if ptrDoor.State == "OPEN" {
		return true
	}
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.State == "ClOSE" {
		return true
	}
	return false
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
