package drr

import "net/http"

func (r *Route) Get(handler http.Handler) {
	r.method = http.MethodGet
	r.Handler(handler)
}

func (r *Route) Head(handler http.Handler) {
	r.method = http.MethodHead
	r.Handler(handler)
}

func (r *Route) Post(handler http.Handler) {
	r.method = http.MethodPost
	r.Handler(handler)
}

func (r *Route) Put(handler http.Handler) {
	r.method = http.MethodPut
	r.Handler(handler)
}

func (r *Route) Patch(handler http.Handler) {
	r.method = http.MethodPatch
	r.Handler(handler)
}

func (r *Route) Delete(handler http.Handler) {
	r.method = http.MethodDelete
	r.Handler(handler)
}

func (r *Route) Connect(handler http.Handler) {
	r.method = http.MethodConnect
	r.Handler(handler)
}

func (r *Route) Options(handler http.Handler) {
	r.method = http.MethodOptions
	r.Handler(handler)
}

func (r *Route) Trace(handler http.Handler) {
	r.method = http.MethodTrace
	r.Handler(handler)
}
