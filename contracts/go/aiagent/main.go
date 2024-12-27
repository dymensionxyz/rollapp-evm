package main

import (
	"context"
	"log"
	"time"

	"agent/contract"
	"agent/external"
	"agent/repository"
)

type Agent struct {
	contract contract.Client
	external external.Client
	repo     repository.Repository
}

func NewAgent() *Agent {
	return &Agent{
		contract: nil,
		external: nil,
		repo:     nil,
	}
}

func (a *Agent) Run(ctx context.Context) {
	const pollInterval = time.Minute
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case err := <-ctx.Done():
			log.Println("context done: ", err)
			return

		case <-ticker.C:
			events, err := a.contract.GetEvents()
			if err != nil {
				// We will try again on the next poll.
				log.Println("error getting events: ", err)
				continue
			}

			// We don't care about batching. We can process each event individually.
			// If there are any errors, we will skip the event and try processing it on the next poll.
			for _, event := range events {
				// External query is not a stateful operation, it can fail without side effects.
				resp, err := a.external.Do(external.Request(event))
				if err != nil {
					log.Println("error processing event: ", err)
					continue
				}

				// Committing the result is a stateful operation. If it fails, we do not want to
				// save the result to the repository. Contract should ensure that commit is atomic.
				err = a.contract.CommitResult(contract.Result(resp))
				if err != nil {
					log.Println("error committing result: ", err)
					continue
				}

				// It's not a big deal if we fail to save the response. At this point, the result
				// is already committed to the contract, so this event will not be processed again.
				err = a.repo.Save([]byte(resp))
				if err != nil {
					log.Println("error saving response: ", err)
					continue
				}
			}
		}
	}
}

func main() {
	_ = NewAgent()
}
