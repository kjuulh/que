package accepttests

import (
	"context"
	"github.com/kjuulh/que"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestRequest struct {
}

func (t TestRequest) GetSignature() string {
	return "test_item"
}

var _ que.Request = TestRequest{}

type TestItemRequestHandler struct{}

func (t TestItemRequestHandler) Handle(ctx context.Context, item any) (*que.Response, error) {
	return &que.Response{}, nil
}

var _ que.RequestHandler = TestItemRequestHandler{}

func TestLib(t *testing.T) {
	q := que.
		New().
		Define(TestRequest{}, TestItemRequestHandler{}).
		Build()

	res, err := q.Handle(context.Background(), TestRequest{})

	require.NoError(t, err)
	require.NotNil(t, res)
}
