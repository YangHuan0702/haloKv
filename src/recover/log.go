package recover

type RaftLog struct {
	Term int32
	Data interface{}
}

type Log struct {
	Key   string
	value string
}
