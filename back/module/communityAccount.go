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

func NewCommunityAccountModuleImpl(communityAccountRepo *repository.CommunityAccountRepository) CommunityAccountModule {
	return &CommunityAccountModuleImpl{CommunityAccountRepo: communityAccountRepo}
}

func (cam *CommunityAccountModuleImpl) CreateCommunityAccount(managerAccountId string, communityAccount *model.CommunityAccount) (err error) {
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
