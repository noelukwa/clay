//go:build js && wasm
// +build js,wasm

package clay

import "syscall/js"

type Request struct {
	jsv    js.Value
	params []string
}

// Param retrieves path parameters by name.
func (r *Request) Param(name string) string {
	// Implementation omitted for brevity; it would match names to captured groups.
	return ""
}
