package dal

import (
	"backend/app/rpc/user/biz/model"
	"backend/app/rpc/user/conf"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
	cache "github.com/mgtv-tech/jetcache-go"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func initConf() conf.Config {
	prefix := "../../conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	c := new(conf.Config)
	err = yaml.Unmarshal(content, c)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(c); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}
	c.Env = GetEnv()
	pretty.Printf("%+v\n", c)
	return *c
}

func TestUserDalImpl_UserDal(t *testing.T) {
	c := initConf()
	Init(c)
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestUserDalImpl_CreateAndGetUser",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: fmt.Sprintf("test_%d", time.Now().Unix()),
					Password: "admin",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDalInstance

			// 1. 创建用户
			if err := s.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			newUser, err := s.GetUserByID(tt.args.ctx, tt.args.user.ID)
			if err != nil {
				t.Errorf("getuserByid() error = %v", err)
			}
			if newUser.Username != tt.args.user.Username {
				t.Errorf("getuserByid() username error = %v", err)
			}

			if newUser.Password != tt.args.user.Password {
				t.Errorf("getuserByid() password error = %v", err)
			}

			// 2. 验证是否重复用户
			if err := s.CreateUser(tt.args.ctx, tt.args.user); err != nil {
				t.Logf("验证重复键成功")
			}

			// 3. 验证ExistUserByUserName 的存在条件是否正常
			if f, err := s.ExistUserByUserName(tt.args.ctx, tt.args.user.Username); err != nil {
				t.Errorf("ExistUserByUserName() error = %v", err)
			} else {
				if !f {
					// 要求是存在的
					t.Errorf("ExistUserByUserName() 存在验证失败")
				} else {
					t.Logf("ExistUserByUserName() 存在验证通过")
				}
			}

			// 4. 验证GetUserByUserName 的不存在条件是否正常
			if u, err := s.ExistUserByUserName(tt.args.ctx, fmt.Sprintf("test_%d", time.Now().Unix())); err != nil {
				t.Errorf("ExistUserByUserName() error = %v", err)
			} else {
				if u != false {
					t.Errorf("ExistUserByUserName 不存在验证失败")
				} else {
					t.Logf("ExistUserByUserName 不存在验证通过")
				}
			}

			// 5. 验证GetUserByUserName 存在的用户名
			if u, err := s.GetUserByUserName(tt.args.ctx, tt.args.user.Username); err != nil {
				t.Errorf("GetUserByUserName() error = %v", err)
			} else {
				if u.Username == tt.args.user.Username && u.Password == tt.args.user.Password {
					t.Logf("GetUserByUserName() 验证用户存在通过")
				} else {
					t.Errorf("GetUserByUserName() 验证用户不存在失败")
				}
			}

			// 6. 验证GetUserByUserName 不存在的用户名
			if _, err := s.GetUserByUserName(tt.args.ctx, fmt.Sprintf("test_%d", time.Now().Unix())); err != nil {
				if err == gorm.ErrRecordNotFound {
					t.Logf("GetUserByUserName() 验证用户不存在通过")
				} else {
					t.Errorf("GetUserByUserName() 验证用户不存在失败")
				}
			}
		})
	}
}

func TestUserDalImpl_DeleteUserByID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			if err := s.DeleteUserByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDalImpl_ExistUserByUserName(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.ExistUserByUserName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExistUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExistUserByUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetFollowersByUserID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx    context.Context
		userID int64
		total  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int64]*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetFollowersByUserID(tt.args.ctx, tt.args.userID, tt.args.total, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFollowersByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFollowersByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetFollowingsByUserID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx    context.Context
		userID int64
		total  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int64]*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetFollowingsByUserID(tt.args.ctx, tt.args.userID, tt.args.total, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFollowingsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFollowingsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetFriendsByUserID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx    context.Context
		userID int64
		total  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int64]*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetFriendsByUserID(tt.args.ctx, tt.args.userID, tt.args.total, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetUserByEmail(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetUserByID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetUserByPhone(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx   context.Context
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetUserByPhone(tt.args.ctx, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetUserByUserName(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx      context.Context
		userName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetUserByUserName(tt.args.ctx, tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_GetUsersByIDs(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx context.Context
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int64]*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			got, err := s.GetUsersByIDs(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersByIDs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDalImpl_UpdateUserByID(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		ctx  context.Context
		id   int64
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			if err := s.UpdateUserByID(tt.args.ctx, tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDalImpl_delayedDoubleDelete(t *testing.T) {
	type fields struct {
		cache         cache.Cache
		db            *gorm.DB
		userByIdCache *cache.T[int64, *model.User]
	}
	type args struct {
		cacheKey     string
		maxRetries   int
		initialDelay time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserDalImpl{
				cache:         tt.fields.cache,
				db:            tt.fields.db,
				userByIdCache: tt.fields.userByIdCache,
			}
			s.delayedDoubleDelete(tt.args.cacheKey, tt.args.maxRetries, tt.args.initialDelay)
		})
	}
}
