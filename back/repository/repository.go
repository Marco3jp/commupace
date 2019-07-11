package repository

import "../model"

type ManagerAccountRepository interface {
	Add(managerAccount *model.ManagerAccount) (id uint, err error)
	FindOne(id uint) (managerAccount *model.ManagerAccount, err error)
	Update(newManagerAccount *model.ManagerAccount) error
	Delete(id uint) error
}

type CommunityAccountRepository interface {
	Add(newCommunityAccount model.CommunityAccount) (id uint, err error)
	FindOne(id uint) (communityAccount *model.CommunityAccount, err error)
	FindFromManagerAccount(managerAccountId uint) (communityAccount []model.CommunityAccount, err error)
	Update(newCommunityAccount *model.CommunityAccount) error
	Delete(id uint) error
}