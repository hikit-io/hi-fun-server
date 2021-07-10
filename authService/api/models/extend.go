package models

import "auth-service/api/caches"

// GetByOptionCache 功能选项模式获取
func (obj *_AuthInfoMgr) GetByOptionCache(opts ...Option) (result AuthInfo, err error) {
	var (
		options = options{
			query: make(map[string]interface{}, len(opts)),
		}
		put func(res AuthInfo)
	)
	for _, o := range opts {
		o.apply(&options)
	}
	id, ok := options.query[AuthInfoColumns.ClientID]
	if ok {
		result, put, err = caches.ClientCacheFunc().GetOneByClientId(obj.ctx, id.(int64))
		if err != nil {
			return result, err
		}
	}
	err = obj.DB.WithContext(obj.ctx).Model(AuthInfo{}).Where(options.query).Find(&result).Error
	go put(result)
	return
}
