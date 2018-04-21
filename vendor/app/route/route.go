package route

import (
	"app/controller"
	"app/middleware"
	"net/http"
)

const fileDIR = "./assets"

func LoadHTTP() *http.ServeMux {
	return router()
}

func router() *http.ServeMux {
	mx := http.NewServeMux()

	fs := http.FileServer(http.Dir(fileDIR))
	mx.Handle("/static/", http.StripPrefix("/static", fs))
	mx.HandleFunc("/register", middleware.Chain(controller.Register, middleware.Logging()))
	mx.HandleFunc("/login", middleware.Chain(controller.Login, middleware.Logging()))
	mx.HandleFunc("/", middleware.Chain(controller.Home, middleware.Logging()))

	return mx
}
