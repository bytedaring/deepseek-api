package deepseek_api_test

import (
	"net/http"
	"testing"
	"time"

	deepseek_api "github.com/ZSLTChenXiYin/deepseek-api"
)

func TestDeepSeekClient_Validation(t *testing.T) {
	api_key := "sk-15d517eb49064e388fb67059a299cd0a"

	deepseek_client := deepseek_api.DefaultDeepSeekClient(api_key)

	t.Run("TestDeepSeekClient_Chat_Validation", func(t *testing.T) {
		basic_message := &deepseek_api.BasicMessage{
			Role:    deepseek_api.ROLE_USER,
			Content: "你好。",
		}

		messages := []deepseek_api.DeepSeekMessage{basic_message}

		chat_request := deepseek_api.NewDeepSeekChatRequest(
			messages,
			deepseek_api.MODEL_DEEPSEEK_CHAT,
		)

		chat_response, err := deepseek_client.Chat(chat_request)
		if err != nil {
			t.Errorf("Chat error: %v", err)
			return
		}

		t.Log("Chat:")
		for _, message := range messages {
			t.Logf("%s: %s", message.GetRole(), message.GetContent())
		}
		for _, choice := range chat_response.Choices {
			t.Logf("%s: %s", choice.Message.Role, choice.Message.Content)
		}
	})

	t.Run("TestDeepSeekClient_Completions_Validation", func(t *testing.T) {
		completion_http_client := &http.Client{Timeout: 120 * time.Second}

		chat_http_client := deepseek_client.GetHttpClient()

		deepseek_client.SetHttpClient(completion_http_client)

		prompt := "从前有一座山，山里有一个庙，庙里有一个老和尚，老和尚给小和尚讲故事："

		completions_request := deepseek_api.NewDeepSeekCompletionsRequest(deepseek_api.MODEL_DEEPSEEK_CHAT, prompt)

		completions_response, err := deepseek_client.Completions(completions_request)
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("Completions:")
		t.Log(prompt)
		for _, choice := range completions_response.Choices {
			t.Logf(choice.Text)
		}

		deepseek_client.SetHttpClient(chat_http_client)
	})

	t.Run("TestDeepSeekClient_Models_Validation", func(t *testing.T) {
		models_response, err := deepseek_client.Models()
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("Models:")
		for _, model := range models_response.Data {
			t.Logf("%s: %s - %s", model.Id, model.Object, model.OwnedBy)
		}
	})

	t.Run("TestDeepSeekClient_Balance_Validation", func(t *testing.T) {
		balance_response, err := deepseek_client.Balance()
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("Balance:")
		if balance_response.IsAvailable {
			if balance_response.BalanceInfos == nil {
				t.Log("BalanceInfos is nil")
				return
			}
			for _, balance := range balance_response.BalanceInfos {
				t.Logf("%s - %s - %s - %s", balance.Currency, balance.TotalBalance, balance.GrantedBalance, balance.ToppedUpBalance)
			}
		} else {
			t.Log("Not available")
		}
	})
}
