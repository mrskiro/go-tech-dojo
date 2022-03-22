package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/purp1eeeee/go-tech-dojo/handlers/http/utils"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
)

func (h Handlers) GetCharacterList(w http.ResponseWriter, r *http.Request, params oapi.GetCharacterListParams) {
	ctx := r.Context()
	me, err := h.userUseCase.GetMe(ctx, params.XToken)
	if err != nil {
		log.Println(err)
		utils.RenderUnauthorized(ctx, w, r)
		return
	}

	cs, err := h.characterUseCase.GetList(ctx, me.ID.String())
	if err != nil {
		log.Println(err)
		utils.RenderInternalServerError(ctx, w, r)
		return
	}

	results := make([]oapi.UserCharacter, len(cs))
	for i, v := range cs {
		userCharacterID := me.ID.String()
		characterID := strconv.Itoa(int(v.ID))
		name := v.Name
		results[i] = oapi.UserCharacter{
			UserCharacterID: &userCharacterID,
			CharacterID:     &characterID,
			Name:            &name,
		}
	}

	utils.RenderJson(ctx, w, r, oapi.CharacterListResponse{
		Characters: &results,
	})
}
