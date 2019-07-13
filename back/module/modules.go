package module

type TokenModule interface {
	CreateToken(managerAccountId string) (accessToken string, refreshToken string, err error)
	RefreshToken(managerAccountId string, refreshToken string) (newAccessToken string, newRefreshToken string, err error)
	// AccessTokenが有効かチェックする。
	// 引数が空など、本来エラーであっても無効であることに変わりはないのでfalseが返る仕様になっている。
	IsValidAccessToken(accessToken string, managerAccountId string) bool
	// isValidRefreshToken()
}

type ManagerAccountModule interface {
	CreateManagerAccount() (managerAccountId string, err error)
	// UpdateManagerAccount()
}
