package external

type Request interface {
	IsRequest()
}

type Response interface {
	IsResponse()
}

type Client interface {
	Do(Request) (Response, error)
}
