package accepttests

import (
	"context"
	"github.com/kjuulh/que"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestRequest struct {
	SomeField string
}

func (t TestRequest) GetSignature() string {
	return "test_item"
}

var _ que.Request = &TestRequest{}

type TestItemRequestHandler[TRequest que.Request] struct{}

func (t TestItemRequestHandler[TRequest]) Handle(ctx context.Context, item TRequest) error {
	return nil
}

func NewTestRequestHandler() *TestItemRequestHandler[TestRequest] {
	return &TestItemRequestHandler[TestRequest]{}
}

func (t TestItemRequestHandler[TRequest]) GetSignature() string {
	return "test_item"
}

func TestLib(t *testing.T) {
	q := que.New()
	que.Define(q, NewTestRequestHandler())

	err := que.Send(q, context.Background(), TestRequest{})

	require.NoError(t, err)
}
