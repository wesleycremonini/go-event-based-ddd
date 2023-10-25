package workforce

type dispatcher struct {
	maxWorkers int
	workerPool chan chan Job
	jobQueue   chan Job
}

func NewDispatcher(maxWorkers int, queueSize int) *dispatcher {
	d := dispatcher{workerPool: make(chan chan Job, maxWorkers), maxWorkers: maxWorkers, jobQueue: make(chan Job, queueSize)}

	for i := 0; i < d.maxWorkers; i++ {
		newWorker(d.workerPool).start()
	}
	go d.dispatch()

	return &d
}

func (d *dispatcher) Enqueue(j Job) {
	go func(job Job) {
		d.jobQueue <- job
	}(j)
}

func (d *dispatcher) dispatch() {
	for j := range d.jobQueue {
		go func(job Job) {
			jobChannel := <-d.workerPool
			jobChannel <- job
		}(j)
	}
}
