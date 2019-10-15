package apiserver

// APIServer ...
type APIServer struct{}

// New initialize APIServer
func New() *APIServer {
	return &APIServer{}
}

// Start start API server
func (s *APIServer) Start() error {
	return nil
}
