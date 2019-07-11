package repository

import "../model"

type ManagerAccountRepository interface {
	Add(managerAccount *model.ManagerAccount) (id uint, err error)
	FindOne(id uint) (managerAccount *model.ManagerAccount, err error)
	Update(newManagerAccount *model.ManagerAccount) error
	Delete(id uint) error
}