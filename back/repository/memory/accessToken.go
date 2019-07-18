package memory

import (
	"sync"
	"github.com/Marco3jp/commupace/back/repository"
	"github.com/Marco3jp/commupace/back/model"
)

type AccessTokenRepositoryImpl struct {
	m *sync.Map
}

func NewAccessTokenRepository(m *sync.Map) repository.AccessTokenRepository {
	return &AccessTokenRepositoryImpl{m: m}
}

func (atr *AccessTokenRepositoryImpl) Add(token string, managerAccountId string, exp int64) {
	values := model.AccessTokenValues{managerAccountId, exp}
	atr.m.Store(token, values)
}

func (atr *AccessTokenRepositoryImpl) FindOne(token string) (managerAccountId string, exp int64, err error) {
	result, ok := atr.m.Load(token)
	if !ok {
		return "", 0, &repository.NotFoundRecordError{"Action: AccessTokenMemory"}
	}

	return result.(model.AccessTokenValues).ManagerAccountID, result.(model.AccessTokenValues).Exp , nil
}

func (atr *AccessTokenRepositoryImpl) Update(token string, managerAccountId string, exp int64) {
	values := model.AccessTokenValues{managerAccountId, exp}
	atr.m.Store(token, values)
}

func (atr *AccessTokenRepositoryImpl) Delete(token string) {
	atr.m.Delete(token)
}
