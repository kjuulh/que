package que

import "context"

type Request interface {
	GetSignature() string
}

type Response interface{}

type RequestHandlerContract[TRequest Request] interface {
	Handle(ctx context.Context, item TRequest) error
	Request
}
