package routes

import (
	"net/http"
)

type Route struct {
	URL string
	Handler http.HandlerFunc
}

var routes = []Route{
	{URL: "/", Handler: indexHandler},
	{URL: "/signup", Handler: signUpHandler},
	{URL: "/app", Handler: appHandler},
}

func NewRouter(routes []Route) http.Handler {
    mux := http.NewServeMux()
    for _, route := range routes {
        mux.HandleFunc(route.URL, route.Handler)
    }
    return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
}

func appHandler(w http.ResponseWriter, r *http.Request) {
}