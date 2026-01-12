package main

import (
	"cbe-error-response/response"
	"fmt"
)

type User struct {
	ID       string
	Fullname string
	Gender   string
}

func main() {
	// Create a new response provider
	provider := response.NewResponseProvider()
	// user := User{
	// 	ID:       "11",
	// 	Fullname: "Jeo",
	// 	Gender:   "Male",
	// }

	userList := make([]User, 0, 3)
	for i := 0; i < 3; i++ {
		userList = append(userList, User{
			ID:       fmt.Sprintf("%d", i+1),
			Fullname: "Jeo",
			Gender:   "Male",
		})
	}

	// Example: Success responses
	success200 := response.Success(provider, response.ResponseParam[[]User]{
		Code: response.CR200,
		Data: userList,
	})
	fmt.Printf("Success 200: HTTP Status=%d, Code=%s, Message=%s, Data=%v\n",
		success200.HTTPStatus, success200.InternalCode, success200.Message, success200.Data)

	success201 := response.Success(provider, response.ResponseParam[any]{Code: response.CR200})
	fmt.Printf("Success 201: HTTP Status=%d, Code=%s, Message=%s\n",
		success201.HTTPStatus, success201.InternalCode, success201.Message)

	// Example: Client error responses
	clientError400 := response.ClientError(provider, response.ResponseParam[any]{Code: response.CR400})
	fmt.Printf("Client Error 400: HTTP Status=%d, Code=%s, Message=%s\n",
		clientError400.HTTPStatus, clientError400.InternalCode, clientError400.Message)

	clientError404 := response.ClientError(provider, response.ResponseParam[any]{Code: response.CR404})
	fmt.Printf("Client Error 404: HTTP Status=%d, Code=%s, Message=%s\n",
		clientError404.HTTPStatus, clientError404.InternalCode, clientError404.Message)

	// Example: Server error responses
	serverError500 := response.ServerError(provider, response.ResponseParam[any]{Code: response.CR500})
	fmt.Printf("Server Error 500: HTTP Status=%d, Code=%s, Message=%s\n",
		serverError500.HTTPStatus, serverError500.InternalCode, serverError500.Message)
}
