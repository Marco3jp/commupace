package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/Marco3jp/commupace/back/model"
)

func (a *APIHandler) SearchCommunity(c *gin.Context) {
	if !a.TokenModule.IsValidAccessToken(c.GetHeader("access-token"), c.GetHeader("manager-account-id")) {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": false,
			"msg":                "Your token is not valid.",
		})
		return
	}

	rawLat := c.Query("lat")
	rawLong := c.Query("long")
	rawZoomLevel := c.Query("zoom")
	if rawLat == "" || rawLong == "" || rawZoomLevel == "" {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request don't have required query string.",
		})
		return
	}

	lat, latErr := strconv.ParseFloat(rawLat, 64)
	long, longErr := strconv.ParseFloat(rawLong, 64)
	zoom, zoomErr := strconv.ParseUint(rawZoomLevel, 10, 16)
	if latErr != nil || longErr != nil || zoomErr != nil {
		c.JSON(400, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Your request query string is incorrect type.",
		})
		return
	}

	communityList, err := a.CommunityModule.SearchCommunityFromCoordinates(model.Coordinates{lat, long}, uint(zoom))
	if err != nil {
		c.JSON(500, gin.H{
			"ok":                 false,
			"isValidAccessToken": true,
			"msg":                "Internal server error.",
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":                 true,
		"isValidAccessToken": true,
		"communityList":      communityList,
	})
	return
}
