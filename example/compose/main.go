//js && wasm
//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func main() {

	js.Global().Set("sayHello", js.FuncOf(sayHello))
	<-make(chan struct{})
}

func sayHello(this js.Value, args []js.Value) interface{} {

	return nil
}
