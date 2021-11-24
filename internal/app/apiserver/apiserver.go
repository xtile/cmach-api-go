package apiserver

type APIServer struct{}

func New(s *Config) *APIServer {
	return &APIServer{
		config: config
	}
}

func (s *APIServer) Start() error {
	return nil
}
