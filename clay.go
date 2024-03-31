//go:build js && wasm
// +build js,wasm

package clay

import (
	"syscall/js"
)

type NextFunc func(ctx *Context)

type Handler func(ctx *Context, next NextFunc) Response

type app struct {
	router *RegexpRouter
}

func NewApp() *app {
	return &app{
		router: NewRegexpRouter(),
	}
}

func (a *app) Get(path string, handler Handler) {
	a.router.Register("GET", path, handler)
}

func (a *app) Post(path string, handler Handler) {
	a.router.Register("POST", path, handler)
}

func (a *app) Put(path string, handler Handler) {
	a.router.Register("PUT", path, handler)
}

func (a *app) Delete(path string, handler Handler) {
	a.router.Register("DELETE", path, handler)
}

func (a *app) Patch(path string, handler Handler) {
	a.router.Register("PATCH", path, handler)
}

func (app *app) Listen() {

	c := make(chan struct{}, 0)

	// register global route handler
	js.Global().Set("gohandle", js.FuncOf(app.router.ServeHTTP))
	js.Global().Set("wrapfunc", js.FuncOf(wrap_request))
	<-c
}
