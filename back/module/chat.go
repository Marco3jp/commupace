package module

import (
	"../repository"
	"../model"
)

type ChatModuleImpl struct {
	PostRepo repository.PostRepository
}

func NewChatModule(pr repository.PostRepository) ChatModule {
	return &ChatModuleImpl{PostRepo: pr}
}

// PostTypeという謎のカラムとPostNumberが空のPost構造体を返します
func (cm *ChatModuleImpl) CreatePost(communityAccountId uint, communityId uint, postText string, postPath string) model.Post {
	return model.Post{
		CommunityAccountID: communityAccountId,
		CommunityID:        communityId,
		PostText:           postText,
		PostPath:           postPath,
	}
}

func (cm *ChatModuleImpl) Post(post model.Post) error {
	_, err := cm.PostRepo.Add(&post)
	if err != nil {
		return err
	}
	return nil
}

func (cm *ChatModuleImpl) GetPosts(communityId uint, count uint) (posts []model.PostData, err error) {
	posts, err = cm.PostRepo.FetchRangeFromCommunityId(communityId, communityId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
