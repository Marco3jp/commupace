package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
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
	if sr.db.Where("location_id = ?", locationId).Find(spaces).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: SpaceTable"}
	}

	return spaces, nil
}

func (sr *SpaceRepositoryImpl) FetchRangeFromLocation(locationId uint, count uint) (spaces []model.Space, err error) {
	if sr.db.Where("location_id = ?", locationId).Last(spaces, count).RecordNotFound() {
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
