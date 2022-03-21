package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/purp1eeeee/go-tech-dojo/handlers/http/utils"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
)

func (h Handlers) PostGachaDraw(w http.ResponseWriter, r *http.Request, params oapi.PostGachaDrawParams) {
	ctx := r.Context()

	var body oapi.PostGachaDrawJSONRequestBody
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

	times, err := strconv.Atoi(*body.Times)
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}

	characters, err := h.gachaUseCase.Drow(ctx, me.ID.String(), times)
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}
	results := make([]oapi.GachaResult, len(characters))
	for i, v := range characters {
		id := strconv.FormatUint(v.ID, 10)
		name := v.Name
		results[i] = oapi.GachaResult{
			CharacterID: &id,
			Name:        &name,
		}
	}

	utils.RenderJson(ctx, w, r, oapi.GachaDrawResponse{
		Results: &results,
	})

}
