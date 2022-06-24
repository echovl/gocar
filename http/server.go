package http

import (
	"github.com/echovl/gocar/gocar"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r   *gin.Engine
	stg gocar.CarStorage
}

func NewServer(stg gocar.CarStorage) *Server {
	return &Server{
		r:   gin.Default(),
		stg: stg,
	}
}

func (s *Server) Listen(addr string) error {
	s.initRoutes()
	return s.r.Run(addr)
}
