package http

import (
	"log"
	"net/http"

	"github.com/purp1eeeee/go-tech-dojo/handlers/http/utils"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
	"github.com/purp1eeeee/go-tech-dojo/repositories/users"
)

func (h Handlers) GetUserGet(w http.ResponseWriter, r *http.Request, params oapi.GetUserGetParams) {
	ctx := r.Context()

	me, err := h.userUseCase.GetMe(ctx, params.XToken)
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}
	utils.RenderJson(ctx, w, r, oapi.UserGetResponse{
		Name: &me.Name,
	})
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
	ctx := r.Context()

	var body oapi.UserUpdateRequest
	err := utils.ParseBody(r, &body)
	if err != nil {
		log.Println(err)
		utils.RenderBadRequest(ctx, w, r)
		return
	}

	me, err := h.userUseCase.GetMe(ctx, params.XToken)
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}

	h.userUseCase.Update(ctx, users.User{
		ID:   me.ID.String(),
		Name: *body.Name,
	})

	utils.RenderNoContent(ctx, w, r)
}
