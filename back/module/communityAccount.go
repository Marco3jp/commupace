package module

import (
	"../repository"
	"../model"
)

type CommunityAccountModuleImpl struct {
	ManagerAccountRepo   repository.ManagerAccountRepository
	CommunityAccountRepo repository.CommunityAccountRepository
	CommunityUserRepo    repository.CommunityUserRepository
}

func NewCommunityAccountModule(mar repository.ManagerAccountRepository, car repository.CommunityAccountRepository, cur repository.CommunityUserRepository) CommunityAccountModule {
	return &CommunityAccountModuleImpl{ManagerAccountRepo: mar, CommunityAccountRepo: car, CommunityUserRepo: cur}
}

func (cam *CommunityAccountModuleImpl) CreateCommunityAccount(managerAccountId string, communityAccount *model.CommunityAccount)  error {
	managerAccount, err := cam.ManagerAccountRepo.FindOneFromManagerAccountId(managerAccountId)
	if err != nil {
		return err
	}

	communityAccount.ManagerAccountID = managerAccount.ID

	_, err = cam.CommunityAccountRepo.Add(communityAccount)
	if err != nil {
		return err
	}
	return nil
}

func (cam *CommunityAccountModuleImpl) IsManagedAccount(managerAccountId string, communityAccountId uint) bool {
	managerAccount, err := cam.ManagerAccountRepo.FindOneFromManagerAccountId(managerAccountId)
	if err != nil {
		return false
	}

	communityAccount, err := cam.CommunityAccountRepo.FindOne(communityAccountId)
	if err != nil {
		return false
	}

	return managerAccount.ID == communityAccount.ManagerAccountID
}

func (cam *CommunityAccountModuleImpl) JoinCommunity(communityAccountId string, communityId string) error {
	communityAccount, err := cam.CommunityAccountRepo.FindOneFromDisplayId(communityAccountId)
	if err != nil {
		return err
	}

	community, err := cam.CommunityAccountRepo.FindOneFromDisplayId(communityId)
	if err != nil {
		return err
	}

	communityUser := &model.CommunityUser{
		CommunityAccountID: communityAccount.ID,
		CommunityID:        community.ID,
	}

	cam.CommunityUserRepo.Add(communityUser)

	return nil
}
