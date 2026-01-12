package response

type Response[T any] struct {
	HTTPStatus   int    `json:"status"`
	InternalCode Code   `json:"code"`
	Message      string `json:"message"`
	Data         T      `json:"data,omitempty"`
	Errors       any    `json:"errors,omitempty"`
}
