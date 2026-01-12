package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"cbe-error-response/response"
)

func writeResp(w http.ResponseWriter, resp response.Response[any]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.HTTPStatus)
	_ = json.NewEncoder(w).Encode(resp)
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Bod       string `json:"bod"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func main() {
	r := chi.NewRouter()
	provider := response.NewResponseProvider()

	makeUsers := func() []User {
		now := time.Now().Format(time.RFC3339)
		return []User{
			{ID: "1", Name: "Alice", Gender: "female", Bod: "1990-01-01", Status: "active", CreatedAt: now},
			{ID: "2", Name: "Bob", Gender: "male", Bod: "1985-05-12", Status: "active", CreatedAt: now},
			{ID: "3", Name: "Carol", Gender: "female", Bod: "1992-09-30", Status: "inactive", CreatedAt: now},
		}
	}

	r.Get("/ok", func(w http.ResponseWriter, r *http.Request) {
		resp := response.Success(provider, response.ResponseParam[any]{
			Code: response.CR200,
			Data: map[string]string{"message": "all good"},
		})
		writeResp(w, resp)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users := makeUsers()
		resp := response.Success(provider, response.ResponseParam[[]User]{
			Code: response.CR200,
			Data: users,
		})
		writeResp(w, resp)
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		users := makeUsers()
		for _, u := range users {
			if u.ID == id {
				resp := response.Success(provider, response.ResponseParam[User]{
					Code: response.CR200,
					Data: u,
				})
				writeResp(w, resp)
				return
			}
		}
		resp := response.ClientError(provider, response.ResponseParam[User]{
			Code:   response.CR404,
			Errors: map[string]string{"message": "user not found", "id": id},
		})
		writeResp(w, resp)
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		resp := response.ClientError(provider, response.ResponseParam[any]{
			Code:   response.CR400,
			Errors: map[string]string{"field": "name", "reason": "required"},
		})
		writeResp(w, resp)
	})

	log.Println("listening :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
