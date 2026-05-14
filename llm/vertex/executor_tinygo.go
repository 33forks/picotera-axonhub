//go:build tinygo

package vertex

import (
	"context"
	"fmt"

	"github.com/looplj/axonhub/llm/httpclient"
	"github.com/looplj/axonhub/llm/pipeline"
	"github.com/looplj/axonhub/llm/streams"
)

type Executor struct{}

func NewExecutorFromJSON(string, string, string) (*Executor, error) {
	return nil, fmt.Errorf("vertex executor is unavailable in tinygo wasm builds")
}

func (e *Executor) Do(context.Context, *httpclient.Request) (*httpclient.Response, error) {
	return nil, fmt.Errorf("vertex executor is unavailable in tinygo wasm builds")
}

func (e *Executor) DoStream(context.Context, *httpclient.Request) (streams.Stream[*httpclient.StreamEvent], error) {
	return nil, fmt.Errorf("vertex executor is unavailable in tinygo wasm builds")
}

var _ pipeline.Executor = (*Executor)(nil)
