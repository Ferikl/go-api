package apiserver

// APIServer ...
type APIServer struct {
	config *Config
}

// New initialize APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start start API server
func (s *APIServer) Start() error {
	return nil
}
