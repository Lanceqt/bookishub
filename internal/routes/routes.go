package routes

import (
	"net/http"
)

type Route struct {
	URL string
	Handler http.HandlerFunc
}

var routes = []Route{
	{URL: "/", Handler: IndexHandler},
	{URL: "/signup", Handler: SignUpHandler},
	{URL: "/app", Handler: AppHandler},
}

func NewRouter(routes []Route) http.Handler {
    mux := http.NewServeMux()
    for _, route := range routes {
        mux.HandleFunc(route.URL, route.Handler)
    }
    return mux
}