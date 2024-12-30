package external

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var _ Client = new(OpenAIClient)

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

func (c *OpenAIClient) Do(r Request) (Response, error) {
	switch req := r.(type) {
	case SubmitPromptRequest:
		return c.SubmitPrompt(req)
	default:
		return nil, fmt.Errorf("unknown request type: %T", r)
	}
}

type SubmitPromptRequest struct {
	Prompt   string
	PromptID uint64
}

func (SubmitPromptRequest) IsRequest() {}

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
func (c *OpenAIClient) SubmitPrompt(req SubmitPromptRequest) (SubmitPromptResponse, error) {
	c.threadMu.Lock()
	defer c.threadMu.Unlock()

	promptMsg, err := c.CreateMessage("user", req.Prompt, req.PromptID)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("create message: %w", err)
	}

	run, err := c.CreateRun(req.PromptID)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("create run: %w", err)
	}

	_, err = c.PollRunResult(run.Id)
	if err != nil {
		return SubmitPromptResponse{}, fmt.Errorf("poll run result: run ID: %s: %w", run.Id, err)
	}

	msgs, err := c.ListMessagesByRun(run.Id)
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
