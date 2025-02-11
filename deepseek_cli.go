package deepseek_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DEFAULT_PROTOCOL = "https"

	DEFAULT_HOST = "api.deepseek.com"

	DEFAULT_CHAT_PATH        = "/chat/completions"
	DEFAULT_COMPLETIONS_PATH = "/beta/completions"
	DEFAULT_MODELS_PATH      = "/models"
	DEFAULT_BALANCE_PATH     = "/user/balance"

	DEFAULT_TIMEOUT = 30 * time.Second
)

type DeepSeekClient struct {
	protocol string
	host     string

	api_key string

	http_client *http.Client
}

type DeepSeekClientOptions func(*DeepSeekClient)

func WithDeepSeekClientCommunication(protocol string, host string) DeepSeekClientOptions {
	return func(dsc *DeepSeekClient) {
		dsc.protocol = protocol
		dsc.host = host
	}
}

func WithDeepSeekClientApi(api_key string) DeepSeekClientOptions {
	return func(dsc *DeepSeekClient) {
		dsc.api_key = api_key
	}
}

func WithDeepSeekClientHttpClient(http_client *http.Client) DeepSeekClientOptions {
	return func(dsc *DeepSeekClient) {
		if http_client == nil {
			http_client = &http.Client{Timeout: DEFAULT_TIMEOUT}
		}
		dsc.http_client = http_client
	}
}

func NewDeepSeekClient(options ...DeepSeekClientOptions) *DeepSeekClient {
	dsc := &DeepSeekClient{}
	for _, option := range options {
		option(dsc)
	}

	if dsc.api_key == "" {
		return nil
	}

	if dsc.protocol == "" {
		dsc.protocol = DEFAULT_PROTOCOL
	}

	if dsc.host == "" {
		dsc.host = DEFAULT_HOST
	}

	return dsc
}

func DefaultDeepSeekClient(api_key string) *DeepSeekClient {
	client_communication := WithDeepSeekClientCommunication(DEFAULT_PROTOCOL, DEFAULT_HOST)

	client_api_key := WithDeepSeekClientApi(api_key)

	client_http_client := WithDeepSeekClientHttpClient(nil)

	return NewDeepSeekClient(client_communication, client_api_key, client_http_client)
}

func (dsc *DeepSeekClient) GetProtocol() string {
	return dsc.protocol
}

func (dsc *DeepSeekClient) SetProtocol(protocol string) *DeepSeekClient {
	dsc.protocol = protocol
	return dsc
}

func (dsc *DeepSeekClient) GetHost() string {
	return dsc.host
}

func (dsc *DeepSeekClient) SetHost(host string) *DeepSeekClient {
	dsc.host = host
	return dsc
}

func (dsc *DeepSeekClient) GetApi() string {
	return dsc.api_key
}

func (dsc *DeepSeekClient) SetApi(api_key string) *DeepSeekClient {
	dsc.api_key = api_key
	return dsc
}

func (dsc *DeepSeekClient) GetHttpClient() *http.Client {
	return dsc.http_client
}

func (dsc *DeepSeekClient) SetHttpClient(http_client *http.Client) *DeepSeekClient {
	dsc.http_client = http_client
	return dsc
}

func (dsc *DeepSeekClient) getUrl(path string) string {
	return dsc.protocol + "://" + dsc.host + path
}

func (dsc *DeepSeekClient) getHeader() http.Header {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer "+dsc.api_key)
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")
	return headers
}

func (dsc *DeepSeekClient) Do(method string, path string, ds_req DeepSeekRequest) (ds_resp DeepSeekResponse, err error) {
	var req *http.Request
	if ds_req == nil {
		req, err = http.NewRequest(method, dsc.getUrl(path), nil)
		if err != nil {
			return nil, err
		}
	} else {
		if ds_req.StreamModel() {
			return nil, fmt.Errorf("streaming is not supported")
		}

		ds_req_json, err := json.Marshal(ds_req)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, dsc.getUrl(path), bytes.NewBuffer(ds_req_json))
		if err != nil {
			return nil, err
		}
	}

	req.Header = dsc.getHeader()

	resp, err := dsc.http_client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dsu_resp := make(DeepSeekUniversalResponse)

	err = json.Unmarshal(resp_body, &dsu_resp)
	if err != nil {
		return nil, err
	}

	if dsu_resp.DeepSeekResponse() != nil {
		ds_resp = &DeepSeekErrorResponse{}
	} else if _, ok := dsu_resp["is_available"]; ok {
		ds_resp = &DeepSeekBalanceResponse{}
	} else if _, ok := dsu_resp["object"]; ok {
		switch dsu_resp["object"] {
		case OBJECT_CHAT_COMPLETION:
			ds_resp = &DeepSeekChatResponse{}
		case OBJECT_TEXT_COMPLETION:
			ds_resp = &DeepSeekCompletionsResponse{}
		case OBJECT_LIST:
			ds_resp = &DeepSeekModelsResponse{}
		default:
			return nil, fmt.Errorf("unknown object type: %s", dsu_resp["object"])
		}
	} else {
		return nil, fmt.Errorf("unknown response type")
	}

	err = json.Unmarshal(resp_body, ds_resp)
	if err != nil {
		return nil, err
	}

	return ds_resp, nil
}

type StreamDoEvent func(response *http.Response, args ...any) error

func (dsc *DeepSeekClient) StreamDo(method string, path string, ds_req DeepSeekRequest, event StreamDoEvent, args ...any) error {
	if !ds_req.StreamModel() {
		return fmt.Errorf("stream must be set to true")
	}

	ds_req_json, err := json.Marshal(ds_req)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, dsc.getUrl(path), bytes.NewBuffer(ds_req_json))
	if err != nil {
		return err
	}

	req.Header = dsc.getHeader()

	resp, err := dsc.http_client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	err = event(resp, args...)
	if err != nil {
		return err
	}

	return nil
}

func (dsc *DeepSeekClient) Chat(dsc_req *DeepSeekChatRequest) (dsc_resp *DeepSeekChatResponse, err error) {
	ds_resp, err := dsc.Do(http.MethodPost, DEFAULT_CHAT_PATH, dsc_req)
	if err != nil {
		return nil, err
	}

	_, ok := ds_resp.(*DeepSeekErrorResponse)
	if ok {
		return nil, ds_resp.DeepSeekResponse()
	}

	return ds_resp.(*DeepSeekChatResponse), nil
}

func (dsc *DeepSeekClient) Completions(dsc_req *DeepSeekCompletionsRequest) (dsc_resp *DeepSeekCompletionsResponse, err error) {
	ds_resp, err := dsc.Do(http.MethodPost, DEFAULT_COMPLETIONS_PATH, dsc_req)
	if err != nil {
		return nil, err
	}

	_, ok := ds_resp.(*DeepSeekErrorResponse)
	if ok {
		return nil, ds_resp.DeepSeekResponse()
	}

	return ds_resp.(*DeepSeekCompletionsResponse), nil
}

func (dsc *DeepSeekClient) Models() (dsm_resp *DeepSeekModelsResponse, err error) {
	ds_resp, err := dsc.Do(http.MethodGet, DEFAULT_MODELS_PATH, nil)
	if err != nil {
		return nil, err
	}

	_, ok := ds_resp.(*DeepSeekErrorResponse)
	if ok {
		return nil, ds_resp.DeepSeekResponse()
	}

	return ds_resp.(*DeepSeekModelsResponse), nil
}

func (dsc *DeepSeekClient) Balance() (dsb_resp *DeepSeekBalanceResponse, err error) {
	ds_resp, err := dsc.Do(http.MethodGet, DEFAULT_BALANCE_PATH, nil)
	if err != nil {
		return nil, err
	}

	_, ok := ds_resp.(*DeepSeekErrorResponse)
	if ok {
		return nil, ds_resp.DeepSeekResponse()
	}

	return ds_resp.(*DeepSeekBalanceResponse), nil
}
