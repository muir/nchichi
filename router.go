package nchichi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muir/nchi"
)

type router struct {
	*nchi.Mux
}

func AsRouter(m *nchi.Mux) chi.Router {
	return router{m}
}

func (r router) Mount(pattern string, h http.Handler) {
	if mux, ok := h.(*nchi.Mux); ok {
		r.Mux.Mout(translatePattern(pattern), mux)
	} else {
		// XXX do better
		r.Mux.Handle(translatePattern(pattern), h.ServeHTTP)
	}
}

func (r router) Middlewares() Middlewares { return nil } // XXX
func (r router) Match(rctx *chi.Context, method, path string) bool {
	return false
	// XXX
}
func (r router) Use(middlewares ...func(http.Handler) http.Handler) {
	r.Mux.Use(asAnySlice(middlewares)...)
}
func (r router) With(middlewares ...func(http.Handler) http.Handler) chi.Router {
	return AsRouter(r.Mux.With(asAnySlice(middlewares)...))
}
func (r router) Group(fn func(r Router)) Router {
	return AsRouter(r.Mux.Group(func(mux *nchi.Mux) {
		fn(AsRouter(mux))
	}))
}
func (r router) Route(pattern string, fn func(r Router)) Router {
	return AsRouter(r.Mux.Route(translatePattern(pattern), func(mux *nchi.Mux) {
		fn(AsRouter(mux))
	}))
}
func (r router) Method(method, pattern string, h http.Handler) {
	r.Mux.Method(method, translatePattern(pattern), h.ServeHTTP)
}
func (r router) MethodFunc(method, pattern string, h http.HandlerFunc) {
	r.Mux.Method(method, translatePattern(pattern), h)
}
func (r router) Connect(pattern string, h http.HandlerFunc) {
	r.Mux.Connect(translatePattern(pattern), h)
}
func (r router) Delete(pattern string, h http.HandlerFunc) {
	r.Mux.Delete(translatePattern(pattern), h)
}
func (r router) Get(pattern string, h http.HandlerFunc) {
	r.Mux.Get(translatePattern(pattern), h)
}
func (r router) Head(pattern string, h http.HandlerFunc) {
	r.Mux.Head(translatePattern(pattern), h)
}
func (r router) Options(pattern string, h http.HandlerFunc) {
	r.Mux.Options(translatePattern(pattern), h)
}
func (r router) Patch(pattern string, h http.HandlerFunc) {
	r.Mux.Patch(translatePattern(pattern), h)
}
func (r router) Post(pattern string, h http.HandlerFunc) {
	r.Mux.Post(translatePattern(pattern), h)
}
func (r router) Put(pattern string, h http.HandlerFunc) {
	r.Mux.Put(translatePattern(pattern), h)
}
func (r router) Trace(pattern string, h http.HandlerFunc) {
	r.Mux.Trace(translatePattern(pattern), h)
}
func (r router) Handle(pattern string, h http.Handler) {
	r.Mux.Handle(translatePattern(pattern), h.ServeHTTP)
}
func (r router) HandleFunc(pattern string, h http.HandlerFunc) {
	r.Mux.Handle(translatePattern(pattern), h.ServeHTTP)
}
func (r router) NotFound(h http.HandlerFunc) {
	r.Mux.NotFound(h)
}
func (r router) MethodNotAllowed(h http.HandlerFunc) {
	r.Mux.MethodNotAllowed(h)
}

func asAnySlice[T any](a []T) []any {
	n := make([]any, len(a))
	for i, v := range a {
		n[i] = v
	}
	return n
}

func translatePattern(p pattern) string {
	// XXX
	return p
}
