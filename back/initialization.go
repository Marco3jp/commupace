package main

import (
	"github.com/gin-gonic/gin"
	"./repository/database"
	"./repository/memory"
	"./module"
	"./handler"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	tokenModule            module.TokenModule
	managerAccountModule   module.ManagerAccountModule
	communityAccountModule module.CommunityAccountModule
	apiHandler             *handler.APIHandler
)

func initObjects(db *gorm.DB) {
	mar := database.NewManagerAccountRepository(db)
	car := database.NewCommunityAccountRepository(db)
	// cur := database.NewCommunityUserRepository(db)
	// lr := database.NewLocationRepository(db)
	// sr := database.NewSpaceRepository(db)
	// cr := database.NewCommunityRepository(db)
	// pr := database.NewPostRepository(db)
	atr := memory.NewAccessTokenRepository(&sync.Map{})
	rtr := memory.NewRefreshTokenRepository(&sync.Map{})

	tokenModule = module.NewTokenModule(atr, rtr)
	managerAccountModule = module.NewManagerAccountModule(mar)
	communityAccountModule = module.NewCommunityAccountModule(car)

	apiHandler = handler.NewAPIHandler(tokenModule, managerAccountModule, communityAccountModule)
}

func initRouting(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/sign_up", apiHandler.SignUp) // return new manager_account token, behave like "create_manager_account"
			// auth.POST("/sign_in")
			// auth.POST("/refresh")
		}
		/*
				managerAccount := api.Group("/manager_account")
				{
					managerAccount.POST("/add_community_account")
					managerAccount.GET("/community_account_list")
					// managerAccount.PATCH("/info")
				}

				communityAccount := api.Group("/community_account")
				{
					communityAccount.GET("/info")
					communityAccount.PATCH("/info")
				}

				community := api.Group("/community")
				{
					community.POST("/create")
					community.GET("/search")
					community.GET("/info")
					community.PATCH("/info")
				}

				chat := community.Group("/chat")
				{
					chat.GET("/post")
					chat.POST("/post")
				}
			*/
	}
}
