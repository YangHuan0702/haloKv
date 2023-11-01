package server

import (
	"haloKv/src/pb"
	"haloKv/src/raft"
	"sync"
	"time"
)

type Server struct {
	data map[string]string

	LogBuffer []*pb.Log

	Lock    sync.Mutex
	Conde   *sync.Cond
	LogChan chan pb.Log

	rf *raft.Raft
}

func (serv *Server) WriteToMap() {
	ticker := time.NewTicker(2 * time.Millisecond)
	for {
		select {
		case l, _ := <-serv.LogChan:
			if len(serv.LogBuffer) >= 10 {
				serv.toWrite()
			} else {
				serv.LogBuffer = append(serv.LogBuffer, &l)
			}
		case <-ticker.C:
			serv.toWrite()
		}
	}
}

func (serv *Server) Put(key, value string) {
	serv.data[key] = value
}

func (serv *Server) Get(key string) string {
	return serv.data[key]
}

func (serv *Server) toWrite() {
	serv.Lock.Lock()
	defer serv.Lock.Unlock()

	for _, l := range serv.LogBuffer {
		serv.data[l.GetKey()] = l.GetValue()
	}

	serv.LogBuffer = serv.LogBuffer[0:0]
}
