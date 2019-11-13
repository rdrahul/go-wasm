package main

import (
	"syscall/js"
)

func main() {

	//getting the document object
	document := js.Global().Get("document")

	//create a p tag by calling createElement
	p := document.Call("createElement", "p")

	//set the innerHtml attribute for that p
	p.Set("innerHTML", "Hello WASM from Go!")

	//append the created p tag to our document
	document.Get("body").Call("appendChild", p)
}
