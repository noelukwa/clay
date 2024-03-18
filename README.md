# clay

_[Experimental]_ lightweight framework for writing cloudflare workers with go

```go
//go:build js && wasm
// +build js,wasm

package main

import (
 "github.com/uchexgod/clay"
)

func main() {

  app := clay.NewApp()

 //send html response
  app.Get("/", func(ctx *clay.Context, next clay.NextFunc) interface{} {

  htmlContent := `
    <!DOCTYPE html>
    <html>
    <head>
    <title>To-Do List</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        #todoInput { padding: 10px; width: 200px; margin-right: 10px; }
        #addButton { padding: 10px; }
        #todoList { margin-top: 20px; }
    </style>
    </head>
    <body>
        <input type="text" id="todoInput" placeholder="Add a new task">
        <button id="addButton">Add Task</button>
        <ul id="todoList"></ul>
        
        <script>
            document.getElementById("addButton").onclick = function() {
                var input = document.getElementById("todoInput");
                var newTodo = input.value;
                if (newTodo) {
                    var li = document.createElement("li");
                    li.textContent = newTodo;
                    document.getElementById("todoList").appendChild(li);
                    input.value = ""; // Clear the input
                }
            };
        </script>
    </body>
    </html>
    `
  return ctx.HTML(htmlContent)
 })

 //send text response
 app.Post("/todo", func(ctx *clay.Context, next clay.NextFunc) interface{} {
  return ctx.Text("Todo Created", 201)
 })


 app.Listen()

}

```
