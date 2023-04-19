package nchichi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muir/nchi"
)

// Shim increases compatibility with chi by adding a
// *chi.Context to the context carried by *http.Request.
//
// This chi.Context is only partially filled out. It will
// have RoutePatterns out.
func Shim(r *http.Request, params nchi.Params) *http.Request {
	paramKeys := make([]string, len(nchi.Params))
	parmaValues := make([]string, len(nchi.Params))
	for i, p := range params {
		paramKeys[i] = p.Key
		paramValues[i] = p.Value
	}
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams = nchi.RouteParams{
		Keys:   paramKeys,
		Values: paramValues,
	}
	return r.WithContext(context.WithValue(chi.RouteCtxKey, chiCtx))
}
