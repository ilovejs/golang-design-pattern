package ch40

// Request take any as input
type Request interface{}

// Response take any as output
type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}
