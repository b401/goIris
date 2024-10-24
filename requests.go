package iris

type RequestBuilder struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    interface{}
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Headers: make(map[string]string),
	}
}

func (rb *RequestBuilder) SetMethod(method string) *RequestBuilder {
	rb.Method = method
	return rb
}

func (rb *RequestBuilder) SetURL(url string) *RequestBuilder {
	rb.URL = url
	return rb
}

func (rb *RequestBuilder) AddHeader(key, value string) *RequestBuilder {
	rb.Headers[key] = value
	return rb
}

func (rb *RequestBuilder) SetBody(body interface{}) *RequestBuilder {
	rb.Body = body
	return rb
}

func (rb *RequestBuilder) Build() *RequestBuilder {
	return rb
}
