package dal

import (
	"backend/app/rpc/user/biz/model"
	"context"
	"gorm.io/gorm"
)

func getUserByID(ctx context.Context, db *gorm.DB, id int64) (*model.User, error) {
	user := &model.User{}
	err := db.WithContext(ctx).Where("id = ?", id).First(user).Error
	return user, err
}

func existUserByID(ctx context.Context, db *gorm.DB, id int64) (bool, error) {
	var cnt int64
	if err := db.WithContext(ctx).Where("id = ?", id).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func existUserByPhone(ctx context.Context, db *gorm.DB, phone string) (bool, error) {
	var cnt int64
	if err := db.WithContext(ctx).Where("phone = ?", phone).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func existUserByEmail(ctx context.Context, db *gorm.DB, email string) (bool, error) {
	var cnt int64
	if err := db.WithContext(ctx).Where("email = ?", email).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func existUserByUserName(ctx context.Context, db *gorm.DB, name string) (bool, error) {
	var cnt int64
	if err := db.WithContext(ctx).Model(&model.User{}).Where("username = ?", name).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func saveUserByID(ctx context.Context, db *gorm.DB, id int64, user *model.User) error {
	return db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Save(user).Error
}

func updateUserByID(ctx context.Context, db *gorm.DB, id int64, user *model.User) error {
	return db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(user).Error
}

func deleteUserByID(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
			return err
		}
		return nil
	})

}

func getUserByUserName(ctx context.Context, db *gorm.DB, name string) (*model.User, error) {
	user := &model.User{}
	err := db.WithContext(ctx).Where("username = ?", name).First(user).Error
	return user, err
}

func getUserByEmail(ctx context.Context, db *gorm.DB, email string) (*model.User, error) {
	user := &model.User{}
	err := db.WithContext(ctx).Where("email = ?", email).First(user).Error
	return user, err
}

func createUser(ctx context.Context, db *gorm.DB, user *model.User) error {
	return db.WithContext(ctx).Create(user).Error
}

func getUserRelevantCountByID(ctx context.Context, db *gorm.DB, id int64) (*model.UserRelevantCount, error) {
	userRelevantCount := &model.UserRelevantCount{}
	err := db.WithContext(ctx).Where("user_id = ?", id).First(userRelevantCount).Error
	return userRelevantCount, err
}

func getUserRelation(ctx context.Context, db *gorm.DB, uid, rid int64) (*model.UserRelationship, error) {
	userRelation := &model.UserRelationship{}
	err := db.WithContext(ctx).Where("user_id = ? and related_user_id", uid, rid).First(userRelation).Error
	return userRelation, err
}

func createUserRelation(ctx context.Context, db *gorm.DB, relation *model.UserRelationship) error {
	return db.WithContext(ctx).Create(relation).Error
}
