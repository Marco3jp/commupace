package module

import (
	"github.com/Marco3jp/commupace/back/repository"
	"github.com/Marco3jp/commupace/back/model"
)

type ManagerAccountModuleImpl struct {
	ManagerAccountRepository repository.ManagerAccountRepository
}

func NewManagerAccountModule(mar repository.ManagerAccountRepository) ManagerAccountModule {
	return &ManagerAccountModuleImpl{ManagerAccountRepository: mar}
}

func (mam *ManagerAccountModuleImpl) CreateManagerAccount() (managerAccountId string, err error) {
	managerAccountId, err = CreateUUIDWithoutHyphen()
	if err != nil {
		return "", err
	}

	managerAccount := &model.ManagerAccount{
		ManagerAccountID: managerAccountId,
	}

	_, err = mam.ManagerAccountRepository.Add(managerAccount)
	if err != nil {
		return "", err
	}

	return managerAccountId, nil
}
