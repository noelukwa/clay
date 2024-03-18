//go:build js && wasm
// +build js,wasm

package clay

import "syscall/js"

type Context struct {
	Request *Request
}

func (ctx *Context) JSON(v map[string]string, status int) interface{} {
	jsObject := js.Global().Get("Object").New()
	for key, value := range v {
		jsObject.Set(key, value)
	}

	jsonString := js.Global().Get("JSON").Call("stringify", jsObject).String()

	responseBody := js.Global().Get("Blob").New([]interface{}{jsonString}, map[string]interface{}{"type": "application/json"})

	responseConstructor := js.Global().Get("Response")
	response := responseConstructor.New(responseBody, map[string]interface{}{
		"status": status,
	})
	return response
}

func (ctx *Context) Text(content string, status ...int) interface{} {
	// Default status code is 200 if not specified
	statusCode := 200
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Create a Blob with the text content
	responseBody := js.Global().Get("Blob").New([]interface{}{content}, map[string]interface{}{"type": "text/plain"})

	// Construct the Response object with the Blob and status code
	responseConstructor := js.Global().Get("Response")
	response := responseConstructor.New(responseBody, map[string]interface{}{
		"status": statusCode,
	})
	return response
}

func (ctx *Context) Redirect(destination string, status ...int) interface{} {
	// Default redirect status is 302 (Found) if not specified
	statusCode := 302
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Manually construct a Response object for the redirect
	responseConstructor := js.Global().Get("Response")
	// An empty body is used because the redirect is indicated by the status code and Location header
	responseBody := js.Global().Get("Blob").New([]interface{}{""}, map[string]interface{}{"type": "text/plain"})
	response := responseConstructor.New(responseBody, map[string]interface{}{
		"status": statusCode,
		"headers": map[string]interface{}{
			"Location": destination,
		},
	})
	return response
}

func (res *Context) HTML(htmlString string) interface{} {
	responseConstructor := js.Global().Get("Response")
	responseBody := js.Global().Get("Blob").New([]interface{}{htmlString}, map[string]interface{}{"type": "text/html"})
	response := responseConstructor.New(responseBody, map[string]interface{}{
		"status": 200,
	})
	return response
}
