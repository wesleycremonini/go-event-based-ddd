package workforce

type Job struct {
	Service Service
	Input   any
}

type Service interface {
	Execute(input any)
}

type Queue interface {
	Enqueue(job Job)
}
