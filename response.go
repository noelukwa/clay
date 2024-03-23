//go:build js && wasm
// +build js,wasm

package clay

import (
	"strings"
	"syscall/js"
)

type Response struct {
	Body    interface{}
	Headers map[string]interface{}
	Status  int
}

func (res Response) write() js.Value {
	responseConstructor := js.Global().Get("Response")

	blobType := "text/plain"

	if contentType, ok := res.Headers["Content-Type"]; ok {

		str, ok := contentType.(string)
		if ok {
			parts := strings.Split(str, ";")
			if len(parts) > 0 {
				blobType = strings.TrimSpace(parts[0])
			}
		}

	}

	responseBody := js.Global().Get("Blob").New([]interface{}{res.Body}, map[string]interface{}{"type": blobType})
	return responseConstructor.New(responseBody, map[string]interface{}{
		"status":  res.Status,
		"headers": res.Headers,
	})
}
