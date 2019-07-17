package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"../model"
)

type parsedPostData struct {
	model.Post
	CommunityAccount model.CommunityAccount
}

// TODO: JoinCommunityが実装されたとき、CommunityAccountのチェックも行う
func (a *APIHandler) GetChat(c *gin.Context) {
	managerAccountId := c.GetHeader("manager-account-id")
	if !a.TokenModule.IsValidAccessToken(c.GetHeader("access-token"), managerAccountId) {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": false,
			"msg":                "Your token is not valid.",
		})
		return
	}
	rawCommunityId := c.Query("communityId")
	rawCount := c.Query("count")

	if rawCommunityId == "" {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request don't have required query string.",
		})
		return
	}

	if rawCount == "" {
		rawCount = "10"
	}

	communityId, communityIdErr := strconv.ParseUint(rawCommunityId, 10, 64)
	count, countErr := strconv.ParseUint(rawCount, 10, 64)

	if communityIdErr != nil || countErr != nil {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request query string is incorrect type.",
		})
		return
	}

	posts, err := a.ChatModule.GetPosts(uint(communityId), uint(count))
	if err != nil {
		c.JSON(500, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Internal Server Error",
		})
		return
	}

	result := make([]parsedPostData, len(posts))

	for e := range posts {
		result[e].ID = posts[e].Post.ID
		result[e].CreatedAt = posts[e].Post.CreatedAt
		result[e].UpdatedAt = posts[e].Post.UpdatedAt
		result[e].CommunityAccountID = posts[e].CommunityAccountID
		result[e].ThreadID = posts[e].ThreadID
		result[e].CommunityID = posts[e].CommunityID
		result[e].PostText = posts[e].PostText
		result[e].CommunityAccountID = posts[e].CommunityAccountID
		result[e].PostNumber = posts[e].PostNumber
		result[e].PostType = posts[e].PostType
		result[e].PostPath = posts[e].PostPath
		result[e].CommunityAccount = model.CommunityAccount{
			DisplayID:   posts[e].DisplayID,
			DisplayName: posts[e].DisplayName,
			Icon:        posts[e].Icon,
			Status:      posts[e].Status,
		}
		result[e].CommunityAccount.ID = posts[e].CommunityAccount.ID
	}

	c.JSON(200, gin.H{
		"ok":                 true,
		"isValidAccessToken": true,
		"posts":              result,
	})
	return
}
