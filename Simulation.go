package main

type Simulation struct {
	scheduler *Scheduler
	now       float64
	running   bool
}

func (sim *Simulation) Run(until float64) {
	for sim.running && sim.scheduler.HasNext() {
		evt := sim.scheduler.Next()
		sim.now = evt.Time
		evt.Action()
		if sim.now >= until {
			break
		}
	}
}
