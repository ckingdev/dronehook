package dronehook

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Server struct {
	Port int
	Path string
	Out  chan Payload
}

func NewServer(port int, path string) *Server {
	return &Server{
		Port: port,
		Path: path,
		Out:  make(chan Payload),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != "POST" {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != s.Path {
		http.Error(w, "404 Not found- expected "+s.Path, http.StatusNotFound)
		return
	}

	respBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "500 Internal Error - Could not read from request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	go s.processPayload(respBody)
}

func (s *Server) processPayload(raw []byte) {
	p, err := makePayload(raw)
	if err != nil {
		panic(err)
		//		log.Printf("Error making payload: %s", err)
		return
	}
	s.Out <- *p
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", s.Port), s)
}

func (s *Server) GoListenAndServe() {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
