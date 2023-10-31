package raft

import (
	"context"
	"haloKv/src/pb"
	"log"
	"math/rand"
	"sort"
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
	log         []*pb.RaftLog
	CommitIndex int
	LastApplied int

	NextIndex  []int
	MatchIndex []int

	State int

	actionTime int64

	serverSize int

	mapChan *chan pb.Log
	me      int

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

func (raft *Raft) VoteRequest(_ context.Context, request *pb.RequestVote) (*pb.ResponseVote, error) {
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
		(len(raft.log) == 0 || request.GetPrevLogTerm() > *raft.log[len(raft.log)-1].Term ||
			(request.GetPrevLogTerm() == *raft.log[len(raft.log)-1].Term && int(request.GetPrevLogIndex()) >= len(raft.log)-1)) {
		*resp.VoteGranted = true
		*resp.Term = raft.CurrentTerm
		raft.actionTime = time.Now().Unix()
		return resp, nil
	}
	return resp, nil
}

func (raft *Raft) follower(term int32, candidateId int) {
	raft.CurrentTerm = term
	raft.VotedFor = candidateId
	raft.State = FOLLOWER
}

func (raft *Raft) AppendEntries(_ context.Context, request *pb.RequestAppendEntries) (*pb.ResponseAppendEntries, error) {
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

	if len(raft.log) == 0 || int(request.GetPrevLogIndex()) > len(raft.log)-1 || *raft.log[int(request.GetPrevLogIndex())].Term != request.GetPrevLogTerm() {
		return &resp, nil
	} else {
		rlog := &pb.RaftLog{}
		*rlog.Term = request.GetTerm()
		*rlog.LogEntries = *request.GetLogs()[0].GetLogEntries()
		raft.log = append(raft.log[:request.GetPrevLogIndex()+1], rlog)
	}

	if int(request.GetLeaderCommit()) > raft.CommitIndex {
		raft.CommitIndex = Min(len(raft.log)-1, int(request.GetLeaderCommit()))
		raft.startWriteStatus()
	}

	raft.actionTime = time.Now().Unix()
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
			*raft.mapChan <- *raft.log[i].GetLogEntries()
		}
		raft.LastApplied = raft.CommitIndex
		raft.WriteStatus = false
		raft.lock.Unlock()
	}
}

const randomTimeOut = 150

func (raft *Raft) electLeader() {
	sleepTime := time.Duration(int64(raft.ElectTime) + rand.Int63n(randomTimeOut))
	startTime := time.Now().Unix()
	for {
		time.Sleep(sleepTime * time.Millisecond)
		raft.lock.Lock()

		lastTime := startTime
		startTime := time.Now().Unix()
		predictTime := raft.actionTime + int64(raft.ElectTime) + rand.Int63n(randomTimeOut)
		diffTime := predictTime - startTime
		if raft.actionTime < lastTime || diffTime <= 0 {
			if raft.State != LEADER {
				// start election
				raft.kickoffLeaderElection()
			}
			sleepTime = time.Duration(int64(raft.ElectTime) + rand.Int63n(randomTimeOut))
		} else {
			sleepTime = time.Duration(diffTime)
		}
		raft.lock.Unlock()
	}
}

func (raft *Raft) kickoffLeaderElection() {
	raft.candidate()
	cond := sync.NewCond(&raft.lock)
	count, electCount := 1, 1

	prevIndex := len(raft.log) - 1
	var prevTerm int32 = -1
	if len(raft.log) > 0 {
		prevTerm = *raft.log[prevIndex].Term
	}

	for i := 0; i < raft.serverSize; i++ {
		if i == raft.me {
			continue
		}

		go func(term int32, targetServer int, prevIndex int, prevTerm int32, candidateId int) {
			request := pb.RequestVote{}
			*request.Term = term
			*request.CandidateId = int32(candidateId)
			*request.PrevLogIndex = int32(prevIndex)
			*request.PrevLogTerm = prevTerm
			*request.For = int32(targetServer)

			raft.lock.Lock()
			defer raft.lock.Unlock()

			resp := raft.sendVoteRequest(&request)

			count++
			if resp.GetVoteGranted() && resp.GetTerm() == raft.CurrentTerm {
				electCount++
			} else if resp.GetTerm() > raft.CurrentTerm {
				raft.follower(resp.GetTerm(), targetServer)
			}
			cond.Broadcast()
		}(raft.CurrentTerm, i, prevIndex, prevTerm, raft.me)
	}

	go func() {
		raft.lock.Lock()
		defer raft.lock.Unlock()

		for raft.State == CANDIDATE && count < raft.serverSize && electCount <= raft.serverSize/2 {
			cond.Wait()
		}

		if raft.State == CANDIDATE && electCount > raft.serverSize/2 {
			raft.leader()
			go raft.leaderSendHeartbeat()
		}
	}()

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (raft *Raft) leaderSendHeartbeat() {
	for {
		raft.lock.Lock()
		if raft.State != LEADER {
			raft.lock.Unlock()
			return
		}

		for i := 0; i < raft.serverSize; i++ {
			if i == raft.me {
				continue
			}
			prevLogIndex := Max(Min(raft.NextIndex[i]-1, len(raft.log)-1), -1)
			var prevLogTerm int32 = -1
			if prevLogIndex >= 0 {
				prevLogTerm = *raft.log[prevLogIndex].Term
			}
			logs := make([]*pb.RaftLog, 0)
			if len(raft.log) > prevLogIndex+1 {
				logs = append(logs, raft.log[prevLogIndex+1])
			}

			go func(i int, term int32, leaderId int, commitIndex int, prevLogIndex int, prevLogTerm int32, rlogs []*pb.RaftLog) {
				reqeust := pb.RequestAppendEntries{}
				*reqeust.Term = term
				*reqeust.LeaderId = int32(leaderId)
				*reqeust.PrevLogIndex = int32(prevLogIndex)
				*reqeust.PrevLogTerm = prevLogTerm
				*reqeust.LeaderCommit = int32(commitIndex)
				*reqeust.For = int32(i)
				reqeust.Logs = rlogs

				raft.lock.Lock()
				defer raft.lock.Unlock()
				resp := raft.sendAppendEntries(&reqeust)
				if resp.GetTerm() > raft.CurrentTerm {
					raft.follower(*resp.Term, i)
					return
				}

				if term != raft.CurrentTerm || *resp.Term != raft.CurrentTerm {
					return
				}

				if len(rlogs) > 0 {
					if !*resp.Success && raft.NextIndex[i] > 0 {
						if *resp.FirstIndex >= 0 && *raft.log[*resp.FirstIndex].Term == *resp.FirstTerm {
							raft.NextIndex[i] = int(*resp.FirstIndex + 1)
						} else {
							raft.NextIndex[i] = int(*resp.FirstIndex - 1)
						}
					}
					if *resp.Success {
						if raft.NextIndex[i] <= 0 {
							raft.NextIndex[i] = len(rlogs)
							raft.MatchIndex[i] = len(rlogs) - 1
						} else {
							raft.NextIndex[i] = prevLogIndex + len(rlogs) + 1
							raft.MatchIndex[i] = raft.NextIndex[i] - 1
						}
					}

					newMatchs := make([]int, raft.serverSize)
					copy(newMatchs, raft.MatchIndex)
					sort.Ints(newMatchs)
					for target := (raft.serverSize - 1) / 2; target > 0 && newMatchs[target] > raft.CommitIndex; target-- {
						// 如果leader提交了一个commitIndex的日志然后宕机了，后来的leader在这个位置上拥有更到的term，就会出现问题
						if newMatchs[target] <= len(raft.log)-1 && *raft.log[newMatchs[target]].Term == raft.CurrentTerm {
							raft.CommitIndex = newMatchs[target]
							raft.startWriteStatus()
							break
						}
					}
				}

			}(i, raft.CurrentTerm, raft.me, raft.CommitIndex, prevLogIndex, prevLogTerm, logs)
		}
		raft.lock.Unlock()
		time.Sleep(time.Duration(raft.HeartbeatTimeOut) * time.Millisecond)
	}
}

func (raft *Raft) leader() {
	raft.State = LEADER
	for i := 0; i < raft.serverSize; i++ {
		if i != raft.me {
			raft.MatchIndex[i] = -1
		}
		raft.NextIndex[i] = len(raft.log) - 1
	}
	raft.actionTime = time.Now().Unix()
}

func (raft *Raft) candidate() {
	raft.State = CANDIDATE
	raft.CurrentTerm += 1
	raft.VotedFor = -1
	// 这里设置时间的意义是保障选举发生的频率
	raft.actionTime = time.Now().Unix()
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (raft *Raft) StartRaftServer(mapChan *chan pb.Log) {
	raft.mapChan = mapChan

	go raft.electLeader()

	go raft.writeStatus()
}
