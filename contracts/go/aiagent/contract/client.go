package contract

type Event struct {
	Body []byte
}

type Result struct {
	Body []byte
}

//type AIOracleClient interface {
//	GetEvents() ([]Event, error)
//	CommitResult(Result) error
//}
