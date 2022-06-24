package http

func (s *Server) initRoutes() {
	s.r.POST("/cars", s.handleCreateCar)
	s.r.GET("/cars", s.handleSearchCars)
	s.r.GET("/cars/export", s.handleExportCars)
}
