package bench

import (
	"github.com/DATA-DOG/fastroute"
	"github.com/gorilla/mux"
	"github.com/gowww/router"
	"github.com/julienschmidt/httprouter"
	"github.com/pressly/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// Different handler formats for routers
	handler           = http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	handlerFunc       = func(_ http.ResponseWriter, _ *http.Request) {}
	handlerHTTPRouter = func(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}
)

func bench(b *testing.B, requests []string, rt http.Handler) {
	b.ReportAllocs()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	for i := 0; i < b.N; i++ {
		for _, path := range requests {
			r.URL.Path = path
			rt.ServeHTTP(w, r)
		}
	}
}

func setChi(reqRoutes []string) http.Handler {
	rt := chi.NewRouter()
	for _, r := range reqRoutes {
		rt.Get(r, handlerFunc)
		rt.Post(r, handlerFunc)
		rt.Put(r, handlerFunc)
		rt.Patch(r, handlerFunc)
		rt.Delete(r, handlerFunc)
	}
	return rt
}

func setFastRoute(reqRoutes []string) http.Handler {
	var routes []fastroute.Router
	for _, r := range reqRoutes {
		routes = append(routes, fastroute.New(r, handlerFunc))
	}
	var tree = map[string]fastroute.Router{
		http.MethodGet:    fastroute.Chain(routes...),
		http.MethodPost:   fastroute.Chain(routes...),
		http.MethodPut:    fastroute.Chain(routes...),
		http.MethodPatch:  fastroute.Chain(routes...),
		http.MethodDelete: fastroute.Chain(routes...),
	}
	return fastroute.RouterFunc(func(r *http.Request) http.Handler {
		return tree[r.Method] // fastroute.Router is also http.Handler
	})
}

func setGorillaMux(reqRoutes []string) http.Handler {
	rt := mux.NewRouter()
	for _, r := range reqRoutes {
		rt.Handle(r, handler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)
	}
	return rt
}

func setGowwwRouter(reqRoutes []string) http.Handler {
	rt := router.New()
	for _, r := range reqRoutes {
		rt.Get(r, handler)
		rt.Post(r, handler)
		rt.Put(r, handler)
		rt.Patch(r, handler)
		rt.Delete(r, handler)
	}
	return rt
}

func setHTTPRouter(reqRoutes []string) http.Handler {
	rt := httprouter.New()
	for _, r := range reqRoutes {
		rt.GET(r, handlerHTTPRouter)
		rt.POST(r, handlerHTTPRouter)
		rt.PUT(r, handlerHTTPRouter)
		rt.PATCH(r, handlerHTTPRouter)
		rt.DELETE(r, handlerHTTPRouter)
	}
	return rt
}
