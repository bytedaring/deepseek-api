package deepseek_api_test

import (
	"testing"

	deepseek_api "github.com/ZSLTChenXiYin/deepseek-api"
)

func TestDeepSeekErrorResponse_DeepSeekResponse_ReturnsError(t *testing.T) {
	errResponse := &deepseek_api.DeepSeekErrorResponse{
		Error: struct {
			Message string         `json:"message"`
			Type    string         `json:"type"`
			Param   map[string]any `json:"param"`
			Code    string         `json:"code"`
		}{
			Message: "test error message",
			Type:    "test type",
			Param:   map[string]any{"key": "value"},
			Code:    "test code",
		},
	}

	err := errResponse.DeepSeekResponse()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if err.Error() != "test error message" {
		t.Errorf("Expected error message 'test error message', but got '%s'", err.Error())
	}
}

func TestDeepSeekCompletionResponse_DeepSeekResponse_ReturnsNil(t *testing.T) {
	completionResponse := &deepseek_api.DeepSeekChatResponse{
		Id:      "test-id",
		Choices: []deepseek_api.ChatChoice{},
		Created: 1234567890,
		Model:   "test-model",
	}

	err := completionResponse.DeepSeekResponse()

	if err != nil {
		t.Errorf("Expected nil, but got error: %v", err)
	}
}
