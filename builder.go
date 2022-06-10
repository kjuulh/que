package que

type queBuilder struct {
	items map[string]RequestHandler
}

func New() *queBuilder {
	return &queBuilder{}
}

func (q *queBuilder) Define(request Request, requestHandler RequestHandler) *queBuilder {
	q.items[request.GetSignature()] = requestHandler

	return q
}

func (q *queBuilder) Build() Que {
	return &que{q}
}
