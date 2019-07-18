package database

import (
	"github.com/jinzhu/gorm"
	"github.com/Marco3jp/commupace/back/repository"
	"github.com/Marco3jp/commupace/back/model"
)

type SpaceRepositoryImpl struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) repository.SpaceRepository {
	return &SpaceRepositoryImpl{db: db}
}

func (sr *SpaceRepositoryImpl) Add(newSpace *model.Space) (id uint, err error) {
	if !sr.db.NewRecord(*newSpace) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := sr.db.Create(newSpace).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newSpace.ID, nil
}

func (sr *SpaceRepositoryImpl) FindOne(id uint) (space *model.Space, err error) {
	if sr.db.First(space, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: SpaceTable"}
	}

	return space, nil
}

func (sr *SpaceRepositoryImpl) FindFromLocation(locationId uint) (spaces []model.Space, err error) {
	sr.db.Where("location_id = ?", locationId).Find(&spaces)
	if len(spaces) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: SpaceTable"}
	}

	return spaces, nil
}

func (sr *SpaceRepositoryImpl) FetchRangeFromLocation(locationId uint, count uint) (spaces []model.Space, err error) {
	sr.db.Where("location_id = ?", locationId).Find(&spaces).Order("id desc").Limit(count)
	if len(spaces) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: SpaceTable"}
	}

	return spaces, nil
}

func (sr *SpaceRepositoryImpl) FetchRangeFromLocations(locationIds []uint, count uint) (spaces []model.Space, err error) {
	sr.db.Where("location_id IN (?)", locationIds).Find(&spaces).Order("id desc").Limit(count)
	if len(spaces) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: SpaceTable"}
	}

	return spaces, nil
}

func (sr *SpaceRepositoryImpl) Update(newSpace *model.Space) error {
	if err := sr.db.Save(newSpace).Error; err != nil {
		return &repository.IOError{err.Error()}
	}
	return nil
}

func (sr *SpaceRepositoryImpl) Delete(id uint) error {
	target := model.Space{}
	target.ID = id

	if err := sr.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
