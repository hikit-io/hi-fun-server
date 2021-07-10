package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var globalIsRelated bool = true // 全局预加载

// prepare for other
type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

// SetCtx set context
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

// Ctx get context
func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

// Cancel cancel context
func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

// New new gorm.新gorm
func (obj *_BaseMgr) New() *gorm.DB {
	return obj.DB.Session(&gorm.Session{Context: obj.ctx})
}

type options struct {
	query map[string]interface{}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// OpenRelated 打开全局预加载
func OpenRelated() {
	globalIsRelated = true
}

// CloseRelated 关闭全局预加载
func CloseRelated() {
	globalIsRelated = true
}

// 自定义sql查询
type Condetion struct {
	list []*condetionInfo
}

// And a condition by and .and 一个条件
func (c *Condetion) And(column string, cases string, value ...interface{}) {
	c.list = append(c.list, &condetionInfo{
		andor:  "and",
		column: column, // 列名
		case_:  cases,  // 条件(and,or,in,>=,<=)
		value:  value,
	})
}

// Or a condition by or .or 一个条件
func (c *Condetion) Or(column string, cases string, value ...interface{}) {
	c.list = append(c.list, &condetionInfo{
		andor:  "or",
		column: column, // 列名
		case_:  cases,  // 条件(and,or,in,>=,<=)
		value:  value,
	})
}

func (c *Condetion) Get() (where string, out []interface{}) {
	firstAnd := -1
	for i := 0; i < len(c.list); i++ { // 查找第一个and
		if c.list[i].andor == "and" {
			where = fmt.Sprintf("`%v` %v ?", c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
			firstAnd = i
			break
		}
	}

	if firstAnd < 0 && len(c.list) > 0 { // 补刀
		where = fmt.Sprintf("`%v` %v ?", c.list[0].column, c.list[0].case_)
		out = append(out, c.list[0].value)
		firstAnd = 0
	}

	for i := 0; i < len(c.list); i++ { // 添加剩余的
		if firstAnd != i {
			where += fmt.Sprintf(" %v `%v` %v ?", c.list[i].andor, c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
		}
	}

	return
}

type condetionInfo struct {
	andor  string
	column string // 列名
	case_  string // 条件(in,>=,<=)
	value  interface{}
}
