package main

type Scheduler struct {
	queue []*Event
	now   float64
}
