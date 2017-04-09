package common

type BusinessError struct {
	Code    string
	Message string
}

func (o *BusinessError) Error() string {
	if o == nil {
		return "<nil>"
	}
	if o.Message != "" {
		return o.Message
	}
	return o.Code
}
