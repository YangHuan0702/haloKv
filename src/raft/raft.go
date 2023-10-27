package raft

import (
	"golang.org/x/net/context"
	"haloKv/src/pb"
)

type Raft struct {
	CurrentTerm int32
	VotedFor    int
	//log	log[]
	CommitIndex int32
	LastApplied int

	NextIndex  []int
	MatchIndex []int32

	pb.UnimplementedRaftRpcServer
}

func (raft *Raft) sendVoteRequest(ctx context.Context, vote *pb.RequestVote) (*pb.ResponseVote, error) {

	return &pb.ResponseVote{}, nil
}

func (raft *Raft) sendAppendEntries(ctx context.Context, request *pb.RequestAppendEntries) (*pb.ResponseAppendEntries, error) {

	return &pb.ResponseAppendEntries{}, nil
}
