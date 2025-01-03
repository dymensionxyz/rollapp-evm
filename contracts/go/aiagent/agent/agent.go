package agent

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"agent/contract"
	"agent/external"
	"agent/repository"
	"github.com/go-chi/chi/v5"
)

type Agent struct {
	logger *slog.Logger
	server *http.Server

	contract *contract.AIOracleClient
	external *external.OpenAIClient
	repo     *repository.DB
}

func NewAgent(
	logger *slog.Logger,
	serverAddr string,
	contract *contract.AIOracleClient,
	external *external.OpenAIClient,
	repo *repository.DB,
) *Agent {
	a := &Agent{
		logger:   logger,
		server:   nil, // is set further
		contract: contract,
		external: external,
		repo:     repo,
	}

	r := chi.NewRouter()
	r.Get("/get-answer/{promptID}", a.getAnswerHandler)
	a.server = &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	return a
}

func (a *Agent) Run(ctx context.Context) {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.With("error", err).Error("HTTP server error")
		}
	}()

	prompts := a.contract.ListenSmartContractEvents(ctx)

	for ps := range prompts {
		a.logger.Info("Got unprocessed prompts", "count", len(ps), "prompts", ps)
		// We don't care about batching. We can process each event individually.
		// If there are any errors, we will skip the event and try processing it on the next poll.
		var wg sync.WaitGroup
		wg.Add(len(ps))
		for _, p := range ps {
			// OpenAI calls may be time-consuming, so we process them concurrently.
			go func() {
				defer wg.Done()
				a.logger.Info("Processing prompt", "promptId", p.PromptId, "prompt", p.Prompt)
				// External query is not a stateful operation, it can fail without side effects.
				r, err := a.external.SubmitPrompt(ctx, p.PromptId, p.Prompt)
				if err != nil {
					a.logger.With("error", err, "promptId", p.PromptId).
						Error("Error on submitting prompt to AI")
					return
				}

				a.logger.Info("Got prompt answer", "promptId", p.PromptId, "answer", r.Answer)
				// Committing the result is a stateful operation. If it fails, we do not want to
				// save the result to the repository. Contract should ensure that commit is atomic.
				err = a.contract.SubmitAnswer(ctx, p.PromptId, r.Answer)
				if err != nil {
					a.logger.With("error", err, "promptId", p.PromptId).
						Info("Error on submitting answer to contract")
					return
				}

				a.logger.Info("Answer committed to contract", "promptId", p.PromptId, "answer", r.Answer)
				// It's not a big deal if we fail to save the response. At this point, the result
				// is already committed to the contract, so this event will not be processed again.
				err = a.repo.Save(p.PromptId, repository.Answer{
					Answer:      r.Answer,
					MessageID:   r.MessageID,
					ThreadID:    r.ThreadID,
					RunID:       r.RunID,
					AssistantID: r.AssistantID,
				})
				if err != nil {
					a.logger.With("error", err, "promptId", p.PromptId).
						Info("Error on saving response to DB")
					return
				}
			}()
		}
		wg.Wait()
	}
}

func (a *Agent) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return a.server.Shutdown(ctx)
}

func (a *Agent) getAnswerHandler(w http.ResponseWriter, r *http.Request) {
	// Extract promptID from the URL path
	promptID := chi.URLParam(r, "promptID")
	if promptID == "" {
		http.Error(w, "promptID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(promptID, 10, 64)
	if err != nil {
		http.Error(w, "promptID must be an integer number", http.StatusBadRequest)
		return
	}

	answer, err := a.repo.Get(id)
	if err != nil {
		http.Error(w, "Error retrieving answer from DB", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(answer); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
