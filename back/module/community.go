package module

import (
	"github.com/Marco3jp/commupace/back/repository"
	"github.com/Marco3jp/commupace/back/model"
)

type CommunityModuleImpl struct {
	LocationRepo  repository.LocationRepository
	SpaceRepo     repository.SpaceRepository
	CommunityRepo repository.CommunityRepository
}

func NewCommunityModule(lr repository.LocationRepository, sr repository.SpaceRepository, cr repository.CommunityRepository) CommunityModule {
	return &CommunityModuleImpl{LocationRepo: lr, SpaceRepo: sr, CommunityRepo: cr}
}

func (cm *CommunityModuleImpl) SearchCommunityFromCoordinates(coordinates model.Coordinates, zoomLevel uint) (communityList []model.Community, err error) {
	const MAX_SEARCH_COUNT = 10

	locationList, err := cm.LocationRepo.FetchRangeFromCoordinates(coordinates, zoomLevel, MAX_SEARCH_COUNT)
	if err != nil {
		return nil, err
	}

	var locationIdList []uint

	for e := range locationList {
		locationIdList = append(locationIdList, locationList[e].ID)
	}

	spaceList, err := cm.SpaceRepo.FetchRangeFromLocations(locationIdList, MAX_SEARCH_COUNT)
	if err != nil {
		return nil, err
	}

	var spaceIdList []uint

	for e := range spaceList {
		spaceIdList = append(spaceIdList, spaceList[e].ID)
	}

	communityList, err = cm.CommunityRepo.FetchRangeFromSpaces(spaceIdList, MAX_SEARCH_COUNT)
	if err != nil {
		return nil, err
	}

	return communityList, nil
}
