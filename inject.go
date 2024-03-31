//go:build js && wasm
// +build js,wasm

package clay

import (
	"strings"
	"syscall/js"
)

var (
	promise = js.Global().Get("Promise")
)

func newPromise() (p js.Value, set func(js.Value)) {
	ch := make(chan js.Value)
	resolver := make(chan js.Value, 1)
	go func() {
		result := <-ch
		resolve := <-resolver
		resolve.Invoke(result)
	}()
	p = promise.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolver <- args[0]
		return nil
	}))
	set = func(v js.Value) {
		ch <- v
	}
	return
}

func await(awaitable js.Value) chan []js.Value {
	ch := make(chan []js.Value)
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args
		return nil
	})
	awaitable.Call("then", cb)
	return ch
}

func wrap_request(this js.Value, args []js.Value) interface{} {
	request := args[0]
	contentType := request.Get("headers").Call("get", "Content-Type").String()
	p, set := newPromise()

	urlObj := js.Global().Get("URL").New(request.Get("url").String())
	method := request.Get("method").String()
	searchParams := urlObj.Get("searchParams")

	queries := make(map[string]interface{})

	it := searchParams.Call("entries")

	for {
		result := it.Call("next")
		if result.Get("done").Bool() {
			break
		}
		key := result.Get("value").Index(0).String()
		value := result.Get("value").Index(1).String()
		queries[key] = value
	}

	headers := make(map[string]interface{})
	requestHeaders := request.Get("headers")
	requestHeaders.Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		headers[args[0].String()] = args[1].String()
		return nil
	}))

	newReq := map[string]interface{}{
		"method":  method,
		"url":     request.Get("url").String(),
		"queries": queries,
		"headers": headers,
	}
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		ch := await(request.Call("json"))

		go func() {
			results := <-ch
			// jsBody := js.Global().Get("JSON").Call("parse", results[0])

			newReq["body"] = results[0]
			set(js.ValueOf(newReq))
		}()
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		ch := await(request.Call("formData"))

		go func() {
			results := <-ch
			body := make(map[string]interface{})
			iter := results[0].Call("entries")
			for {
				next := iter.Call("next")
				done := next.Get("done").Bool()
				if done {
					break
				}

				val := next.Get("value")
				key := val.Index(0).String()
				value := val.Index(1).String()
				body[key] = value
			}
			newReq["body"] = body
			set(js.ValueOf(newReq))
		}()

	case strings.HasPrefix(contentType, "multipart/form-data"):
		ch := await(request.Call("formData"))

		go func() {
			results := <-ch
			body := make(map[string]interface{})
			iter := results[0].Call("entries")
			for {
				next := iter.Call("next")
				done := next.Get("done").Bool()
				if done {
					break
				}

				val := next.Get("value")
				key := val.Index(0).String()

				var value interface{}

				if val.Index(1).Type() == js.TypeObject {
					fileConstructor := val.Index(1).Get("constructor")
					if fileConstructor.Type() == js.TypeObject {
						constructorName := fileConstructor.Get("name").String()
						file := constructorName == "File"
						if file {
							fileValue := val.Index(1)
							fileReader := js.Global().Get("FileReader").New()
							fileReader.Call("readAsDataURL", fileValue)
							value = <-await(fileReader.Call("load"))
							base64Data := strings.SplitN(value.(string), ",", 2)[1]
							value = base64Data
						} else {
							value = val.Index(1).String()
						}
					}
				} else {
					value = val.Index(1).String()
				}
				body[key] = value
			}
			newReq["body"] = body
			set(js.ValueOf(newReq))
		}()
	case strings.HasPrefix(contentType, "text/plain"):
		ch := await(request.Call("text"))

		go func() {
			results := <-ch
			newReq["body"] = results[0].String()
			set(js.ValueOf(newReq))
		}()

	default:
		ch := await(args[0].Call("arrayBuffer"))
		go func() {
			results := <-ch

			jsBody := js.Global().Get("Uint8Array").New(results[0])

			body := make([]byte, jsBody.Get("length").Int())
			js.CopyBytesToGo(body, jsBody)

			newReq["body"] = string(body)

			set(js.ValueOf(newReq))
		}()
	}

	return p
}
