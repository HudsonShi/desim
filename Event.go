package main

type Event struct {
	Time   float64
	Action func()
}
