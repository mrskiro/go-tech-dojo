package http

import "github.com/purp1eeeee/go-tech-dojo/oapi"

type Handlers struct{}

func NewHandlers() oapi.ServerInterface {
	return Handlers{}
}
