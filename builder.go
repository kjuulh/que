package que

type queBuilder struct {
	items map[string]any
}

func New() *queBuilder {
	return &queBuilder{
		items: make(map[string]any, 0),
	}
}

func Define[TRequest Request](queBuilder *queBuilder, requestHandler RequestHandlerContract[TRequest]) *queBuilder {
	signature := requestHandler.GetSignature()

	queBuilder.items[signature] = requestHandler.Handle

	return queBuilder
}

func (q *queBuilder) Build() *queBuilder {
	return q
}
