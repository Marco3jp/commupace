package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type CommunityUserRepositoryImpl struct {
	db *gorm.DB
}

func CommunityUserRepository(db *gorm.DB) repository.CommunityUserRepository {
	return &CommunityUserRepositoryImpl{db: db}
}

func (cur *CommunityUserRepositoryImpl) Add(newCommunityUser *model.CommunityUser) (id uint, err error) {
	if !cur.db.NewRecord(*newCommunityUser) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := cur.db.Create(newCommunityUser).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newCommunityUser.ID, nil

}
func (cur *CommunityUserRepositoryImpl) FindOne(id uint) (communityUser *model.CommunityUser, err error) {
	if cur.db.First(communityUser, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: communityUserTable"}
	}

	return communityUser, nil

}
func (cur *CommunityUserRepositoryImpl) FindFromCommunity(communityId uint) (communityUsers []model.CommunityUser, err error) {
	if cur.db.Where("community_id = ?", communityId).Find(communityUsers).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: communityUserTable"}
	}

	return communityUsers, nil

}
func (cur *CommunityUserRepositoryImpl) FindFromCommunityAccount(communityAccountId uint) (communityUsers []model.CommunityUser, err error) {
	if cur.db.Where("community_account_id = ?", communityAccountId).Find(communityUsers).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: communityUserTable"}
	}

	return communityUsers, nil

}
func (cur *CommunityUserRepositoryImpl) FetchRangeFromCommunity(communityId uint, count uint) (communityUsers []model.CommunityUser, err error) {
	if cur.db.Where("community_id = ?", communityId).Last(communityUsers, count).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: communityUserTable"}
	}

	return communityUsers, nil
}
func (cur *CommunityUserRepositoryImpl) FetchRangeFromCommunityAccount(communityAccountId uint, count uint) (communityUsers []model.CommunityUser, err error) {
	if cur.db.Where("community_account_id = ?", communityAccountId).Last(communityUsers, count).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: communityUserTable"}
	}

	return communityUsers, nil
}
func (cur *CommunityUserRepositoryImpl) Delete(id uint) error {
	target := model.CommunityUser{}
	target.ID = id

	if err := cur.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
