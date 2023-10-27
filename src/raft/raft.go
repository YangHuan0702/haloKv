package raft

import (
	"context"
	"haloKv/src/pb"
	rc "haloKv/src/recover"
	"log"
	"sync"
	"time"
)

const (
	LEADER    = 1
	FOLLOWER  = 2
	CANDIDATE = 3
)

type RpcCall struct {
}

type Raft struct {
	CurrentTerm int32
	VotedFor    int
	log         []rc.RaftLog
	CommitIndex int32
	LastApplied int

	NextIndex  []int
	MatchIndex []int

	State int

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

func (raft *Raft) VoteRequest(ctx context.Context, request *pb.RequestVote) (*pb.ResponseVote, error) {
	resp := &pb.ResponseVote{}
	*resp.VoteGranted = false
	*resp.Term = raft.CurrentTerm

	if raft.CurrentTerm > *request.Term {
		return resp, nil
	}

	if *request.Term > raft.CurrentTerm {
		raft.follower(request.GetTerm(), int(request.GetCandidateId()))
	}
	*resp.Term = raft.CurrentTerm

	if (raft.VotedFor == -1 || raft.VotedFor == int(request.GetCandidateId())) &&
		(len(raft.log) == 0 || request.GetPrevLogTerm() > raft.log[len(raft.log)-1].Term ||
			(request.GetPrevLogTerm() == raft.log[len(raft.log)-1].Term && int(request.GetPrevLogIndex()) >= len(raft.log)-1)) {
		*resp.VoteGranted = true
		*resp.Term = raft.CurrentTerm
		return resp, nil
	}

	return resp, nil
}

func (raft *Raft) follower(term int32, candidateId int) {
	raft.CurrentTerm = term
	raft.VotedFor = candidateId
	raft.State = FOLLOWER
}

func (raft *Raft) AppendEntries(ctx context.Context, in *pb.RequestAppendEntries) (*pb.ResponseAppendEntries, error) {

	return nil, nil
}
