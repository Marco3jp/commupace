package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type ManagerAccountRepositoryImpl struct {
	db *gorm.DB
}

func NewManagerAccountRepository(db *gorm.DB) repository.ManagerAccountRepository {
	return &ManagerAccountRepositoryImpl{db: db}
}

func (mar *ManagerAccountRepositoryImpl) Add(newManagerAccount *model.ManagerAccount) (id uint, err error) {
	if !mar.db.NewRecord(*newManagerAccount) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := mar.db.Create(newManagerAccount).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newManagerAccount.ID, nil
}

func (mar *ManagerAccountRepositoryImpl) FindOne(id uint) (managerAccount *model.ManagerAccount, err error) {
	if mar.db.First(managerAccount, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: ManagerAccountTable"}
	}

	return managerAccount, nil
}

func (mar *ManagerAccountRepositoryImpl) Update(newManagerAccount *model.ManagerAccount) error {
	if err := mar.db.Save(newManagerAccount).Error; err != nil {
		return &repository.IOError{err.Error()}
	}
	return nil
}

func (mar *ManagerAccountRepositoryImpl) Delete(id uint) error {
	target := model.ManagerAccount{}
	target.ID = id

	if err := mar.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
