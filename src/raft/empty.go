package raft

// Raft /*
type Raft struct {
	CurrentTerm int32
	VotedFor    int
	//log	log[]
	CommitIndex int32
	LastApplied int

	NextIndex  []int
	MatchIndex []int32
}

/*
------------------------ request & response------------------------------
*/
type RequestVote struct {
	Term         int32
	CandidateId  int
	LastLogIndex int32
	LastLogTerm  int32
}

type ResponseVote struct {
	Term        int32
	VoteGranted bool
}

type RequestAppendEntries struct {
	Term         int32
	LeaderId     int
	PrevLogIndex int32
	PrevLogTerm  int32
	LeaderCommit int32
}

type ResponseAppendEntries struct {
	Term    int32
	Success bool
}
