package handler

import "github.com/gin-gonic/gin"

type PostChatRequestBody struct {
	CommunityAccountId uint   `json:"communityAccountId"`
	ThreadId           uint   `json:"threadId"`
	CommunityId        uint   `json:"communityId"`
	PostText           string `json:"postText"`
	PostPath           string `json:"postPath"`
}

func (a *APIHandler) PostChat(c *gin.Context) {
	managerAccountId := c.GetHeader("manager-account-id")
	if !a.TokenModule.IsValidAccessToken(c.GetHeader("access-token"), managerAccountId) {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": false,
			"msg":                "Your token is not valid.",
		})
		return
	}

	body := PostChatRequestBody{}
	c.BindJSON(&body)
	if body.CommunityAccountId == 0 || body.CommunityId == 0 || body.PostText == "" || body.PostPath == "" {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request body don't have some params.",
		})
		return
	}

	if !a.CommunityAccountModule.IsManagedAccount(managerAccountId, body.CommunityAccountId) {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "You are not this manager account owner. PLZ DON'T SPOOFING.",
		})
		return
	}

	post := a.ChatModule.CreatePost(body.CommunityAccountId, body.CommunityId, body.PostText, body.PostPath)
	err := a.ChatModule.Post(post)
	if err != nil {
		c.JSON(500, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":                 true,
		"isValidAccessToken": true,
		"yourPost":           post,
	})
	return
}
