package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type CommunityAccountRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityAccountRepository(db *gorm.DB) repository.CommunityAccountRepository {
	return &CommunityAccountRepositoryImpl{db: db}
}

func (car *CommunityAccountRepositoryImpl) Add(newCommunityAccount *model.CommunityAccount) (id uint, err error) {
	if !car.db.NewRecord(*newCommunityAccount) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := car.db.Create(newCommunityAccount).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newCommunityAccount.ID, nil
}

func (car *CommunityAccountRepositoryImpl) FindOne(id uint) (communityAccount *model.CommunityAccount, err error) {
	if car.db.First(communityAccount, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityAccountTable"}
	}

	return communityAccount, nil
}

func (car *CommunityAccountRepositoryImpl) FindOneFromDisplayId(displayId string) (communityAccount *model.CommunityAccount, err error) {
	if car.db.Where("display_id = ?", displayId).First(communityAccount).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityAccountTable"}
	}

	return communityAccount, nil
}

func (car *CommunityAccountRepositoryImpl) FindFromManagerAccount(managerAccountId uint) (communityAccounts []model.CommunityAccount, err error) {
	if car.db.Where("manager_account_id = ?", managerAccountId).Find(communityAccounts).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityAccountTable"}
	}

	return communityAccounts, nil
}

func (car *CommunityAccountRepositoryImpl) Update(newCommunityAccount *model.CommunityAccount) error {
	if err := car.db.Save(newCommunityAccount).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}

func (car *CommunityAccountRepositoryImpl) Delete(id uint) error {
	target := model.CommunityAccount{}
	target.ID = id

	if err := car.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
