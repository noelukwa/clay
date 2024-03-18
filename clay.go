//go:build js && wasm
// +build js,wasm

package clay

import (
	"regexp"
	"syscall/js"
)

type NextFunc func(ctx *Context)

type Handler func(ctx *Context, next NextFunc) interface{}

type Response interface {
	func() interface{}
}

type Route struct {
	pattern *regexp.Regexp
	handler Handler
	method  string
}

type RegexpRouter struct {
	routes []Route
}

func NewRegexpRouter() *RegexpRouter {
	return &RegexpRouter{}
}

type Request struct {
	jsv    js.Value
	params []string
}

type app struct {
	router *RegexpRouter
}

func NewApp() *app {
	return &app{
		router: NewRegexpRouter(),
	}
}

type Context struct {
	Request *Request
}

func (app *app) Listen() {

	c := make(chan struct{}, 0)

	// register global route handler
	<-c
}
