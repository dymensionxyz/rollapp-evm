package external

type Request[ReqT any] struct {
	Body ReqT
}

type Response[RespT any] struct {
	Body RespT
}

type Client[ReqT, RespT any] interface {
	Do(Request[ReqT]) (Response[RespT], error)
}
