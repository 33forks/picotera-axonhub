//go:build tinygo

package bedrock

import (
	"context"
	"fmt"
	"io"

	"github.com/looplj/axonhub/llm/httpclient"
)

func NewAWSEventStreamDecoder(context.Context, io.ReadCloser) httpclient.StreamDecoder {
	return &unavailableEventStreamDecoder{err: fmt.Errorf("bedrock event stream decoder is unavailable in tinygo wasm builds")}
}

type unavailableEventStreamDecoder struct {
	err error
}

func (d *unavailableEventStreamDecoder) Next() bool {
	return false
}

func (d *unavailableEventStreamDecoder) Current() *httpclient.StreamEvent {
	return nil
}

func (d *unavailableEventStreamDecoder) Err() error {
	return d.err
}

func (d *unavailableEventStreamDecoder) Close() error {
	return nil
}

var _ httpclient.StreamDecoder = (*unavailableEventStreamDecoder)(nil)
