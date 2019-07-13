package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type CommunityRepositoryImpl struct {
	db *gorm.DB
}

func CommunityRepository(db *gorm.DB) repository.CommunityRepository {
	return &CommunityRepositoryImpl{db: db}
}

func (cr *CommunityRepositoryImpl) Add(newCommunity *model.Community) (id uint, err error) {
	if !cr.db.NewRecord(*newCommunity) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := cr.db.Create(newCommunity).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newCommunity.ID, nil
}

func (cr *CommunityRepositoryImpl) FindOne(id uint) (community *model.Community, err error) {
	if cr.db.First(community, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityTable"}
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) FindOneFromCommunityId(communityId string) (community *model.Community, err error) {
	if cr.db.Where("community_id = ?", communityId).First(community).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityTable"}
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) FindFromSpace(spaceId uint) (communities []model.Community, err error) {
	if cr.db.Where("space_id = ?", spaceId).Find(communities).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityTable"}
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) FetchRangeFromSpace(spaceId uint, count uint) (communities []model.Community, err error) {
	if cr.db.Where("space_id = ?", spaceId).Last(communities, count).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: CommunityTable"}
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) Update(newCommunity *model.Community) error {
	if err := cr.db.Save(newCommunity).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}

func (cr *CommunityRepositoryImpl) Delete(id uint) error {
	target := model.Community{}
	target.ID = id

	if err := cr.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
