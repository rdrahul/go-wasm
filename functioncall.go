package main

import (
	"syscall/js"
)

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)

}

func main() {
	js.Global().Set("printMessage", js.FuncOf(printMessage))
}
