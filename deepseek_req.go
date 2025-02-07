package deepseek_api

import "errors"

const (
	RESPONSE_FORMAT_TEXT        = "text"
	RESPONSE_FORMAT_JSON_OBJECT = "json_object"
)

type DeepSeekRequest interface {
	DeepSeekRequest() error
	GetStream() bool
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type StreamOption struct {
	IncludeUsage bool `json:"include_usage"`
}

type Tool struct {
	Type     string `json:"type"`
	Function struct {
		Description string         `json:"description"`
		Name        string         `json:"name"`
		Parameters  map[string]any `json:"parameters"`
	} `json:"function"`
}

const (
	MODEL_DEEPSEEK_CHAT     = "deepseek-chat"
	MODEL_DEEPSEEK_REASONER = "deepseek-reasoner"

	TOOL_CHOICE_NONE     = "none"
	TOOL_CHOICE_AUTO     = "auto"
	TOOL_CHOICE_REQUIRED = "required"
)

type DeepSeekChatRequest struct {
	Messages         []DeepSeekMessage `json:"messages"`
	Model            string            `json:"model"`
	FrequencyPenalty float64           `json:"frequency_penalty"`
	MaxTokens        int64             `json:"max_tokens"`
	PresencePenalty  float64           `json:"presence_penalty"`
	ResponseFormat   ResponseFormat    `json:"response_format"`
	Stop             []string          `json:"stop"`
	Stream           bool              `json:"stream"`
	StreamOptions    *StreamOption     `json:"stream_options"`
	Temperature      float64           `json:"temperature"`
	TopP             float64           `json:"top_p"`
	Tools            []Tool            `json:"tools"`
	ToolChoice       string            `json:"tool_choice"`
	Logprobs         bool              `json:"logprobs"`
	TopLogprobs      *int64            `json:"top_logprobs"`
}

func NewDeepSeekChatRequest(messages []DeepSeekMessage, model string) *DeepSeekChatRequest {
	return &DeepSeekChatRequest{
		Messages:         messages,
		Model:            model,
		FrequencyPenalty: 0,
		MaxTokens:        4096,
		PresencePenalty:  0,
		ResponseFormat: ResponseFormat{
			Type: RESPONSE_FORMAT_TEXT,
		},
		Stop:          nil,
		Stream:        false,
		StreamOptions: nil,
		Temperature:   1,
		TopP:          1,
		Tools:         nil,
		ToolChoice:    TOOL_CHOICE_NONE,
		Logprobs:      false,
		TopLogprobs:   nil,
	}
}

func (dsr *DeepSeekChatRequest) DeepSeekRequest() error {
	var err error

	if len(dsr.Messages) < 1 {
		err = errors.Join(err, errors.New("messages must be at least one"))
	}

	if dsr.Model == "" {
		err = errors.Join(err, errors.New("model must be set"))
	}

	if dsr.FrequencyPenalty < -2 || dsr.FrequencyPenalty > 2 {
		err = errors.Join(err, errors.New("frequency_penalty must be between -2 and 2"))
	}

	if dsr.MaxTokens < 1 || dsr.MaxTokens > 8192 {
		err = errors.Join(err, errors.New("max_tokens must be between 1 and 8192"))
	}

	if dsr.PresencePenalty < -2 || dsr.PresencePenalty > 2 {
		err = errors.Join(err, errors.New("presence_penalty must be between -2 and 2"))
	}

	if dsr.ResponseFormat.Type != RESPONSE_FORMAT_TEXT && dsr.ResponseFormat.Type != RESPONSE_FORMAT_JSON_OBJECT {
		err = errors.Join(err, errors.New("response_format.type must be text or json_object"))
	}

	if len(dsr.Stop) > 16 {
		err = errors.Join(err, errors.New("stop must be less than 16 lists"))
	}

	if dsr.Stream {
		if dsr.StreamOptions == nil {
			err = errors.Join(err, errors.New("stream_options must be set when stream is true"))
		}
	} else {
		if dsr.StreamOptions != nil {
			err = errors.Join(err, errors.New("stream_options must be nil when stream is false"))
		}
	}

	if dsr.Temperature < 0 || dsr.Temperature > 2 {
		err = errors.Join(err, errors.New("temperature must be between 0 and 2"))
	}

	if dsr.TopP < 0 || dsr.TopP > 1 {
		err = errors.Join(err, errors.New("top_p must be between 0 and 1"))
	}

	switch dsr.ToolChoice {
	case TOOL_CHOICE_NONE, TOOL_CHOICE_AUTO:
	case TOOL_CHOICE_REQUIRED:
		if len(dsr.Tools) < 1 {
			err = errors.Join(err, errors.New("tools must be defined when tool_choice is required"))
		}
	default:
		err = errors.Join(err, errors.New("tool_choice must be one of none, auto, or required"))
	}

	if dsr.TopLogprobs != nil {
		if !dsr.Logprobs {
			err = errors.Join(err, errors.New("top_logprobs must be defined when logprobs is true"))
		}
		if *dsr.TopLogprobs < 0 || *dsr.TopLogprobs > 20 {
			err = errors.Join(err, errors.New("top_logprobs must be between 0 and 20"))
		}
	}

	return err
}

func (dsr *DeepSeekChatRequest) GetStream() bool {
	return dsr.Stream
}

type DeepSeekCompletionsRequest struct {
	Model            string        `json:"model"`
	Prompt           string        `json:"prompt"`
	Echo             bool          `json:"echo"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	Logprobs         int64         `json:"logprobs"`
	MaxTokens        int64         `json:"max_tokens"`
	PresencePenalty  float64       `json:"presence_penalty"`
	Stop             []string      `json:"stop"`
	Stream           bool          `json:"stream"`
	StreamOptions    *StreamOption `json:"stream_options"`
	Suffix           *string       `json:"suffix"`
	Temperature      float64       `json:"temperature"`
	TopP             float64       `json:"top_p"`
}

func NewDeepSeekCompletionsRequest(model string, prompt string) *DeepSeekCompletionsRequest {
	return &DeepSeekCompletionsRequest{
		Model:            model,
		Prompt:           prompt,
		Echo:             false,
		FrequencyPenalty: 0,
		Logprobs:         0,
		MaxTokens:        1024,
		PresencePenalty:  0,
		Stop:             nil,
		Stream:           false,
		StreamOptions:    nil,
		Suffix:           nil,
		Temperature:      1,
		TopP:             1,
	}
}

func (dsr *DeepSeekCompletionsRequest) DeepSeekRequest() error {
	var err error

	if dsr.Model == "" {
		err = errors.Join(err, errors.New("model must be set"))
	}

	if dsr.Prompt == "" {
		err = errors.Join(err, errors.New("prompt must be set"))
	}

	if dsr.FrequencyPenalty < -2 || dsr.FrequencyPenalty > 2 {
		err = errors.Join(err, errors.New("frequency_penalty must be between -2 and 2"))
	}

	if dsr.Logprobs < 0 || dsr.Logprobs > 20 {
		err = errors.Join(err, errors.New("logprobs must be between 0 and 20"))
	}

	if dsr.MaxTokens < 1 {
		err = errors.Join(err, errors.New("max_tokens must be greater than 0"))
	}

	if dsr.PresencePenalty < -2 || dsr.PresencePenalty > 2 {
		err = errors.Join(err, errors.New("presence_penalty must be between -2 and 2"))
	}

	if len(dsr.Stop) > 16 {
		err = errors.Join(err, errors.New("stop must be less than 16 lists"))
	}

	if dsr.Stream {
		if dsr.StreamOptions == nil {
			err = errors.Join(err, errors.New("stream_options must be set when stream is true"))
		}
	} else {
		if dsr.StreamOptions != nil {
			err = errors.Join(err, errors.New("stream_options must be nil when stream is false"))
		}
	}

	if dsr.Temperature < 0 || dsr.Temperature > 2 {
		err = errors.Join(err, errors.New("temperature must be between 0 and 2"))
	}

	if dsr.TopP < 0 || dsr.TopP > 1 {
		err = errors.Join(err, errors.New("top_p must be between 0 and 1"))
	}

	return err
}

func (dsr *DeepSeekCompletionsRequest) GetStream() bool {
	return dsr.Stream
}
