package contract

type Event struct {
	Body []byte
}

type Result struct {
	Body []byte
}

type Client interface {
	GetEvents() ([]Event, error)
	CommitResult(Result) error
}
