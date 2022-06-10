package que

import (
	"context"
	"fmt"
)

type que struct {
	*queBuilder
}

func (q *que) Handle(ctx context.Context, request Request) (*Response, error) {
	requestHandler, ok := q.items[request.GetSignature()]
	if !ok {
		return nil, fmt.Errorf("could not find request handler")
	}

	return requestHandler.Handle(ctx, request)
}

type Que interface {
	Handle(ctx context.Context, request Request) (*Response, error)
}
