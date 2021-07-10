package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AuthInfoMgr struct {
	*_BaseMgr
}

// AuthInfoMgr open func
func AuthInfoMgr(db *gorm.DB) *_AuthInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AuthInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AuthInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("auth_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AuthInfoMgr) GetTableName() string {
	return "auth_info"
}

// Get 获取
func (obj *_AuthInfoMgr) Get() (result AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AuthInfoMgr) Gets() (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AuthInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithClientName client_name获取 用户名称
func (obj *_AuthInfoMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

// WithClientSecret client_secret获取 客户端秘钥
func (obj *_AuthInfoMgr) WithClientSecret(clientSecret string) Option {
	return optionFunc(func(o *options) { o.query["client_secret"] = clientSecret })
}

// WithClientID client_id获取 客户端名称
func (obj *_AuthInfoMgr) WithClientID(clientID string) Option {
	return optionFunc(func(o *options) { o.query["client_id"] = clientID })
}

// WithCallbackURL callback_url获取 客户端回调地址
func (obj *_AuthInfoMgr) WithCallbackURL(callbackURL string) Option {
	return optionFunc(func(o *options) { o.query["callback_url"] = callbackURL })
}

// WithIsDelete is_delete获取 是否删除
func (obj *_AuthInfoMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithUpdateTime update_time获取
func (obj *_AuthInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateTime create_time获取
func (obj *_AuthInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_AuthInfoMgr) GetByOption(opts ...Option) (result AuthInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AuthInfoMgr) GetByOptions(opts ...Option) (results []*AuthInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AuthInfoMgr) GetFromID(id int64) (result AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AuthInfoMgr) GetBatchFromID(ids []int64) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromClientName 通过client_name获取内容 用户名称
func (obj *_AuthInfoMgr) GetFromClientName(clientName string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_name` = ?", clientName).Find(&results).Error

	return
}

// GetBatchFromClientName 批量查找 用户名称
func (obj *_AuthInfoMgr) GetBatchFromClientName(clientNames []string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_name` IN (?)", clientNames).Find(&results).Error

	return
}

// GetFromClientSecret 通过client_secret获取内容 客户端秘钥
func (obj *_AuthInfoMgr) GetFromClientSecret(clientSecret string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_secret` = ?", clientSecret).Find(&results).Error

	return
}

// GetBatchFromClientSecret 批量查找 客户端秘钥
func (obj *_AuthInfoMgr) GetBatchFromClientSecret(clientSecrets []string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_secret` IN (?)", clientSecrets).Find(&results).Error

	return
}

// GetFromClientID 通过client_id获取内容 客户端名称
func (obj *_AuthInfoMgr) GetFromClientID(clientID string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_id` = ?", clientID).Find(&results).Error

	return
}

// GetBatchFromClientID 批量查找 客户端名称
func (obj *_AuthInfoMgr) GetBatchFromClientID(clientIDs []string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`client_id` IN (?)", clientIDs).Find(&results).Error

	return
}

// GetFromCallbackURL 通过callback_url获取内容 客户端回调地址
func (obj *_AuthInfoMgr) GetFromCallbackURL(callbackURL string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`callback_url` = ?", callbackURL).Find(&results).Error

	return
}

// GetBatchFromCallbackURL 批量查找 客户端回调地址
func (obj *_AuthInfoMgr) GetBatchFromCallbackURL(callbackURLs []string) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`callback_url` IN (?)", callbackURLs).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 是否删除
func (obj *_AuthInfoMgr) GetFromIsDelete(isDelete bool) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_AuthInfoMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_AuthInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_AuthInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_AuthInfoMgr) GetFromCreateTime(createTime time.Time) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_AuthInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AuthInfoMgr) FetchByPrimaryKey(id int64) (result AuthInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where("`id` = ?", id).Find(&result).Error

	return
}
