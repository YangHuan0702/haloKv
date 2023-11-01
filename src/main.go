package main

import _ "haloKv/src/config"
import "haloKv/src/raft"
import s "haloKv/src/server"

func main() {
	serv := s.GetServer()

	raftServer := raft.GetRaftInstance(&serv.LogChan)

	raftServer.StartRaftServer()

}
