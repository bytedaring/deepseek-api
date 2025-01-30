package deepseek_api_test

import (
	"errors"
	"testing"

	deepseek_api "github.com/ZSLTChenXiYin/deepseek-api"
)

func TestDeepSeekChatRequest_Validation(t *testing.T) {
	tests := []struct {
		name     string
		request  *deepseek_api.DeepSeekChatRequest
		expected error
	}{
		{
			name: "Valid Request",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: nil,
		},
		{
			name: "Invalid Messages",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("messages must be at least one"),
		},
		{
			name: "Invalid Model",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("model must be set"),
		},
		{
			name: "Invalid FrequencyPenalty",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 3,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("frequency_penalty must be between -2 and 2"),
		},
		{
			name: "Invalid MaxTokens",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        10000,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("max_tokens must be between 1 and 8192"),
		},
		{
			name: "Invalid PresencePenalty",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  3,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("presence_penalty must be between -2 and 2"),
		},
		{
			name: "Invalid ResponseFormat",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: "invalid"},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("response_format.type must be text or json_object"),
		},
		{
			name: "Invalid Stop",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             make([]string, 17),
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("stop must be less than 16 lists"),
		},
		{
			name: "Invalid StreamOptions",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           true,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("stream_options must be set when stream is true"),
		},
		{
			name: "Invalid Temperature",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      3,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("temperature must be between 0 and 2"),
		},
		{
			name: "Invalid TopP",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             2,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
			},
			expected: errors.New("top_p must be between 0 and 1"),
		},
		{
			name: "Invalid ToolChoice",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       "invalid",
			},
			expected: errors.New("tool_choice must be one of none, auto, or required"),
		},
		{
			name: "Invalid Tools",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_REQUIRED,
				Tools:            []deepseek_api.Tool{},
			},
			expected: errors.New("tools must be defined when tool_choice is required"),
		},
		{
			name: "Invalid TopLogprobs",
			request: &deepseek_api.DeepSeekChatRequest{
				Messages:         []deepseek_api.DeepSeekMessage{&deepseek_api.BasicMessage{Role: "user", Content: "Hello"}},
				Model:            "deepseek-chat",
				FrequencyPenalty: 0,
				MaxTokens:        100,
				PresencePenalty:  0,
				ResponseFormat:   deepseek_api.ResponseFormat{Type: deepseek_api.RESPONSE_FORMAT_TEXT},
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
				ToolChoice:       deepseek_api.TOOL_CHOICE_NONE,
				Logprobs:         true,
				TopLogprobs: func() *int64 {
					ptr := new(int64)
					*ptr = 21
					return ptr
				}(),
			},
			expected: errors.New("top_logprobs must be between 0 and 20"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.request.DeepSeekRequest()
			if err != nil && test.expected == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && test.expected != nil {
				t.Errorf("Expected error: %v, but got none", test.expected)
			} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expected, err)
			}
		})
	}
}

func TestDeepSeekCompletionsRequest_Validation(t *testing.T) {
	tests := []struct {
		name     string
		request  *deepseek_api.DeepSeekCompletionsRequest
		expected error
	}{
		{
			name: "Valid Request",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: nil,
		},
		{
			name: "Invalid Model",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("model must be set"),
		},
		{
			name: "Invalid Prompt",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("prompt must be set"),
		},
		{
			name: "Invalid FrequencyPenalty",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 3,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("frequency_penalty must be between -2 and 2"),
		},
		{
			name: "Invalid Logprobs",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         21,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("logprobs must be between 0 and 20"),
		},
		{
			name: "Invalid MaxTokens",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        0,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("max_tokens must be greater than 0"),
		},
		{
			name: "Invalid PresencePenalty",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  3,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("presence_penalty must be between -2 and 2"),
		},
		{
			name: "Invalid Stop",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             make([]string, 17),
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("stop must be less than 16 lists"),
		},
		{
			name: "Invalid StreamOptions",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           true,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             0.9,
			},
			expected: errors.New("stream_options must be set when stream is true"),
		},
		{
			name: "Invalid Temperature",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      3,
				TopP:             0.9,
			},
			expected: errors.New("temperature must be between 0 and 2"),
		},
		{
			name: "Invalid TopP",
			request: &deepseek_api.DeepSeekCompletionsRequest{
				Model:            "deepseek-chat",
				Prompt:           "Hello",
				FrequencyPenalty: 0,
				Logprobs:         0,
				MaxTokens:        100,
				PresencePenalty:  0,
				Stop:             []string{},
				Stream:           false,
				StreamOptions:    nil,
				Temperature:      0.7,
				TopP:             2,
			},
			expected: errors.New("top_p must be between 0 and 1"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.request.DeepSeekRequest()
			if err != nil && test.expected == nil {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && test.expected != nil {
				t.Errorf("Expected error: %v, but got none", test.expected)
			} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error: %v, but got: %v", test.expected, err)
			}
		})
	}
}
