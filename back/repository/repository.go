package repository

import "github.com/Marco3jp/commupace/back/model"

type ManagerAccountRepository interface {
	Add(newManagerAccount *model.ManagerAccount) (id uint, err error)
	FindOne(id uint) (managerAccount *model.ManagerAccount, err error)
	FindOneFromManagerAccountId(managerAccountId string) (managerAccount *model.ManagerAccount, err error)
	Update(newManagerAccount *model.ManagerAccount) error
	Delete(id uint) error
}

type AccessTokenRepository interface {
	Add(token string, managerAccountId string, exp int64)
	FindOne(token string) (managerAccountId string, exp int64, err error)
	Update(token string, managerAccountId string, exp int64)
	Delete(token string)
}

type RefreshTokenRepository interface {
	Add(token string, managerAccountId string, exp int64)
	FindOne(token string) (managerAccountId string, exp int64, err error)
	Update(token string, managerAccountId string, exp int64)
	Delete(token string)
}

type CommunityAccountRepository interface {
	Add(newCommunityAccount *model.CommunityAccount) (id uint, err error)
	FindOne(id uint) (communityAccount *model.CommunityAccount, err error)
	FindOneFromDisplayId(displayId string) (communityAccount *model.CommunityAccount, err error)
	FindFromCommunityAccountIds(communityAccountIds []uint) (communityAccounts []model.CommunityAccount, err error)
	FindFromManagerAccount(managerAccountId uint) (communityAccounts []model.CommunityAccount, err error)
	Update(newCommunityAccount *model.CommunityAccount) error
	Delete(id uint) error
}

type LocationRepository interface {
	Add(newLocation *model.Location) (id uint, err error)
	FindOne(id uint) (location *model.Location, err error)
	// 以下のメソッドは、引数の場所から近い場所をIDの新しい順でcount個返します(近い順ではない)
	FetchRangeFromCoordinates(coordinates model.Coordinates, zoomLevel uint, count uint) (locations []model.Location, err error)
	Update(newLocation *model.Location) error
	Delete(id uint) error
}

type SpaceRepository interface {
	Add(newSpace *model.Space) (id uint, err error)
	FindOne(id uint) (space *model.Space, err error)
	FindFromLocation(locationId uint) (spaces []model.Space, err error)
	// 以下のメソッドは、引数の場所にある空間を、IDの新しい順でcount個返します
	FetchRangeFromLocation(locationId uint, count uint) (spaces []model.Space, err error)
	// 以下のメソッドは、引数の場所たちにある空間を、IDの新しい順でcount個返します
	FetchRangeFromLocations(locationIds []uint, count uint) (spaces []model.Space, err error)
	Update(newSpace *model.Space) error
	Delete(id uint) error
}

type CommunityRepository interface {
	Add(newCommunity *model.Community) (id uint, err error)
	FindOne(id uint) (community *model.Community, err error)
	FindOneFromCommunityId(communityId string) (community *model.Community, err error)
	FindFromSpace(spaceId uint) (communities []model.Community, err error)
	FetchRangeFromSpace(spaceId uint, count uint) (communities []model.Community, err error)
	FetchRangeFromSpaces(spaceIds []uint, count uint) (communities []model.Community, err error)
	Update(newCommunity *model.Community) error
	Delete(id uint) error
}

type CommunityUserRepository interface {
	Add(newCommunityUser *model.CommunityUser) (id uint, err error)
	FindOne(id uint) (communityUser *model.CommunityUser, err error)
	FindFromCommunity(communityId uint) (communityUsers []model.CommunityUser, err error)
	FindFromCommunityAccount(communityAccountId uint) (communityUsers []model.CommunityUser, err error)
	FetchRangeFromCommunity(communityId uint, count uint) (communityUsers []model.CommunityUser, err error)
	FetchRangeFromCommunityAccount(communityAccountId uint, count uint) (communityUsers []model.CommunityUser, err error)
	Delete(id uint) error
}

type PostRepository interface {
	Add(newPost *model.Post) (id uint, err error)
	FindOne(id uint) (post *model.Post, err error)
	FindFromCommunityId(communityId uint) (posts []model.Post, err error)
	//FindFromThread(threadId uint) (posts []model.Post, err error) TODO: Threadが実装されたら追加
	FetchRangeFromCommunityId(communityId uint, count uint) (result []model.PostData, err error)
	//FetchRangeFromThread(threadId uint, count uint) (posts []model.Post, err error) TODO: Threadが実装されたら追加
	Update(newPost *model.Post) error
	Delete(id uint) error
}
