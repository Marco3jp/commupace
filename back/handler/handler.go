package handler

import (
	"../module"
)

type APIHandler struct {
	TokenModule            module.TokenModule
	ManagerAccountModule   module.ManagerAccountModule
	CommunityAccountModule module.CommunityAccountModule
}

func NewAPIHandler(tokenModule module.TokenModule, managerAccountModule module.ManagerAccountModule, communityAccountModule module.CommunityAccountModule) *APIHandler {
	return &APIHandler{
		TokenModule:            tokenModule,
		ManagerAccountModule:   managerAccountModule,
		CommunityAccountModule: communityAccountModule,
	}
}
