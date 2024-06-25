//go:build js && wasm
// +build js,wasm

package clay

import (
	"syscall/js"
)

type Context struct {
	Request *Request
}

func (ctx *Context) JSON(v map[string]string, status ...int) Response {
	statusCode := 200
	if len(status) > 0 {
		statusCode = status[0]
	}

	jsObject := js.Global().Get("Object").New()
	for key, value := range v {
		jsObject.Set(key, value)
	}

	jsonString := js.Global().Get("JSON").Call("stringify", jsObject).String()

	return Response{
		Body:   jsonString,
		Status: statusCode,
		Headers: map[string]interface{}{
			"Content-Type": "application/json; charset=utf-8",
		},
	}
}

func (ctx *Context) Text(content interface{}, status ...int) Response {

	statusCode := 200
	if len(status) > 0 {
		statusCode = status[0]
	}

	return Response{
		Body:   content,
		Status: statusCode,
		Headers: map[string]interface{}{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}
}

func (ctx *Context) Redirect(destination string, status ...int) Response {
	statusCode := 302
	if len(status) > 0 {
		statusCode = status[0]
	}

	return Response{
		Body:   "",
		Status: statusCode,
		Headers: map[string]interface{}{
			"Content-Type": "text/plain; charset=utf-8",
			"Location":     destination,
		},
	}

}

func (res *Context) HTML(htmlString string, status ...int) Response {
	statusCode := 200
	if len(status) > 0 {
		statusCode = status[0]
	}

	return Response{
		Body:   htmlString,
		Status: statusCode,
		Headers: map[string]interface{}{
			"Content-Type": "text/html; charset=utf-8",
		},
	}
}
