package que

import (
	"context"
	"fmt"
)

func Send[TRequest Request](queBuilder *queBuilder, ctx context.Context, request TRequest) error {
	requestHandler, ok := queBuilder.items[request.GetSignature()]
	if !ok {
		return fmt.Errorf("could not find request handler with signature: %s", request.GetSignature())
	}

	rh, ok := requestHandler.(func(context.Context, TRequest) error)
	if !ok {
		return fmt.Errorf("could not find request handler")
	}

	return rh(ctx, request)
}

type Que interface {
	Handle(ctx context.Context, request Request) (*Response, error)
}
