package main

import (
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)

	exit()
	return nil
}

func callbackMessage(this js.Value, inputs []js.Value) interface{} {
	callback := inputs[len(inputs)-1:][0]
	message := inputs[0].String()

	callback.Invoke(js.Null(), "Did you mean "+message)
	return nil
}

func exit() interface{} {
	c <- true
	return nil
}

func main() {
	//created a channel
	c = make(chan bool)

	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("callbackMessage", js.FuncOf(callbackMessage))

	//listening for it
	<-c
}
