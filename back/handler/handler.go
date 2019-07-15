package handler

import (
	"../module"
)

type APIHandler struct {
	TokenModule            module.TokenModule
	ManagerAccountModule   module.ManagerAccountModule
	CommunityAccountModule module.CommunityAccountModule
	CommunityModule        module.CommunityModule
	ChatModule             module.ChatModule
}

func NewAPIHandler(
	tokenModule module.TokenModule,
	managerAccountModule module.ManagerAccountModule,
	communityAccountModule module.CommunityAccountModule,
	communityModule module.CommunityModule,
	chatModule module.ChatModule) *APIHandler {

	return &APIHandler{
		TokenModule:            tokenModule,
		ManagerAccountModule:   managerAccountModule,
		CommunityAccountModule: communityAccountModule,
		CommunityModule:        communityModule,
		ChatModule:             chatModule,
	}
}
