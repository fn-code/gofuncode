package server

import (
	"app/middleware"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type Server struct {
	Hostname string
	HTTPPort string
}

// Run is running or start server
func Run(mx *http.ServeMux) {
	s := Server{"", "8080"}
	httpAddr := httpAddress(s)
	sck, err := GenerateSCK(32)
	if err != nil {
		log.Println(err)
	}
	CSRF := csrf.Protect(
		sck,
		csrf.Secure(false),
		csrf.ErrorHandler(middleware.CSRFError()),
	)
	srv := http.Server{
		Addr:    httpAddr,
		Handler: CSRF(mx),
	}
	log.Println("Server Running On Port : " + s.HTTPPort)
	srv.ListenAndServe()
}

func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%s", s.HTTPPort)
}

//GenerateSCK is generate 32byte secret key
func GenerateSCK(s int) ([]byte, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
