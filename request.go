package que

import "context"

type Request interface {
	GetSignature() string
}

type Response struct{}

type RequestResponse struct{}

type RequestHandler interface {
	Handle(ctx context.Context, item any) (*Response, error)
}
