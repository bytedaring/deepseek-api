package deepseek_api_test

import (
	"errors"
	"testing"

	deepseek_api "github.com/ZSLTChenXiYin/deepseek-api"
)

func TestBasicMessage_DeepSeekMessage(t *testing.T) {
	tests := []struct {
		msg      deepseek_api.BasicMessage
		expected error
	}{
		{deepseek_api.BasicMessage{Role: ""}, errors.New("role cannot be empty")},
		{deepseek_api.BasicMessage{Role: "invalid"}, errors.New("role must be one of system, user, assistant, or tool")},
		{deepseek_api.BasicMessage{Role: deepseek_api.ROLE_SYSTEM, Content: ""}, errors.New("content cannot be empty")},
		{deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER, Content: "Hello"}, nil},
	}

	for _, test := range tests {
		err := test.msg.DeepSeekMessage()
		if err != test.expected {
			t.Errorf("DeepSeekMessage() = %v, want %v", err, test.expected)
		}
	}
}

func TestSystemMessage_DeepSeekMessage(t *testing.T) {
	tests := []struct {
		msg      deepseek_api.SystemMessage
		expected error
	}{
		{deepseek_api.SystemMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER}}, errors.New("role must be system")},
		{deepseek_api.SystemMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_SYSTEM, Content: ""}}, errors.New("content cannot be empty")},
		{deepseek_api.SystemMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_SYSTEM, Content: "System message"}}, nil},
	}

	for _, test := range tests {
		err := test.msg.DeepSeekMessage()
		if err != test.expected {
			t.Errorf("DeepSeekMessage() = %v, want %v", err, test.expected)
		}
	}
}

func TestUserMessage_DeepSeekMessage(t *testing.T) {
	tests := []struct {
		msg      deepseek_api.UserMessage
		expected error
	}{
		{deepseek_api.UserMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_SYSTEM}}, errors.New("role must be user")},
		{deepseek_api.UserMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER, Content: ""}}, errors.New("content cannot be empty")},
		{deepseek_api.UserMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER, Content: "User message"}}, nil},
	}

	for _, test := range tests {
		err := test.msg.DeepSeekMessage()
		if err != test.expected {
			t.Errorf("DeepSeekMessage() = %v, want %v", err, test.expected)
		}
	}
}

func TestAssistantMessage_DeepSeekMessage(t *testing.T) {
	tests := []struct {
		msg      deepseek_api.AssistantMessage
		expected error
	}{
		{deepseek_api.AssistantMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER}}, errors.New("role must be assistant")},
		{deepseek_api.AssistantMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_ASSISTANT, Content: ""}}, errors.New("content cannot be empty")},
		{deepseek_api.AssistantMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_ASSISTANT, Content: "Assistant message"}, ReasoningContent: "Reasoning", Prefix: false}, errors.New("prefix must be true if reasoning_context is not empty")},
		{deepseek_api.AssistantMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_ASSISTANT, Content: "Assistant message"}, ReasoningContent: "Reasoning", Prefix: true}, nil},
	}

	for _, test := range tests {
		err := test.msg.DeepSeekMessage()
		if err != test.expected {
			t.Errorf("DeepSeekMessage() = %v, want %v", err, test.expected)
		}
	}
}

func TestToolMessage_DeepSeekMessage(t *testing.T) {
	tests := []struct {
		msg      deepseek_api.ToolMessage
		expected error
	}{
		{deepseek_api.ToolMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_USER}}, errors.New("role must be tool")},
		{deepseek_api.ToolMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_TOOL, Content: ""}}, errors.New("content cannot be empty")},
		{deepseek_api.ToolMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_TOOL, Content: "Tool message"}, ToolCallId: ""}, errors.New("tool_call_id cannot be empty")},
		{deepseek_api.ToolMessage{BasicMessage: deepseek_api.BasicMessage{Role: deepseek_api.ROLE_TOOL, Content: "Tool message"}, ToolCallId: "123"}, nil},
	}

	for _, test := range tests {
		err := test.msg.DeepSeekMessage()
		if err != test.expected {
			t.Errorf("DeepSeekMessage() = %v, want %v", err, test.expected)
		}
	}
}
