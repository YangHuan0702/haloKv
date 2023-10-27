package raft

import (
	"context"
	"haloKv/src/pb"
	"log"
	"sync"
	"time"
)

type RpcCall struct {
}

type Raft struct {
	CurrentTerm int32
	VotedFor    int
	//log	log[]
	CommitIndex int32
	LastApplied int

	NextIndex  []int
	MatchIndex []int

	ElectTime        int
	HeartbeatTimeOut int

	serverMap map[int32]pb.RaftRpcClient

	lock sync.Mutex
	cv   *sync.Cond
	pb.UnimplementedRaftRpcServer
}

func (raft *Raft) sendVoteRequest(vote *pb.RequestVote) *pb.ResponseVote {
	server := raft.serverMap[*vote.For]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := server.VoteRequest(ctx, vote)
	if err != nil {
		log.Fatal("Request VoteRequest Fatal : {}", err)
	}
	return r
}

func (raft *Raft) sendAppendEntries(request *pb.RequestAppendEntries) *pb.ResponseAppendEntries {
	server := raft.serverMap[*request.For]
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	r, err := server.AppendEntries(ctx, request)
	if err != nil {
		log.Fatal("Request AppendEntries Fatal : {}", err)
	}
	return r
}

func (raft *Raft) VoteRequest(ctx context.Context, in *pb.RequestVote) (*pb.ResponseVote, error) {

	return nil, nil
}
func (raft *Raft) AppendEntries(ctx context.Context, in *pb.RequestAppendEntries) (*pb.ResponseAppendEntries, error) {

	return nil, nil
}
