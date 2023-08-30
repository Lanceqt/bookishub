package server

import (
	"net/http"
)

func NewFileServer(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}
