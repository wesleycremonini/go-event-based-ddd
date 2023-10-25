package workforce

import (
	"go.uber.org/zap"
)

type worker struct {
	workerPool chan chan Job
	jobChannel chan Job
	quit       chan bool
}

func newWorker(workerPool chan chan Job) worker {
	return worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w worker) start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				zap.L().Error("Error: panic recovered in worker", zap.Any("error", r))
			}
		}()

		for {
			w.workerPool <- w.jobChannel
			job := <-w.jobChannel
			job.Service.Execute(job.Input)
		}
	}()
}
