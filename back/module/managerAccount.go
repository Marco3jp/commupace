package module

import (
	"../repository"
	"../model"
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
		Email:            nil,
		Password:         nil,
	}

	_, err = mam.ManagerAccountRepository.Add(managerAccount)
	if err != nil {
		return "", err
	}

	return managerAccountId, nil
}
