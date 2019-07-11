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
	FindFromManagerAccount(managerAccountId uint) (communityAccounts []model.CommunityAccount, err error)
	Update(newCommunityAccount *model.CommunityAccount) error
	Delete(id uint) error
}

type LocationRepository interface {
	Add(newLocation model.Location) (id uint, err error)
	FindOne(id uint) (location *model.Location, err error)
	FindOneFromCoordinates(coordinates model.Coordinates) (locations []model.Location, err error)
	Update(newLocation model.Location) error
	Delete(id uint) error
}
