package external

import (
	"context"
	"fmt"
	"log/slog"

	"agent/config"
	"github.com/go-resty/resty/v2"
)

type OpenAIClient struct {
	logger *slog.Logger
	http   *resty.Client
	config config.OpenAIConfig
}

// NewOpenAIClient creates and returns a new instance of OpenAIClient.
func NewOpenAIClient(logger *slog.Logger, config config.OpenAIConfig) *OpenAIClient {
	return &OpenAIClient{
		logger: logger,
		http: resty.New().
			SetBaseURL(config.BaseURL).
			SetAuthToken(config.APIKey).
			SetAuthScheme("Bearer").
			SetHeader("Content-Type", "application/json").
			SetHeader("OpenAI-Beta", "assistants=v2"),
		config: config,
	}
}

type SubmitPromptResponse struct {
	Answer      string
	MessageID   string
	ThreadID    string
	RunID       string
	AssistantID string
}

func (SubmitPromptResponse) IsResponse() {}

// SubmitPrompt sends a prompt to the OpenAI API.
func (c *OpenAIClient) SubmitPrompt(ctx context.Context, promptID uint64, prompt string) (SubmitPromptResponse, error) {
	run, err := c.CreateThreadAndRunMessage(ctx, "user", prompt, promptID)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("create run: %w", err)
	}

	_, err = c.PollRunResult(ctx, run.ThreadId, run.Id)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("poll run result: thread ID: %s: run ID: %s: %w", run.ThreadId, run.Id, err)
	}

	msgs, err := c.ListMessagesByRun(ctx, run.ThreadId, run.Id)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("list messages by run: thread ID: %s: run ID: %s: %w", run.ThreadId, run.Id, err)
	}

	if len(msgs.Data) != 1 {
		return SubmitPromptResponse{}, fmt.Errorf("expected 1 message from AI in a single run, got %d: thread ID: %s: run ID %s", len(msgs.Data), run.ThreadId, run.Id)
	}

	if len(msgs.Data[0].Content) != 1 {
		return SubmitPromptResponse{}, fmt.Errorf("expected 1 answer from AI in a single message, got %d: thread ID: %s: run ID %s: message ID %s", len(msgs.Data), run.ThreadId, run.Id, msgs.Data[0].Id)
	}

	return SubmitPromptResponse{
		Answer:      msgs.Data[0].Content[0].Text.Value,
		MessageID:   msgs.Data[0].Id,
		ThreadID:    msgs.Data[0].ThreadId,
		RunID:       msgs.Data[0].RunId,
		AssistantID: msgs.Data[0].AssistantId,
	}, nil
}
