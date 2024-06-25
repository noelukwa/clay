//go:build js && wasm
// +build js,wasm

package clay

import (
	"regexp"
	"syscall/js"
)

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

func (r *RegexpRouter) Register(method, pattern string, handler Handler) {
	pattern = regexp.QuoteMeta(pattern)
	pattern = "^" + pattern + "$"
	pattern = regexp.MustCompile(`\\:([a-zA-Z0-9_]+)`).ReplaceAllString(pattern, `([a-zA-Z0-9_]+)`)

	// Include the method in the stored route information
	r.routes = append(r.routes, Route{regexp.MustCompile(pattern), handler, method})
}

func (r *RegexpRouter) ServeHTTP(this js.Value, args []js.Value) interface{} {
	js.Global().Get("console").Call("log", args[0])
	req := &Request{}
	request := args[0]
	method := args[0].Get("method").String()
	urlObj := js.Global().Get("URL").New(request.Get("url").String())
	pathname := urlObj.Get("pathname").String()

	js.Global().Get("console").Call("log", request)
	methodAllowed := false
	for _, route := range r.routes {
		if matches := route.pattern.FindStringSubmatch(pathname); matches != nil {
			if route.method != method {
				methodAllowed = true
				continue
			}

			params := make(map[string]string, 0)
			for i, name := range route.pattern.SubexpNames() {
				if i != 0 && name != "" {
					params[name] = matches[i]
				}
			}

			req.params = params
			req.path = pathname
			req.Method = method

			ctx := &Context{
				Request: req,
			}
			next := func(ctx *Context) {}
			return route.handler(ctx, next).write()
		}
	}

	if methodAllowed {
		return methodNotAllowedResponse()
	}

	return notFoundResponse()
}

func notFoundResponse() js.Value {
	htmlContent := "<h1>404 Not Found</h1>"
	responseConstructor := js.Global().Get("Response")
	responseBody := js.Global().Get("Blob").New([]interface{}{htmlContent}, map[string]interface{}{"type": "text/html"})
	return responseConstructor.New(responseBody, map[string]interface{}{
		"status": 404,
	})
}

func serverErrorResponse() js.Value {
	responseConstructor := js.Global().Get("Response")
	responseBody := js.Global().Get("Blob").New([]interface{}{"Internal Server Error"}, map[string]interface{}{"type": "text/plain"})
	return responseConstructor.New(responseBody, map[string]interface{}{
		"status": 500,
	})
}

func methodNotAllowedResponse() js.Value {
	responseConstructor := js.Global().Get("Response")
	return responseConstructor.New(nil, map[string]interface{}{
		"status": 405,
	})
}
