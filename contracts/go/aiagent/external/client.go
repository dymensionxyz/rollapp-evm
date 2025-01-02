package external

import "context"

type Request interface {
	IsRequest()
}

type Response interface {
	IsResponse()
}

type Client interface {
	Do(context.Context, Request) (Response, error)
}
