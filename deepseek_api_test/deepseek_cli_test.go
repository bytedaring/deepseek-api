package deepseek_api_test

import (
	"testing"

	deepseek_api "github.com/ZSLTChenXiYin/deepseek-api"
)

func TestDeepSeekClient_Validation(t *testing.T) {
	api_key := "YOUR_API_KEY"

	basic_message := &deepseek_api.BasicMessage{
		Role:    deepseek_api.ROLE_USER,
		Content: "你好。",
	}

	messages := []deepseek_api.DeepSeekMessage{basic_message}

	chat_request := deepseek_api.NewDeepSeekChatRequest(
		messages,
		deepseek_api.MODEL_DEEPSEEK_CHAT,
	)

	deepseek_client := deepseek_api.DefaultDeepSeekClient(api_key)

	chat_response, err := deepseek_client.Chat(chat_request)
	if err != nil {
		t.Fatalf("Chat error: %v", err)
	}

	t.Log("Chat request:")
	for _, message := range messages {
		t.Logf("%s: %s", message.GetRole(), message.GetContent())
	}

	t.Log("Chat response:")
	for _, choice := range chat_response.Choices {
		t.Logf("%s: %s", choice.Message.Role, choice.Message.Content)
	}
}
