package post

import (
	"cloud/dao"
	"cloud/initial"
	"context"

	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
)

// system_post 职位
// SystemPostCreate 创建数据
func SystemPostCreate(ctx context.Context, data dao.SystemPost) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_post`").Insert(data)
	if err != nil {
		return
	}
	res, err = db.Insert(ctx, query, args...)
	return
}

// SystemPostUpdate 更新数据
func SystemPostUpdate(ctx context.Context, systemPostId int64, data dao.SystemPost) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_post`").Where("`id`", systemPostId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemPostDelete 删除数据
func SystemPostDelete(ctx context.Context, systemPostId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 1
	query, args, err := builder.Table("`system_post`").Where("`id`", systemPostId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemPost 查询单条数据
func SystemPost(ctx context.Context, systemPostId int64) (res dao.SystemPost, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	query, args, err := builder.Table("`system_post`").Where("`id`", systemPostId).Row()
	if err != nil {
		return
	}
	err = db.QueryRow(ctx, query, args...).ToStruct(&res)
	return
}

// SystemPostRecover 恢复数据
func SystemPostRecover(ctx context.Context, systemPostId int64) (res int64, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Write()
	builder := sql.NewBuilder()
	data := make(map[string]any)
	data["deleted"] = 0
	query, args, err := builder.Table("`system_post`").Where("`id`", systemPostId).Update(data)
	if err != nil {
		return
	}
	res, err = db.Update(ctx, query, args...)
	return
}

// SystemPostList 查询列表数据
func SystemPostList(ctx context.Context, condition map[string]any) (res []dao.SystemPost, err error) {
	db := initial.Core.Store.LoadSQL("mysql").Read()
	builder := sql.NewBuilder()
	builder.Table("`system_post`")
	if val, ok := condition["systemTenantId"]; ok {
		builder.Where("`system_tenant_id`", val)
	}
	if val, ok := condition["deleted"]; ok {
		builder.Where("`deleted`", val)
	}
	if val, ok := condition["status"]; ok {
		builder.Where("`status`", val)
	}
	if val, ok := condition["name"]; ok {
		builder.Like("`name`", "%"+cast.ToString(val)+"%")
	}

	builder.OrderBy("`sort`", sql.ASC)
	builder.OrderBy("`id`", sql.DESC)
	query, args, err := builder.Rows()
	if err != nil {
		return
	}
	err = db.QueryRows(ctx, query, args...).ToStruct(&res)
	return
}
