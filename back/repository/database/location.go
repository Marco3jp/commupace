package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type LocationRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) repository.LocationRepository {
	return &LocationRepositoryImpl{db: db}
}

func (lr *LocationRepositoryImpl) Add(newLocation *model.Location) (id uint, err error) {
	if !lr.db.NewRecord(*newLocation) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := lr.db.Create(newLocation).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

	return newLocation.ID, nil
}

func (lr *LocationRepositoryImpl) FindOne(id uint) (location *model.Location, err error) {
	if lr.db.First(location, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: ManagerAccountTable"}
	}

	return location, nil
}

// zoomLevelはGoogle Mapsと同様で、1 - 世界、20 - 建物となっている。
// 180をzoomLevel回半分にして、coordinatesに対して加減算することで両角の座標を求める。
// TODO: Module並の責務をRepositoryに置いてしまったので解体する, 暫定的に全件取得を用意して、パフォーマンスを見つつZoomLevelのバリデーションでうまくコントロールすることも考えられる
func (lr *LocationRepositoryImpl) FetchRangeFromCoordinates(coordinates model.Coordinates, zoomLevel uint, count uint) (locations []model.Location, err error) {
	searchRange := float64(180)
	for i := uint(0); i < zoomLevel; i++ {
		searchRange /= 2
	}

	westTop := model.Coordinates{}
	// TODO: 経度はまだしも緯度をまたぐ地域は早めに対処する
	if coordinates.Latitude < searchRange {
		westTop.Latitude = 0
	} else {
		westTop.Latitude = coordinates.Latitude - searchRange
	}

	if coordinates.Longitude < searchRange {
		westTop.Longitude = 0
	} else {
		westTop.Longitude = coordinates.Longitude - searchRange
	}

	eastBottom := model.Coordinates{}
	eastBottom.Latitude = coordinates.Latitude + searchRange
	eastBottom.Longitude = coordinates.Longitude + searchRange

	if lr.db.
		Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?", westTop.Latitude, eastBottom.Latitude, westTop.Longitude, eastBottom.Longitude).
		Last(locations, count).
		RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: LocationTable"}
	}

	return locations, nil
}

func (lr *LocationRepositoryImpl) Update(newLocation *model.Location) error {
	if err := lr.db.Save(newLocation).Error; err != nil {
		return &repository.IOError{err.Error()}
	}
	return nil
}

func (lr *LocationRepositoryImpl) Delete(id uint) error {
	target := model.Location{}
	target.ID = id

	if err := lr.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
