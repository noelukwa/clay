//go:build js && wasm
// +build js,wasm

package clay

type Request struct {
	params  map[string]string
	path    string
	queries map[string]string
	headers map[string]string
	body    map[string]string
	Method  string
}

func (r *Request) Param(name string) string {
	if value, ok := r.params[name]; ok {
		return value
	}
	return ""
}

func (r *Request) Query(name string) string {
	if value, ok := r.queries[name]; ok {
		return value
	}
	return ""
}

func (r *Request) Queries() map[string]string {
	return r.queries
}

func (r *Request) Header(name string) string {
	if value, ok := r.headers[name]; ok {
		return value
	}
	return ""
}

func (r *Request) Headers(name string) map[string]string {
	return r.headers
}

func (r *Request) Body() map[string]string {
	return r.headers
}

func (r *Request) Text() (body string) {
	return "Hello"
}
