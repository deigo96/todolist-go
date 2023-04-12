package domain

type EmptyObj struct{}

// Response is used for static shape json return
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildError(status, message string) Response {
	res := Response{
		Status:  status,
		Message: message,
	}

	return res
}
