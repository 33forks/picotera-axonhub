//go:build tinygo

package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/looplj/axonhub/llm/streams"
)

// HttpClient is intentionally inert in TinyGo/WASI builds. PicoTera's
// llmbridge wasm module only uses axonhub for in-memory format conversion.
type HttpClient struct{}

type ClientOption func(*clientOptions)

type clientOptions struct{}

func WithInsecureSkipVerify(bool) ClientOption {
	return func(*clientOptions) {}
}

func NewHttpClientWithProxy(*ProxyConfig, ...ClientOption) *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) WithProxy(*ProxyConfig) *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) GetNativeClient() *http.Client {
	return nil
}

func NewHttpClient(...ClientOption) *HttpClient {
	return &HttpClient{}
}

func NewHttpClientWithClient(*http.Client) *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) Do(context.Context, *Request) (*Response, error) {
	return nil, fmt.Errorf("httpclient: network requests are unavailable in tinygo wasm builds")
}

func (hc *HttpClient) DoStream(context.Context, *Request) (streams.Stream[*StreamEvent], error) {
	return nil, fmt.Errorf("httpclient: network requests are unavailable in tinygo wasm builds")
}

func BuildHttpRequest(ctx context.Context, request *Request) (*http.Request, error) {
	var body io.Reader
	if len(request.Body) > 0 {
		body = bytes.NewReader(request.Body)
	}

	httpReq, err := http.NewRequestWithContext(ctx, request.Method, request.URL, body)
	if err != nil {
		return nil, err
	}

	httpReq.Header = request.Headers
	if httpReq.Header == nil {
		httpReq.Header = make(http.Header)
	}
	if request.ContentType != "" {
		httpReq.Header.Set("Content-Type", request.ContentType)
	}
	if request.Auth != nil {
		if err := applyAuth(httpReq.Header, request.Auth); err != nil {
			return nil, fmt.Errorf("failed to apply authentication: %w", err)
		}
	}
	if len(request.Query) > 0 {
		if httpReq.URL.RawQuery != "" {
			httpReq.URL.RawQuery += "&"
		}
		httpReq.URL.RawQuery += request.Query.Encode()
	}
	return httpReq, nil
}

func (hc *HttpClient) BuildHttpRequest(ctx context.Context, request *Request) (*http.Request, error) {
	return BuildHttpRequest(ctx, request)
}

func applyAuth(headers http.Header, auth *AuthConfig) error {
	switch auth.Type {
	case "bearer":
		if auth.APIKey == "" {
			return fmt.Errorf("bearer token is required")
		}
		headers.Set("Authorization", "Bearer "+auth.APIKey)
	case "api_key":
		if auth.HeaderKey == "" {
			return fmt.Errorf("header key is required")
		}
		headers.Set(auth.HeaderKey, auth.APIKey)
	default:
		return fmt.Errorf("unsupported auth type: %s", auth.Type)
	}
	return nil
}

func (hc *HttpClient) extractHeaders(headers http.Header) map[string]string {
	result := make(map[string]string)
	for key, values := range headers {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}
