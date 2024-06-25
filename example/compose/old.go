package main

// //go:build js && wasm
// // +build js,wasm

// package main

// import (
// 	"fmt"
// 	"syscall/js"
// )

// var (
// 	fetch   = js.Global().Get("fetch")
// 	promise = js.Global().Get("Promise")
// )

// func newPromise() (p js.Value, set func(js.Value)) {
// 	ch := make(chan js.Value)
// 	resolver := make(chan js.Value, 1)
// 	go func() {
// 		result := <-ch
// 		resolve := <-resolver
// 		resolve.Invoke(result)
// 	}()
// 	p = promise.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		resolver <- args[0]
// 		return nil
// 	}))
// 	set = func(v js.Value) {
// 		ch <- v
// 	}
// 	return
// }

// func await(awaitable js.Value) chan []js.Value {
// 	ch := make(chan []js.Value)
// 	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		ch <- args
// 		return nil
// 	})
// 	awaitable.Call("then", cb)
// 	return ch
// }

// func main() {

// 	// app := clay.NewApp()

// 	//	app.Get("/", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	// log.Println("Hello World")
// 	// 	htmlContent := `
// 	// 	<!DOCTYPE html>
// 	// 	<html>
// 	// 	<head>
// 	// 		<title>To-Do List</title>
// 	// 		<style>
// 	// 			body { font-family: Arial, sans-serif; margin: 20px; }
// 	// 			#todoInput { padding: 10px; width: 200px; margin-right: 10px; }
// 	// 			#addButton { padding: 10px; }
// 	// 			#todoList { margin-top: 20px; }
// 	// 		</style>
// 	// 	</head>
// 	// 	<body>
// 	// 		<input type="text" id="todoInput" placeholder="Add a new task">
// 	// 		<button id="addButton">Add Task</button>
// 	// 		<ul id="todoList"></ul>

// 	// 		<script>
// 	// 			document.getElementById("addButton").onclick = function() {
// 	// 				var input = document.getElementById("todoInput");
// 	// 				var newTodo = input.value;
// 	// 				if (newTodo) {
// 	// 					var li = document.createElement("li");
// 	// 					li.textContent = newTodo;
// 	// 					document.getElementById("todoList").appendChild(li);
// 	// 					input.value = ""; // Clear the input
// 	// 				}
// 	// 			};
// 	// 		</script>
// 	// 	</body>
// 	// 	</html>
// 	// 	`
// 	// 	return ctx.HTML(htmlContent)
// 	// })

// 	// app.Get("/hey", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	res := map[string]string{
// 	// 		"message": "hello world!",
// 	// 	}

// 	// 	return ctx.JSON(res, 200)
// 	// })

// 	// app.Get("/lorem", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	return ctx.Text("ipsum dot")
// 	// })

// 	// app.Get("/ping", func(ctx *clay.Context, next clay.NextFunc) clay.Response {

// 	// 	return ctx.Redirect("https://google.com", 301)
// 	// })

// 	// app.Post("/todo", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	return ctx.Text("Todo Created", 201)
// 	// })

// 	// app.Put("/todo", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	return ctx.Text("Todo Updated", 201)
// 	// })

// 	// app.Delete("/todo", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	return ctx.Text("Todo Delete", 201)
// 	// })

// 	// app.Patch("/todo", func(ctx *clay.Context, next clay.NextFunc) clay.Response {
// 	// 	return ctx.Text("Todo Patched", 200)
// 	// })

// 	// app.Post("/body", func(ctx *clay.Context, next clay.NextFunc) clay.Response {

// 	// 	now := time.Now()
// 	// 	dateConstructor := js.Global().Get("Date")
// 	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
// 	// if err != nil {
// 	// 	js.Global().Get("console").Call("log", err.Error())
// 	// }
// 	// ch := await(fetch.Invoke("https://jsonplaceholder.typicode.com/todos/1"))
// 	// js.Global().Get("console").Call("log", res.StatusCode)
// 	// waitorr := make(chan string)
// 	// p, set := newPromise()
// 	// go func() {
// 	// 	results := <-ch
// 	// 	js.Global().Get("console").Call("log", results)

// 	// 	rsp := results[0]

// 	// 	results = <-await(rsp.Call("json"))
// 	// 	js.Global().Get("console").Call("log", results[0].Call("toString"))

// 	// 	waitorr <- results[0].Call("toString").String()
// 	// 	js.Global().Get("console").Call("log", <-waitorr)
// 	// 	// set(js.ValueOf(ctx.Text(<-waitorr, 201)))
// 	// }()

// 	// return ctx.Text(dateConstructor.New(now.Unix()*1000), 201)
// 	// })

// 	// app.Listen()
// 	js.Global().Set("sayHello", js.FuncOf(sayHello))
// 	<-make(chan struct{})
// }

// //go:export
// func sayHello(this js.Value, args []js.Value) interface{} {
// 	fmt.Println("Go Web Assembly")
// 	return nil
// }

// package main

// import (
// 	"strings"
// 	"syscall/js"

// 	"github.com/uchexgod/clay"
// )

// func main() {

// 	clay.NewApp()
// 	js.Global().Set("gohandle", js.FuncOf(handle_fetch))
// 	js.Global().Set("wrapfunc", js.FuncOf(wrap_request))
// 	select {}
// }

// var (
// 	fetch   = js.Global().Get("fetch")
// 	promise = js.Global().Get("Promise")
// )

// func newPromise() (p js.Value, set func(js.Value)) {
// 	ch := make(chan js.Value)
// 	resolver := make(chan js.Value, 1)
// 	go func() {
// 		result := <-ch
// 		resolve := <-resolver
// 		resolve.Invoke(result)
// 	}()
// 	p = promise.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		resolver <- args[0]
// 		return nil
// 	}))
// 	set = func(v js.Value) {
// 		ch <- v
// 	}
// 	return
// }

// func await(awaitable js.Value) chan []js.Value {
// 	ch := make(chan []js.Value)
// 	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		ch <- args
// 		return nil
// 	})
// 	awaitable.Call("then", cb)
// 	return ch
// }

// func text(req js.Value) string {
// 	res := make(chan string, 1)

// 	awaitable := req.Call("text")
// 	ch := make(chan []js.Value)
// 	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		ch <- args
// 		return nil
// 	})

// 	awaitable.Call("then", cb)

// 	go func() {
// 		results := <-ch
// 		// js.Global().Get("console").Call("log", results[0])
// 		res <- results[0].String()
// 	}()

// 	return <-res
// }

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func handle_fetch(this js.Value, args []js.Value) interface{} {

// 	// ch := await(args[0].Call("text"))
// 	p, set := newPromise()

// 	go func() {
// 		// results := <-ch
// 		set(js.ValueOf("hiiii"))
// 	}()

// 	body := args[0].Get("body")
// 	js.Global().Get("console").Call("log", body.Get("name"))
// 	js.Global().Get("console").Call("log", body)
// 	// err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&user)
// 	// // err := json.Unmarshal(body, &user)
// 	// if err != nil {
// 	// 	js.Global().Get("console").Call("log", err.Error())
// 	// }

// 	// defer func() {
// 	// 	js.Global().Get("console").Call("log", p.String())
// 	// }()

// 	// response := js.Global().Get("Response")
// 	// blobType := "text/plain"
// 	// responseBody := js.Global().Get("Blob").New([]interface{}{p}, map[string]interface{}{"type": blobType})
// 	// return response.New(responseBody, map[string]interface{}{"status": 200})

// 	return p
// }

// func wrap_request(this js.Value, args []js.Value) interface{} {
// 	request := args[0]
// 	contentType := request.Get("headers").Call("get", "Content-Type").String()
// 	p, set := newPromise()

// 	urlObj := js.Global().Get("URL").New(request.Get("url").String())
// 	method := request.Get("method").String()
// 	searchParams := urlObj.Get("searchParams")

// 	queries := make(map[string]string)

// 	it := searchParams.Call("entries")

// 	for {
// 		result := it.Call("next")
// 		if result.Get("done").Bool() {
// 			break
// 		}
// 		key := result.Get("value").Index(0).String()
// 		value := result.Get("value").Index(1).String()
// 		queries[key] = value
// 	}

// 	headers := make(map[string]string)
// 	requestHeaders := request.Get("headers")
// 	requestHeaders.Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		headers[args[0].String()] = args[1].String()
// 		return nil
// 	}))

// 	newReq := map[string]interface{}{
// 		"method":  method,
// 		"url":     urlObj,
// 		"queries": queries,
// 		"headers": headers,
// 	}
// 	switch {
// 	case strings.HasPrefix(contentType, "application/json"):
// 		ch := await(request.Call("json"))

// 		go func() {
// 			results := <-ch
// 			// jsBody := js.Global().Get("JSON").Call("parse", results[0])

// 			newReq["body"] = results[0]
// 			set(js.ValueOf(newReq))
// 		}()
// 	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
// 		ch := await(request.Call("formData"))

// 		go func() {
// 			results := <-ch
// 			body := make(map[string]interface{})
// 			iter := results[0].Call("entries")
// 			for {
// 				next := iter.Call("next")
// 				done := next.Get("done").Bool()
// 				if done {
// 					break
// 				}

// 				val := next.Get("value")
// 				key := val.Index(0).String()
// 				value := val.Index(1).String()
// 				body[key] = value
// 			}
// 			newReq["body"] = body
// 			set(js.ValueOf(newReq))
// 		}()

// 	case strings.HasPrefix(contentType, "multipart/form-data"):
// 		ch := await(request.Call("formData"))

// 		go func() {
// 			results := <-ch
// 			body := make(map[string]interface{})
// 			iter := results[0].Call("entries")
// 			for {
// 				next := iter.Call("next")
// 				done := next.Get("done").Bool()
// 				if done {
// 					break
// 				}

// 				val := next.Get("value")
// 				key := val.Index(0).String()

// 				var value interface{}

// 				if val.Index(1).Type() == js.TypeObject {
// 					fileConstructor := val.Index(1).Get("constructor")
// 					if fileConstructor.Type() == js.TypeObject {
// 						constructorName := fileConstructor.Get("name").String()
// 						file := constructorName == "File"
// 						if file {
// 							fileValue := val.Index(1)
// 							fileReader := js.Global().Get("FileReader").New()
// 							fileReader.Call("readAsDataURL", fileValue)
// 							value = <-await(fileReader.Call("load"))
// 							base64Data := strings.SplitN(value.(string), ",", 2)[1]
// 							value = base64Data
// 						} else {
// 							value = val.Index(1).String()
// 						}
// 					}
// 				} else {
// 					value = val.Index(1).String()
// 				}
// 				body[key] = value
// 			}
// 			newReq["body"] = body
// 			set(js.ValueOf(newReq))
// 		}()
// 	case strings.HasPrefix(contentType, "text/plain"):
// 		ch := await(request.Call("text"))

// 		go func() {
// 			results := <-ch
// 			newReq["body"] = results[0].String()
// 			set(js.ValueOf(newReq))
// 		}()

// 	default:
// 		ch := await(args[0].Call("arrayBuffer"))
// 		go func() {
// 			results := <-ch

// 			jsBody := js.Global().Get("Uint8Array").New(results[0])

// 			body := make([]byte, jsBody.Get("length").Int())
// 			js.CopyBytesToGo(body, jsBody)

// 			newReq["body"] = string(body)

// 			set(js.ValueOf(newReq))
// 		}()
// 	}

// 	return p
// }
