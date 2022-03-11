package server

import "github.com/gin-gonic/gin"

type Register func(engine *gin.Engine) error

type Server struct {
	engine *gin.Engine
	rs     []Register
}

func (s *Server) Run(addr ...string) error {
	for _, r := range s.rs {
		if err := r(s.engine); err != nil {
			return err
		}
	}

	return s.engine.Run(addr...)
}

func (s *Server) With(r Register) *Server {
	s.rs = append(s.rs, r)
	return s
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
		rs:     make([]Register, 0),
	}
}
