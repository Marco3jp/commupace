package handler

import (
	"github.com/gin-gonic/gin"
	"../model"
)

type AddCommunityAccountRequestBody struct {
	DisplayId   string `json:"displayId"`
	DisplayName string `json:"displayName"`
	Icon        string `json:"icon"`
	Status      string `json:"status"`
}

func (a *APIHandler) AddCommunityAccount(c *gin.Context) {
	managerAccountId := c.GetHeader("manager-account-id")
	if !a.TokenModule.IsValidAccessToken(c.GetHeader("access-token"), managerAccountId) {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": false,
			"msg":                "Your token is not valid.",
		})
		return
	}

	body := AddCommunityAccountRequestBody{}
	c.ShouldBindJSON(&body)
	if body.DisplayId == "" || body.DisplayName == "" {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request body don't have some params.",
		})
		return
	}

	communityAccount := &model.CommunityAccount{
		DisplayID:   body.DisplayId,
		DisplayName: body.DisplayName,
		Icon:        body.Icon,
		Status:      body.Status,
	}

	err := a.CommunityAccountModule.CreateCommunityAccount(managerAccountId, communityAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":                 true,
		"isValidAccessToken": true,
		"communityAccount":   communityAccount,
	})
	return
}
