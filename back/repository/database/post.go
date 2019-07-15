package database

import (
	"github.com/jinzhu/gorm"
	".."
	"../../model"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (pr *PostRepositoryImpl) Add(newPost *model.Post) (id uint, err error) {
	if !pr.db.NewRecord(*newPost) {
		return 0, &repository.IDError{err.Error()}
	}

	if err := pr.db.Create(newPost).Error; err != nil {
		return 0, &repository.IOError{err.Error()}
	}

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

func (pr *PostRepositoryImpl) FetchRangeFromCommunityId(communityId uint, count uint) (posts []model.Post, err error) {
	pr.db.Where("community_id = ?", communityId).Order("post_number desc").Limit(count).Find(&posts)
	if len(posts) == 0 {
		return nil, &repository.NotFoundRecordError{"Action: postTable"}
	}

	return posts, nil
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
