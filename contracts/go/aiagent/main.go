package main

import (
	"context"
	"log"

	"agent/contract"
	"agent/external"
	"agent/repository"
)

type Agent struct {
	contract *contract.AIOracleClient
	external *external.OpenAIClient
	repo     *repository.DB
}

func NewAgent() *Agent {
	return &Agent{
		contract: nil,
		external: nil,
		repo:     nil,
	}
}

func (a *Agent) Run(ctx context.Context) {
	prompts := a.contract.ListenSmartContractEvents(ctx)

	for {
		select {
		case err := <-ctx.Done():
			log.Println("context done: ", err)
			return

		case ps := <-prompts:
			// We don't care about batching. We can process each event individually.
			// If there are any errors, we will skip the event and try processing it on the next poll.
			for _, p := range ps {
				// External query is not a stateful operation, it can fail without side effects.
				resp, err := a.external.Do(ctx, external.SubmitPromptRequest{
					PromptID: p.PromptId,
					Prompt:   p.Prompt,
				})
				if err != nil {
					log.Println("error processing request: ", err)
					continue
				}

				// Committing the result is a stateful operation. If it fails, we do not want to
				// save the result to the repository. Contract should ensure that commit is atomic.
				r := resp.(external.SubmitPromptResponse)
				err = a.contract.SubmitAnswer(ctx, p.PromptId, r.Answer)
				if err != nil {
					log.Println("error committing result: ", err)
					continue
				}

				// It's not a big deal if we fail to save the response. At this point, the result
				// is already committed to the contract, so this event will not be processed again.
				err = a.repo.Save(r.MustToBytes())
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
