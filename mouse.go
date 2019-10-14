package wa

import (
	"unsafe"

	"github.com/lxn/win"
)

//GetCursorPos mouse pos
func GetCursorPos() win.POINT {
	pt := win.POINT{X: 0, Y: 0}
	win.GetCursorPos(&pt)
	return pt
}

//SetCursorPos set x, y
func SetCursorPos(x, y int32) {
	win.SetCursorPos(x, y)
}

//MouseClick click
func MouseClick(args ...string) {
	if len(args) == 0 {
		var input, input2 win.MOUSE_INPUT
		input.Mi.DwFlags = win.MOUSEEVENTF_LEFTDOWN
		input2.Mi.DwFlags = win.MOUSEEVENTF_LEFTUP
		win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
		// Sleep(1)
		win.SendInput(1, unsafe.Pointer(&input2), int32(unsafe.Sizeof(input2)))
		return
	}
	switch args[0] {
	case "", "left":
		var input, input2 win.MOUSE_INPUT
		input.Mi.DwFlags = win.MOUSEEVENTF_LEFTDOWN
		input2.Mi.DwFlags = win.MOUSEEVENTF_LEFTUP
		win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
		// Sleep(1)
		win.SendInput(1, unsafe.Pointer(&input2), int32(unsafe.Sizeof(input2)))
		return
	case "right":
		var input, input2 win.MOUSE_INPUT
		input.Mi.DwFlags = win.MOUSEEVENTF_RIGHTDOWN
		input2.Mi.DwFlags = win.MOUSEEVENTF_RIGHTUP
		win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
		// Sleep(1)
		win.SendInput(1, unsafe.Pointer(&input2), int32(unsafe.Sizeof(input2)))
		return
	}
}

//MouseClickPos x, y click
func MouseClickPos(x, y int32) {
	win.SetCursorPos(x, y)
	MouseClick()
}
