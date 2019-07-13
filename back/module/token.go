package module

import (
	"../repository"
	"time"
)

type TokenModuleImpl struct {
	AccessTokenRepo  repository.AccessTokenRepository
	RefreshTokenRepo repository.RefreshTokenRepository
}

func NewTokenModule(atr *repository.AccessTokenRepository, rtr *repository.RefreshTokenRepository) TokenModule {
	return &TokenModuleImpl{AccessTokenRepo: atr, RefreshTokenRepo: rtr}
}

func (tm *TokenModuleImpl) CreateToken(managerAccountId string) (accessToken string, refreshToken string, err error) {
	accessToken, err = CreateUUIDWithoutHyphen()
	if err != nil {
		return "", "", err
	}

	refreshToken, err = CreateUUIDWithoutHyphen()
	if err != nil {
		return "", "", err
	}

	// accessToken 1week
	tm.AccessTokenRepo.Add(accessToken, managerAccountId, time.Now().Unix()+60*60*24*7)
	// refreshToken 30days
	tm.RefreshTokenRepo.Add(refreshToken, managerAccountId, time.Now().Unix()+60*60*24*30)

	return accessToken, refreshToken, nil
}

func (tm *TokenModuleImpl) RefreshToken(managerAccountId string, refreshToken string) (newAccessToken string, newRefreshToken string, err error) {
	maid, exp, err := tm.RefreshTokenRepo.FindOne(refreshToken)
	if err != nil {
		return "", "", err
	}

	if isExpiredToken(exp) {
		return "", "", &ExpiredRefreshTokenError{err.Error()}
	}

	if maid != managerAccountId {
		return "", "", &InvalidRequestError{err.Error()}
	}

	newAccessToken, err = CreateUUIDWithoutHyphen()
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err = CreateUUIDWithoutHyphen()
	if err != nil {
		return "", "", err
	}

	// accessToken 1week
	tm.AccessTokenRepo.Update(newAccessToken, managerAccountId, time.Now().Unix()+60*60*24*7)
	// refreshToken 30days
	tm.RefreshTokenRepo.Update(newRefreshToken, managerAccountId, time.Now().Unix()+60*60*24*30)

	return newAccessToken, newRefreshToken, nil
}

func (tm *TokenModuleImpl) IsValidAccessToken(accessToken string, managerAccountId string) bool {
	maid, exp, err := tm.AccessTokenRepo.FindOne(accessToken)

	if err != nil || isExpiredToken(exp) || maid != managerAccountId {
		return false
	}

	return true
}

// 期限切れでないかを返します
func isExpiredToken(exp int64) bool {
	if exp < time.Now().Unix() {
		return true
	} else {
		return false
	}
}
