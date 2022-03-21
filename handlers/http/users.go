package http

import (
	"fmt"
	"net/http"

	"github.com/purp1eeeee/go-tech-dojo/oapi"
)

func (h Handlers) GetUserGet(w http.ResponseWriter, r *http.Request, params oapi.GetUserGetParams) {
	fmt.Println("getUserGet")
}

func (h Handlers) PostUserCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostUserCreate")
}

func (h Handlers) PutUserUpdate(w http.ResponseWriter, r *http.Request, params oapi.PutUserUpdateParams) {
}
