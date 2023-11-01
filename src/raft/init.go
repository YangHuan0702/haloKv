package raft

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"haloKv/src/config"
	"haloKv/src/pb"
	"log"
	"net"
	"strings"
	"sync"
)

func GetRaftInstance(mapChan *chan pb.Log) *Raft {
	c := config.ReadConfig()
	raftConfig := c.Raft
	servers := strings.Split(raftConfig.Servers, ",")

	raft := Raft{}
	raft.CurrentTerm = -1
	raft.VotedFor = -1
	raft.CommitIndex = -1
	raft.LastApplied = 0
	raft.lock = sync.Mutex{}
	raft.cv = sync.NewCond(&raft.lock)
	raft.WriteStatus = false
	raft.mapChan = mapChan

	raft.NextIndex = make([]int, len(servers))
	raft.serverMap = make(map[int32]pb.RaftRpcClient)
	raft.MatchIndex = make([]int, len(servers))
	raft.ElectTime = raftConfig.ElectTime
	raft.HeartbeatTimeOut = raftConfig.HeartbeatTimeOut
	raft.State = CANDIDATE
	raft.serverSize = len(servers)

	for i, server := range servers {
		if i == raftConfig.Me {
			lis, err := net.Listen("tcp", server)
			if err != nil {
				log.Printf("RPC: listener localhost panic: {%s}", err)
				return nil
			}
			// 创建gRPC服务器
			s := grpc.NewServer()
			// 在grpc注册服务
			pb.RegisterRaftRpcServer(s, &raft)
			raft.me = i
			if s.Serve(lis) != nil {
				log.Printf("RPC: Server Starting Panic...")
				return nil
			}
		} else {
			flag.Parsed()
			conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("%s connection err : %s", server, err)
				return nil
			}
			client := pb.NewRaftRpcClient(conn)
			raft.serverMap[int32(i)] = client
		}
	}
	return &raft
}
