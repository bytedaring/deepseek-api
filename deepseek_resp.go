package deepseek_api

import (
	"fmt"
)

type DeepSeekResponse interface {
	DeepSeekResponse() error
}

type DeepSeekUniversalResponse map[string]any

func (dsr DeepSeekUniversalResponse) DeepSeekResponse() error {
	if dsr["error"] != nil {
		return fmt.Errorf("deepseek error: %s", dsr["error"].(map[string]any)["message"].(string))
	}
	return nil
}

type DeepSeekErrorResponse struct {
	Error struct {
		Message string         `json:"message"`
		Type    string         `json:"type"`
		Param   map[string]any `json:"param"`
		Code    string         `json:"code"`
	} `json:"error"`
}

func (dsr *DeepSeekErrorResponse) DeepSeekResponse() error {
	return fmt.Errorf("deepseek error: %s", dsr.Error.Message)
}

type Content struct {
	Token       string  `json:"token"`
	Logprob     float64 `json:"logprob"`
	Bytes       []int64 `json:"bytes"`
	TopLogprobs []struct {
		Token   string  `json:"token"`
		Logprob float64 `json:"logprob"`
		Bytes   []int64 `json:"bytes"`
	} `json:"top_logprobs"`
}

type ChatChoice struct {
	FinishReason string          `json:"finish_reason"`
	Index        int64           `json:"index"`
	Message      ResponseMessage `json:"message"`
	Logprobs     *struct {
		Content []Content `json:"content"`
	} `json:"logprobs"`
}

type Usage struct {
	CompletionTokens        int64  `json:"completion_tokens"`
	PromptTokens            int64  `json:"prompt_tokens"`
	PromptCacheHitTokens    *int64 `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens   *int64 `json:"prompt_cache_miss_tokens"`
	TotalTokens             int64  `json:"total_tokens"`
	CompletionTokensDetails *struct {
		ReasoningTokens int64 `json:"reasoning_tokens"`
	} `json:"completion_tokens_details"`
}

const (
	OBJECT_CHAT_COMPLETION = "chat.completion"
	OBJECT_TEXT_COMPLETION = "text_completion"
	OBJECT_LIST            = "list"
)

type DeepSeekChatResponse struct {
	Id                string       `json:"id"`
	Choices           []ChatChoice `json:"choices"`
	Created           int64        `json:"created"`
	Model             string       `json:"model"`
	SystemFingerprint *string      `json:"system_fingerprint"`
	Object            string       `json:"object"`
	Usage             Usage        `json:"usage"`
}

func (dsr *DeepSeekChatResponse) DeepSeekResponse() error {
	return nil
}

type CompletionsChoice struct {
	FinishReason string `json:"finish_reason"`
	Index        int64  `json:"index"`
	Logprobs     *struct {
		TextOffset    []int64        `json:"text_offset"`
		TokenLogprobs []float64      `json:"token_logprobs"`
		Tokens        []string       `json:"tokens"`
		TopLogprobs   map[string]any `json:"top_logprobs"`
	} `json:"logprobs"`
	Text string `json:"text"`
}

type DeepSeekCompletionsResponse struct {
	Id                string              `json:"id"`
	Choices           []CompletionsChoice `json:"choices"`
	Created           int64               `json:"created"`
	Model             string              `json:"model"`
	SystemFingerprint *string             `json:"system_fingerprint"`
	Object            string              `json:"object"`
	Usage             Usage               `json:"usage"`
}

func (dsr *DeepSeekCompletionsResponse) DeepSeekResponse() error {
	return nil
}

type DeepSeekModelsResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Id      string `json:"id"`
		Object  string `json:"object"`
		OwnedBy string `json:"owned_by"`
	}
}

func (dsr *DeepSeekModelsResponse) DeepSeekResponse() error {
	return nil
}

type DeepSeekBalanceResponse struct {
	IsAvailable  bool `json:"is_available"`
	BalanceInfos []struct {
		Currency        string `json:"currency"`
		TotalBalance    string `json:"total_balance"`
		GrantedBalance  string `json:"granted_balance"`
		ToppedUpBalance string `json:"topped_up_balance"`
	}
}

func (dsr *DeepSeekBalanceResponse) DeepSeekResponse() error {
	return nil
}
