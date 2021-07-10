package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserInfoMgr struct {
	*_BaseMgr
}

// UserInfoMgr open func
func UserInfoMgr(db *gorm.DB) *_UserInfoMgr {
	if db == nil {
		panic(fmt.Errorf("UserInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserInfoMgr) GetTableName() string {
	return "user_info"
}

// Get 获取
func (obj *_UserInfoMgr) Get() (result UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserInfoMgr) Gets() (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_UserInfoMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取
func (obj *_UserInfoMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithSalt salt获取
func (obj *_UserInfoMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithIsDelete is_delete获取 是否删除
func (obj *_UserInfoMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithUpdateTime update_time获取
func (obj *_UserInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateTime create_time获取
func (obj *_UserInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_UserInfoMgr) GetByOption(opts ...Option) (result UserInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserInfoMgr) GetByOptions(opts ...Option) (results []*UserInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserInfoMgr) GetFromID(id int64) (result UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserInfoMgr) GetBatchFromID(ids []int64) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_UserInfoMgr) GetFromUsername(username string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找
func (obj *_UserInfoMgr) GetBatchFromUsername(usernames []string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UserInfoMgr) GetFromPassword(password string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_UserInfoMgr) GetBatchFromPassword(passwords []string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容
func (obj *_UserInfoMgr) GetFromSalt(salt string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`salt` = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量查找
func (obj *_UserInfoMgr) GetBatchFromSalt(salts []string) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`salt` IN (?)", salts).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 是否删除
func (obj *_UserInfoMgr) GetFromIsDelete(isDelete bool) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_UserInfoMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_UserInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_UserInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_UserInfoMgr) GetFromCreateTime(createTime time.Time) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_UserInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserInfoMgr) FetchByPrimaryKey(id int64) (result UserInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfo{}).Where("`id` = ?", id).Find(&result).Error

	return
}
