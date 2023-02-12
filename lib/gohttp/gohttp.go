package gohttp

// NewClient new request object
func NewClient(opts ...Options) *Request {
	req := &Request{}

	if len(opts) > 0 {
		req.opts = opts[0]
	} else {
		req.opts = Options{}
	}

	return req
}

// Get send get request
func Get(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Get(uri, opts...)
}

// Post send post request
func Post(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Post(uri, opts...)
}

// Put send put request
func Put(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Post(uri, opts...)
}

// Patch send patch request
func Patch(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Patch(uri, opts...)
}

// Delete send delete request
func Delete(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Delete(uri, opts...)
}

// Download file
func FastGet(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.FastGet(uri, opts...)
}
