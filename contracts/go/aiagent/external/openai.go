package external

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type OpenAIClient struct {
	http *resty.Client

	// Answer polling
	pollRetryCount       int
	pollRetryWaitTime    time.Duration
	pollRetryMaxWaitTime time.Duration

	threadMu sync.Mutex
}

// NewOpenAIClient creates and returns a new instance of OpenAIClient.
func NewOpenAIClient(
	apiKey, baseUrl string,
	pollRetryCount int,
	pollRetryWaitTime time.Duration,
	pollRetryMaxWaitTime time.Duration,
) *OpenAIClient {
	return &OpenAIClient{
		http: resty.New().
			SetBaseURL(baseUrl).
			SetAuthToken(apiKey).
			SetAuthScheme("Bearer").
			SetHeader("Content-Type", "application/json").
			SetHeader("OpenAI-Beta", "assistants=v2"),
		pollRetryCount:       pollRetryCount,
		pollRetryWaitTime:    pollRetryWaitTime,
		pollRetryMaxWaitTime: pollRetryMaxWaitTime,
		threadMu:             sync.Mutex{},
	}
}

type SubmitPromptResponse struct {
	Answer          string
	PromptMessageID string
	AnswerMessageID string
	ThreadID        string
	RunID           string
	AssistantID     string
}

func (SubmitPromptResponse) IsResponse() {}

// SubmitPrompt sends a prompt to the OpenAI API.
func (c *OpenAIClient) SubmitPrompt(ctx context.Context, promptID uint64, prompt string) (SubmitPromptResponse, error) {
	c.threadMu.Lock()
	defer c.threadMu.Unlock()

	promptMsg, err := c.CreateMessage(ctx, "user", prompt, promptID)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("create message: %w", err)
	}

	run, err := c.CreateRun(ctx, promptID)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("create run: %w", err)
	}

	_, err = c.PollRunResult(ctx, run.Id)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("poll run result: run ID: %s: %w", run.Id, err)
	}

	msgs, err := c.ListMessagesByRun(ctx, run.Id)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("list messages by run: run ID: %s: %w", run.Id, err)
	}

	if len(msgs.Data) != 1 {
		return SubmitPromptResponse{}, fmt.Errorf("expected 1 message from AI in a single run, got %d: run ID %s", len(msgs.Data), run.Id)
	}

	if len(msgs.Data[0].Content) != 1 {
		return SubmitPromptResponse{}, fmt.Errorf("expected 1 answer from AI in a single message, got %d: run ID %s: message ID %s", len(msgs.Data), run.Id, msgs.Data[0].Id)
	}

	return SubmitPromptResponse{
		Answer:          msgs.Data[0].Content[0].Text.Value,
		PromptMessageID: promptMsg.Id,
		AnswerMessageID: msgs.Data[0].Id,
		ThreadID:        msgs.Data[0].ThreadId,
		RunID:           msgs.Data[0].RunId,
		AssistantID:     msgs.Data[0].AssistantId,
	}, nil
}
