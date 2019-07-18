package database

import (
	"github.com/jinzhu/gorm"
	"github.com/Marco3jp/commupace/back/repository"
	"github.com/Marco3jp/commupace/back/model"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (pr *PostRepositoryImpl) Add(newPost *model.Post) (id uint, err error) {
	tx := pr.db.Begin()

	var lastPost model.Post
	postNum := uint(0)

	if !tx.Where("community_id = ?", newPost.CommunityID).Order("post_number desc").Limit(1).Find(&lastPost).RecordNotFound() {
		postNum = lastPost.PostNumber
	}

	newPost.PostNumber = postNum + 1

	if !tx.NewRecord(*newPost) {
		tx.Rollback()
		return 0, &repository.IDError{err.Error()}
	}

	if err := tx.Create(newPost).Error; err != nil {
		tx.Rollback()
		return 0, &repository.IOError{err.Error()}
	}

	tx.Commit()
	return newPost.ID, nil
}

func (pr *PostRepositoryImpl) FindOne(id uint) (post *model.Post, err error) {
	if pr.db.First(post, id).RecordNotFound() {
		return nil, &repository.NotFoundRecordError{"Action: postTable"}
	}

	return post, nil
}

func (pr *PostRepositoryImpl) FindFromCommunityId(communityId uint) (posts []model.Post, err error) {
	pr.db.Where("community_id = ?", communityId).Find(&posts)
	if len(posts) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: postTable"}
	}

	return posts, nil
}

// TODO: Threadが実装されたら追加
// func (pr *PostRepositoryImpl) FindFromThread(threadId uint) (posts []model.Post, err error){}

func (pr *PostRepositoryImpl) FetchRangeFromCommunityId(communityId uint, count uint) (result []model.PostData, err error) {
	// pr.db.Where("community_id = ?", communityId).Order("post_number desc").Limit(count).Find(&posts)
	result = make([]model.PostData, 0)
	pr.db.Table("posts").
		Select("posts.id, posts.created_at, posts.updated_at, posts.community_account_id, posts.thread_id, posts.post_text, posts.post_number, posts.post_type, posts.post_path, posts.community_id, community_accounts.id, community_accounts.display_id, community_accounts.display_name, community_accounts.icon, community_accounts.status").
		Joins("inner join community_accounts on posts.community_id = ? and posts.community_account_id = community_accounts.id", communityId).
		Order("post_number desc").
		Limit(count).
		Scan(&result)

	if len(result) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: postTable"}
	}

	return result, nil
}

// TODO: Threadが実装されたら追加
// func (pr *PostRepositoryImpl) FetchRangeFromThread(threadId uint, count uint) (posts []model.Post, err error){}

func (pr *PostRepositoryImpl) Update(newPost *model.Post) error {
	if err := pr.db.Save(newPost).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}

func (pr *PostRepositoryImpl) Delete(id uint) error {
	target := model.Post{}
	target.ID = id

	if err := pr.db.Delete(&target).Error; err != nil {
		return &repository.IOError{err.Error()}
	}

	return nil
}
