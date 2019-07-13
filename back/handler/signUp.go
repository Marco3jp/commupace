package handler

import "github.com/gin-gonic/gin"

func (a *APIHandler) SignUp(c *gin.Context) {
	managerAccountId, err := a.ManagerAccountModule.CreateManagerAccount()
	if err != nil {
		c.JSON(500, gin.H{
			"ok":  false,
			"msg": "Failed create manager account",
		})
		return
	}

	accessToken, refreshToken, err := a.TokenModule.CreateToken(managerAccountId)
	if err != nil {
		c.JSON(500, gin.H{
			"ok":  false,
			"msg": "Failed create token",
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":           true,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
	return
}
