package deepseek_api

import "errors"

type DeepSeekMessage interface {
	DeepSeekMessage() error
	GetContent() string
	GetRole() string
}

const (
	ROLE_SYSTEM    = "system"
	ROLE_USER      = "user"
	ROLE_ASSISTANT = "assistant"
	ROLE_TOOL      = "tool"
)

type BasicMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

func (m *BasicMessage) DeepSeekMessage() error {
	if m.Role == "" {
		return errors.New("role cannot be empty")
	}

	switch m.Role {
	case ROLE_SYSTEM, ROLE_USER, ROLE_ASSISTANT, ROLE_TOOL:
		if m.Content == "" {
			return errors.New("content cannot be empty")
		}
	default:
		return errors.New("role must be one of system, user, assistant, or tool")
	}

	return nil
}

func (m *BasicMessage) GetContent() string {
	return m.Content
}

func (m *BasicMessage) GetRole() string {
	return m.Role
}

type SystemMessage struct {
	BasicMessage
	Name string `json:"name"`
}

func (m *SystemMessage) DeepSeekMessage() error {
	if m.Role != ROLE_SYSTEM {
		return errors.New("role must be system")
	}

	if m.Content == "" {
		return errors.New("content cannot be empty")
	}

	return nil
}

type UserMessage struct {
	BasicMessage
	Name string `json:"name"`
}

func (m *UserMessage) DeepSeekMessage() error {
	if m.Role != ROLE_USER {
		return errors.New("role must be user")
	}

	if m.Content == "" {
		return errors.New("content cannot be empty")
	}

	return nil
}

type AssistantMessage struct {
	BasicMessage
	Name             string `json:"name"`
	Prefix           bool   `json:"prefix"`
	ReasoningContent string `json:"reasoning_content"`
}

func (m *AssistantMessage) DeepSeekMessage() error {
	if m.Role != ROLE_ASSISTANT {
		return errors.New("role must be assistant")
	}

	if m.Content == "" {
		return errors.New("content cannot be empty")
	}

	if m.ReasoningContent != "" {
		if !m.Prefix {
			return errors.New("prefix must be true if reasoning_context is not empty")
		}
	}

	return nil
}

type ToolMessage struct {
	BasicMessage
	ToolCallId string `json:"tool_call_id"`
}

func (m *ToolMessage) DeepSeekMessage() error {
	if m.Role != ROLE_TOOL {
		return errors.New("role must be tool")
	}

	if m.Content == "" {
		return errors.New("content cannot be empty")
	}

	if m.ToolCallId == "" {
		return errors.New("tool_call_id cannot be empty")
	}

	return nil
}

type ToolCall struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type ResponseMessage struct {
	BasicMessage
	ReasoningContent string     `json:"reasoning_content"`
	ToolCalls        []ToolCall `json:"tool_calls"`
}
