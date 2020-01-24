package main

import (
	"fmt"
	"gRPC_exampleStream/server/counter"
	"time"
)

var ticker *time.Ticker
var countStatus bool

func (s *server) timer() {
	timer := time.NewTicker(1 * time.Second)
	ticker = timer

	count := 0

	s.broadcast <- &counter.Count{Value: 0, Timestamp: time.Stamp}
	fmt.Printf("count: %d\n", count)

	for range timer.C {
		count += 1
		s.broadcast <- &counter.Count{Value: int32(count), Timestamp: time.Stamp}
		fmt.Printf("count: %d\n", count)
	}
}

func (s *server) startCounter() {
	if !countStatus {
		go s.timer()
	}
	countStatus = true
}

func (s *server) stopCounter() {
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}
	countStatus = false
}

func (s *server) resetCounter() {
	s.stopCounter()
	s.startCounter()
}