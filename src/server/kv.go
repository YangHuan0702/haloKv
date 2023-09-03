package server

type Server struct {
	data map[string]string
}

func (serv *Server) Put(key, value string) {
	serv.data[key] = value
}

func (serv *Server) Get(key string) string {
	return serv.data[key]
}
