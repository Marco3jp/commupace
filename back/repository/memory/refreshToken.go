package memory

import (
	"sync"
	".."
	"../../model"
)

type RefreshTokenRepositoryImpl struct {
	m *sync.Map
}

func NewRefreshTokenRepository(m *sync.Map) repository.RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{m: m}
}

func (rtr *RefreshTokenRepositoryImpl) Add(token string, managerAccountId string, exp int64) {
	values := model.RefreshTokenValues{managerAccountId, exp}
	rtr.m.Store(managerAccountId, values)
}

func (rtr *RefreshTokenRepositoryImpl) FindOne(token string) (managerAccountId string, exp int64, err error) {
	result, ok := rtr.m.Load(token)
	if !ok {
		return "", 0, &repository.NotFoundRecordError{"Action: RefreshTokenMemory"}
	}

	return result.(model.RefreshTokenValues).ManagerAccountID, result.(model.RefreshTokenValues).Exp, nil
}

func (rtr *RefreshTokenRepositoryImpl) Update(token string, managerAccountId string, exp int64) {
	values := model.RefreshTokenValues{managerAccountId, exp}
	rtr.m.Store(managerAccountId, values)
}

func (rtr *RefreshTokenRepositoryImpl) Delete(token string) {
	rtr.m.Delete(token)
}
