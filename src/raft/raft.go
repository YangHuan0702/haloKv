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
	CommitIndex int
	LastApplied int

	NextIndex  []int
	MatchIndex []int

	State int

	ElectTime        int
	HeartbeatTimeOut int
	WriteStatus      bool
	StopServer       bool

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
	raft.lock.Lock()
	defer raft.lock.Unlock()

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

func (raft *Raft) AppendEntries(ctx context.Context, request *pb.RequestAppendEntries) (*pb.ResponseAppendEntries, error) {
	raft.lock.Lock()
	defer raft.lock.Unlock()

	resp := pb.ResponseAppendEntries{}
	*resp.Success = false
	*resp.Term = raft.CurrentTerm

	if request.GetTerm() < raft.CurrentTerm {
		raft.VotedFor = -1
		return &resp, nil
	}

	if request.GetTerm() > raft.CurrentTerm {
		raft.follower(request.GetTerm(), int(request.GetLeaderId()))
	}

	if len(raft.log) == 0 || int(request.GetPrevLogIndex()) > len(raft.log)-1 || raft.log[int(request.GetPrevLogIndex())].Term != request.GetPrevLogTerm() {
		return &resp, nil
	} else {
		raft.log = append(raft.log[:request.GetPrevLogIndex()+1], rc.RaftLog{Term: request.GetTerm(), Data: request.GetData()})
	}

	if int(request.GetLeaderCommit()) > raft.CommitIndex {
		raft.CommitIndex = Min(len(raft.log)-1, int(request.GetLeaderCommit()))
		// TODO write to Map & Log
		raft.startWriteStatus()
	}

	*resp.Success = true
	return &resp, nil
}

func (raft *Raft) startWriteStatus() {
	raft.WriteStatus = true
	raft.cv.Broadcast()
}

func (raft *Raft) writeStatus() {
	for {
		raft.lock.Lock()
		if !raft.WriteStatus {
			raft.cv.Wait()
		}
		for i := raft.LastApplied; i <= raft.CommitIndex; i++ {
			//TODO Write To Map
		}
		raft.LastApplied = raft.CommitIndex
		raft.WriteStatus = false
		raft.lock.Unlock()
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
