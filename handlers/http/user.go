package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/purp1eeeee/go-tech-dojo/handlers/http/utils"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
)

func (h Handlers) GetUserGet(w http.ResponseWriter, r *http.Request, params oapi.GetUserGetParams) {
	fmt.Println("getUserGet")
}

func (h Handlers) PostUserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body oapi.UserCreateRequest
	err := utils.ParseBody(r, &body)
	if err != nil {
		log.Println(err)
		utils.RenderBadRequest(ctx, w, r)
		return
	}
	token, err := h.userUseCase.Register(ctx, *body.Name)
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}

	utils.RenderJson(ctx, w, r, oapi.UserCreateResponse{
		Token: &token,
	})
}

func (h Handlers) PutUserUpdate(w http.ResponseWriter, r *http.Request, params oapi.PutUserUpdateParams) {
}
