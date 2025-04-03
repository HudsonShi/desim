package main

type Resource struct {
	capacity int
	users    int
	queue    []*Event
}
