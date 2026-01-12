package response

type Provider interface {
	Success(param ResponseParam[any]) Response[any]
	ClientError(param ResponseParam[any]) Response[any]
	ServerError(param ResponseParam[any]) Response[any]
}
