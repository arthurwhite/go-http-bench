package bench

import "testing"

var (
	paramsRoutesColon = []string{
		"/item/:page",
		"/user/:item",
		"/users/:id/carriage",
		"/users/:id/car",
	}

	paramsRoutesCurlyBracket = []string{
		"/item/{page}",
		"/user/{item}",
		"/users/{id}/carriage",
		"/users/{id}/car",
	}

	paramsRequests = []string{
		"/users/42/car",
		"/page/notfound",
		"/item/about",
		"/user/files",
		"/users/42/carriage",
	}
)

func BenchmarkParamsChi(b *testing.B) {
	bench(b, paramsRequests, setChi(paramsRoutesColon))
}

func BenchmarkParamsFastRoute(b *testing.B) {
	bench(b, paramsRequests, setFastRoute(paramsRoutesColon))
}

func BenchmarkParamsGorillaMux(b *testing.B) {
	bench(b, staticRequests, setGorillaMux(paramsRoutesCurlyBracket))
}

func BenchmarkParamsGowwwRouter(b *testing.B) {
	bench(b, paramsRequests, setGowwwRouter(paramsRoutesColon))
}

func BenchmarkParamsHTTPRouter(b *testing.B) {
	bench(b, paramsRequests, setHTTPRouter(paramsRoutesColon))
}
