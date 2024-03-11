package types

import "time"

type BurstStats struct {
	BurstNumber     int
	SuccessfulCount int
	ClientErrors    int
	ServerErrors    int
	TotalLatency    time.Duration
	MaxLatency      time.Duration
	MinLatency      time.Duration
	AverageLatency  time.Duration
}
