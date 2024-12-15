package response

type GenericError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}

func (v *GenericError) Error() string {
	return v.Err.Error()
}
