package server

import (
	"haloKv/src/pb"
	"sync"
)

func GetServer() *Server {
	server := Server{}
	server.data = make(map[string]string)
	server.LogBuffer = make([]*pb.Log, 0, 10)
	server.Lock = sync.Mutex{}
	server.Conde = sync.NewCond(&server.Lock)
	server.LogChan = make(chan pb.Log)

	return &server
}
