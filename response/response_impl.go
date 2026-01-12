package response

import "net/http"

type ResponseProvider struct{}

func NewResponseProvider() Provider {
	return &ResponseProvider{}
}

type ResponseParam[T any] struct {
	Code   Code
	Data   T
	Errors any
}

// Generic helper functions for type-safe usage
func Success[T any](provider Provider, param ResponseParam[T]) Response[any] {
	return provider.Success(ResponseParam[any]{
		Code:   param.Code,
		Data:   param.Data,
		Errors: param.Errors,
	})
}

func ClientError[T any](provider Provider, param ResponseParam[T]) Response[any] {
	return provider.ClientError(ResponseParam[any]{
		Code:   param.Code,
		Data:   param.Data,
		Errors: param.Errors,
	})
}

func ServerError[T any](provider Provider, param ResponseParam[T]) Response[any] {
	return provider.ServerError(ResponseParam[any]{
		Code:   param.Code,
		Data:   param.Data,
		Errors: param.Errors,
	})
}

func (e *ResponseProvider) Success(param ResponseParam[any]) Response[any] {
	switch param.Code {
	case CR200:
		return Response[any]{
			HTTPStatus:   http.StatusOK,
			InternalCode: CR200,
			Message:      "Response Successfully!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR201:
		return Response[any]{
			HTTPStatus:   http.StatusCreated,
			InternalCode: CR201,
			Message:      "Created Successfully!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR203:
		return Response[any]{
			HTTPStatus:   http.StatusAccepted,
			InternalCode: CR203,
			Message:      "Service Accepted",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR204:
		return Response[any]{
			HTTPStatus:   http.StatusNoContent,
			InternalCode: CR204,
			Message:      "Success, No Content!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	default:
		return Response[any]{
			HTTPStatus:   http.StatusOK,
			InternalCode: "",
			Message:      "Unknown Success Code",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	}
}

func (e *ResponseProvider) ClientError(param ResponseParam[any]) Response[any] {
	switch param.Code {
	case CR400:
		return Response[any]{
			HTTPStatus:   http.StatusBadRequest,
			InternalCode: CR400,
			Message:      "Invalid Request!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR401:
		return Response[any]{
			HTTPStatus:   http.StatusUnauthorized,
			InternalCode: CR401,
			Message:      "Authorization Required!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR402:
		return Response[any]{
			HTTPStatus:   http.StatusForbidden,
			InternalCode: CR402,
			Message:      "Access Denied!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR403:
		return Response[any]{
			HTTPStatus:   http.StatusNotFound,
			InternalCode: CR403,
			Message:      "Content/Resource Not Found!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR404:
		return Response[any]{
			HTTPStatus:   http.StatusNotFound,
			InternalCode: CR404,
			Message:      "Resource Not Found!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR405:
		return Response[any]{
			HTTPStatus:   http.StatusMethodNotAllowed,
			InternalCode: CR405,
			Message:      "Method Not Allowed!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR406:
		return Response[any]{
			HTTPStatus:   http.StatusConflict,
			InternalCode: CR406,
			Message:      "Conflict!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR407:
		return Response[any]{
			HTTPStatus:   http.StatusUnprocessableEntity,
			InternalCode: CR407,
			Message:      "Validation Error!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	default:
		return Response[any]{
			HTTPStatus:   http.StatusBadRequest,
			InternalCode: "",
			Message:      "Unknown Client Error Code",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	}
}

func (e *ResponseProvider) ServerError(param ResponseParam[any]) Response[any] {
	switch param.Code {
	case CR500:
		return Response[any]{
			HTTPStatus:   http.StatusInternalServerError,
			InternalCode: CR500,
			Message:      "Internal Server Error!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR501:
		return Response[any]{
			HTTPStatus:   http.StatusNotImplemented,
			InternalCode: CR501,
			Message:      "Not Implemented!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR502:
		return Response[any]{
			HTTPStatus:   http.StatusBadGateway,
			InternalCode: CR502,
			Message:      "Bad Gateway!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR503:
		return Response[any]{
			HTTPStatus:   http.StatusServiceUnavailable,
			InternalCode: CR503,
			Message:      "Service Unavailable!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	case CR504:
		return Response[any]{
			HTTPStatus:   http.StatusGatewayTimeout,
			InternalCode: CR504,
			Message:      "Gateway Timeout!",
			Data:         param.Data,
			Errors:       param.Errors,
		}
	default:
		return Response[any]{
			HTTPStatus:   http.StatusInternalServerError,
			InternalCode: "",
			Message:      "Unknown Error",
			Data:         param.Data,
			Errors:       param.Errors,
		}

	}
}
