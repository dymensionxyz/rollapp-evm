package external

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	defaultThreadID    = "thread_Ke1ZN1NaAPD5duORxiwwb5fi"
	defaultAssistantID = "asst_qKNNnqSrtSDLS8g9bSTe1TR8"

	defaultLimit = "20"
	defaultOrder = "desc"
)

func (c *OpenAIClient) CreateMessage(ctx context.Context, role, content string, promptID uint64) (ThreadMessage, error) {
	var result ThreadMessage

	resp, err := c.http.R().
		SetContext(ctx).
		SetPathParam("thread_id", defaultThreadID).
		SetBody(CreateMessageReq{
			Role:    role,
			Content: content,
			Metadata: map[string]string{
				"prompt_id": fmt.Sprintf("%d", promptID),
			},
		}).
		SetResult(&result).
		SetError(&ErrorResp{}).
		Post("/v1/threads/{thread_id}/messages")

	if err != nil {
		return ThreadMessage{}, fmt.Errorf("failed to create message: %w", err)
	}
	if resp.IsError() {
		return ThreadMessage{}, fmt.Errorf("failed to create message: %s", resp.Error().(*ErrorResp).Error.Message)
	}

	return result, nil
}

func (c *OpenAIClient) RetrieveMessage(ctx context.Context, messageID string) (ThreadMessage, error) {
	var result ThreadMessage

	resp, err := c.http.R().
		SetContext(ctx).
		SetPathParam("thread_id", defaultThreadID).
		SetPathParam("message_id", messageID).
		SetResult(&result).
		SetError(&ErrorResp{}).
		Get("/v1/threads/{thread_id}/messages/{message_id}")

	if err != nil {
		return ThreadMessage{}, fmt.Errorf("failed to retrieve message: %w", err)
	}
	if resp.IsError() {
		return ThreadMessage{}, fmt.Errorf("failed to retrieve message: %s", resp.Error().(*ErrorResp).Error.Message)
	}

	return result, nil
}

func (c *OpenAIClient) ListMessagesByRun(ctx context.Context, threadID, runID string) (ThreadMessageList, error) {
	return c.listMessages(ctx, threadID, map[string]string{
		"run_id": runID,
		"limit":  defaultLimit,
		"order":  defaultOrder,
	})
}

func (c *OpenAIClient) ListMessages(ctx context.Context, threadID string) (ThreadMessageList, error) {
	return c.listMessages(ctx, threadID, map[string]string{
		"limit": defaultLimit,
		"order": defaultOrder,
	})
}

func (c *OpenAIClient) listMessages(ctx context.Context, threadID string, queryParams map[string]string) (ThreadMessageList, error) {
	var result ThreadMessageList

	resp, err := c.http.R().
		SetContext(ctx).
		SetPathParam("thread_id", threadID).
		SetQueryParams(queryParams).
		SetResult(&result).
		SetError(&ErrorResp{}).
		Get("/v1/threads/{thread_id}/messages")

	if err != nil {
		return ThreadMessageList{}, fmt.Errorf("failed to list messages: %w", err)
	}
	if resp.IsError() {
		return ThreadMessageList{}, fmt.Errorf("failed to list messages: %s", resp.Error().(*ErrorResp).Error.Message)
	}

	return result, nil
}

// CreateThreadAndRunMessage creates a thread and runs it with the message constructed from the given role and content.
func (c *OpenAIClient) CreateThreadAndRunMessage(ctx context.Context, role, content string, promptID uint64) (ThreadRun, error) {
	var result ThreadRun

	resp, err := c.http.R().
		SetContext(ctx).
		SetBody(CreateRunReq{
			AssistantId: defaultAssistantID,
			Thread: CreateThreadReq{
				Messages: []CreateMessageReq{{
					Role:    role,
					Content: content,
				}},
				Metadata: map[string]string{
					"prompt_id": fmt.Sprintf("%d", promptID),
				},
			},
		}).
		SetResult(&result).
		SetError(&ErrorResp{}).
		Post("/v1/threads/runs")

	if err != nil {
		return ThreadRun{}, fmt.Errorf("failed to create run: %w", err)
	}
	if resp.IsError() {
		return ThreadRun{}, fmt.Errorf("failed to create run: %s", resp.Error().(*ErrorResp).Error.Message)
	}

	return result, nil
}

func (c *OpenAIClient) RetrieveRun(ctx context.Context, threadID, runID string) (ThreadRun, error) {
	return retrieveRun(ctx, c.http, threadID, runID)
}

func (c *OpenAIClient) PollRunResult(ctx context.Context, threadID, runID string) (ThreadRun, error) {
	// Do `Clone` to avoid modifying the original client.
	pollingClient := c.http.Clone().
		SetRetryCount(c.config.PollRetryCount).
		SetRetryWaitTime(c.config.PollRetryWaitTime).
		SetRetryMaxWaitTime(c.config.PollRetryMaxWaitTime).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			c.logger.Debug("Polling run status...", "threadId", threadID, "runId", runID)
			return r.Result().(*ThreadRun).Status != "completed"
		})
	return retrieveRun(ctx, pollingClient, threadID, runID)
}

func retrieveRun(ctx context.Context, client *resty.Client, threadID, runID string) (ThreadRun, error) {
	var result ThreadRun

	resp, err := client.R().
		SetContext(ctx).
		SetPathParam("thread_id", threadID).
		SetPathParam("run_id", runID).
		SetResult(&result).
		SetError(&ErrorResp{}).
		Get("/v1/threads/{thread_id}/runs/{run_id}")

	if err != nil {
		return ThreadRun{}, fmt.Errorf("failed to retrieve run: %w", err)
	}
	if resp.IsError() {
		return ThreadRun{}, fmt.Errorf("failed to retrieve run: %s", resp.Error().(*ErrorResp).Error.Message)
	}

	return result, nil
}

type CreateMessageReq struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	// Optional. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	Metadata map[string]string `json:"metadata"`
}

// ThreadMessage represents a message in a thread. Example:
//
//	{
//	  "id": "msg_abc123",
//	  "object": "thread.message",
//	  "created_at": 1698983503,
//	  "thread_id": "thread_abc123",
//	  "role": "assistant",
//	  "content": [
//	    {
//	      "type": "text",
//	      "text": {
//	        "value": "Hi! How can I help you today?",
//	        "annotations": []
//	      }
//	    }
//	  ],
//	  "assistant_id": "asst_abc123",
//	  "run_id": "run_abc123",
//	  "attachments": [],
//	  "metadata": {}
//	}
type ThreadMessage struct {
	Id          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int               `json:"created_at"`
	ThreadId    string            `json:"thread_id"`
	Role        string            `json:"role"`
	Content     []MessageContent  `json:"content"`
	AssistantId string            `json:"assistant_id"`
	RunId       string            `json:"run_id"`
	Attachments []interface{}     `json:"attachments"`
	Metadata    map[string]string `json:"metadata"`
}

// MessageContent represents the content of a message. Example:
//
//	{
//	  "type": "text",
//	  "text": {
//	    "value": "Hi! How can I help you today?",
//	    "annotations": []
//	  }
//	}
type MessageContent struct {
	Type string `json:"type"`
	Text struct {
		Value       string        `json:"value"`
		Annotations []interface{} `json:"annotations"`
	} `json:"text"`
}

type ThreadMessageList struct {
	Object  string          `json:"object"` // Must be "list"
	Data    []ThreadMessage `json:"data"`
	FirstId string          `json:"first_id"`
	LastId  string          `json:"last_id"`
	HasMore bool            `json:"has_more"`
}

type CreateRunReq struct {
	AssistantId string          `json:"assistant_id"`
	Thread      CreateThreadReq `json:"thread"`
	// Optional. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	Metadata map[string]string `json:"metadata"`
}

type CreateThreadReq struct {
	Messages []CreateMessageReq `json:"messages"`
	// Optional. Keys can be a maximum of 64 characters long and values can be a maximum of 512 characters long.
	Metadata map[string]string `json:"metadata"`
}

// ThreadRun represents a run in a thread. Example:
//
//	{
//	  "id": "run_abc123",
//	  "object": "thread.run",
//	  "created_at": 1698107661,
//	  "assistant_id": "asst_abc123",
//	  "thread_id": "thread_abc123",
//	  "status": "completed",
//	  "started_at": 1699073476,
//	  "expires_at": null,
//	  "cancelled_at": null,
//	  "failed_at": null,
//	  "completed_at": 1699073498,
//	  "last_error": null,
//	  "model": "gpt-4o",
//	  "instructions": null,
//	  "tools": [{"type": "file_search"}, {"type": "code_interpreter"}],
//	  "metadata": {},
//	  "incomplete_details": null,
//	  "usage": {
//	    "prompt_tokens": 123,
//	    "completion_tokens": 456,
//	    "total_tokens": 579
//	  },
//	  "temperature": 1.0,
//	  "top_p": 1.0,
//	  "max_prompt_tokens": 1000,
//	  "max_completion_tokens": 1000,
//	  "truncation_strategy": {
//	    "type": "auto",
//	    "last_messages": null
//	  },
//	  "response_format": "auto",
//	  "tool_choice": "auto",
//	  "parallel_tool_calls": true
//	}
type ThreadRun struct {
	Id                  string             `json:"id"`
	Object              string             `json:"object"`
	CreatedAt           int                `json:"created_at"`
	AssistantId         string             `json:"assistant_id"`
	ThreadId            string             `json:"thread_id"`
	Status              string             `json:"status"`
	StartedAt           int                `json:"started_at"`
	ExpiresAt           interface{}        `json:"expires_at"`
	CancelledAt         interface{}        `json:"cancelled_at"`
	FailedAt            interface{}        `json:"failed_at"`
	CompletedAt         int                `json:"completed_at"`
	LastError           interface{}        `json:"last_error"`
	Model               string             `json:"model"`
	Instructions        interface{}        `json:"instructions"`
	Metadata            map[string]string  `json:"metadata"`
	IncompleteDetails   interface{}        `json:"incomplete_details"`
	Temperature         float64            `json:"temperature"`
	TopP                float64            `json:"top_p"`
	MaxPromptTokens     int                `json:"max_prompt_tokens"`
	MaxCompletionTokens int                `json:"max_completion_tokens"`
	TruncationStrategy  TruncationStrategy `json:"truncation_strategy"`
	ResponseFormat      string             `json:"response_format"`
	ToolChoice          string             `json:"tool_choice"`
	ParallelToolCalls   bool               `json:"parallel_tool_calls"`
}

type TruncationStrategy struct {
	Type         string `json:"type"`
	LastMessages *int   `json:"last_messages"`
}

type ErrorResp struct {
	Error struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    string      `json:"code"`
	} `json:"error"`
}
