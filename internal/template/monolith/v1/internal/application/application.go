package application

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Configure(cfg string) {

}

func (s *Server) StartDev(frontendDir string) {
	s.startSrv(frontendDir, 8080)
}
