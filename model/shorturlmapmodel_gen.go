// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shortUrlMapFieldNames          = builder.RawFieldNames(&ShortUrlMap{})
	shortUrlMapRows                = strings.Join(shortUrlMapFieldNames, ",")
	shortUrlMapRowsExpectAutoSet   = strings.Join(stringx.Remove(shortUrlMapFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	shortUrlMapRowsWithPlaceHolder = strings.Join(stringx.Remove(shortUrlMapFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheShortenerShortUrlMapIdPrefix   = "cache:shortener:shortUrlMap:id:"
	cacheShortenerShortUrlMapMd5Prefix  = "cache:shortener:shortUrlMap:md5:"
	cacheShortenerShortUrlMapSurlPrefix = "cache:shortener:shortUrlMap:surl:"
)

type (
	shortUrlMapModel interface {
		Insert(ctx context.Context, data *ShortUrlMap) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ShortUrlMap, error)
		FindOneByMd5(ctx context.Context, md5 sql.NullString) (*ShortUrlMap, error)
		FindOneBySurl(ctx context.Context, surl sql.NullString) (*ShortUrlMap, error)
		Update(ctx context.Context, data *ShortUrlMap) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShortUrlMapModel struct {
		sqlc.CachedConn
		table string
	}

	ShortUrlMap struct {
		Id       int64          `db:"id"`        // 主键
		CreateAt time.Time      `db:"create_at"` // 创建时间
		CreateBy string         `db:"create_by"` // 创建者
		IsDel    int64          `db:"is_del"`    // 是否删除：0正常1删除
		Lurl     sql.NullString `db:"lurl"`      // 长链接
		Md5      sql.NullString `db:"md5"`       // 长链接MD5
		Surl     sql.NullString `db:"surl"`      // 短链接
	}
)

func newShortUrlMapModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultShortUrlMapModel {
	return &defaultShortUrlMapModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`short_url_map`",
	}
}

func (m *defaultShortUrlMapModel) withSession(session sqlx.Session) *defaultShortUrlMapModel {
	return &defaultShortUrlMapModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`short_url_map`",
	}
}

func (m *defaultShortUrlMapModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	shortenerShortUrlMapIdKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapIdPrefix, id)
	shortenerShortUrlMapMd5Key := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapMd5Prefix, data.Md5)
	shortenerShortUrlMapSurlKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapSurlPrefix, data.Surl)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, shortenerShortUrlMapIdKey, shortenerShortUrlMapMd5Key, shortenerShortUrlMapSurlKey)
	return err
}

func (m *defaultShortUrlMapModel) FindOne(ctx context.Context, id int64) (*ShortUrlMap, error) {
	shortenerShortUrlMapIdKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapIdPrefix, id)
	var resp ShortUrlMap
	err := m.QueryRowCtx(ctx, &resp, shortenerShortUrlMapIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shortUrlMapRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) FindOneByMd5(ctx context.Context, md5 sql.NullString) (*ShortUrlMap, error) {
	shortenerShortUrlMapMd5Key := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapMd5Prefix, md5)
	var resp ShortUrlMap
	err := m.QueryRowIndexCtx(ctx, &resp, shortenerShortUrlMapMd5Key, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `md5` = ? limit 1", shortUrlMapRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, md5); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) FindOneBySurl(ctx context.Context, surl sql.NullString) (*ShortUrlMap, error) {
	shortenerShortUrlMapSurlKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapSurlPrefix, surl)
	var resp ShortUrlMap
	err := m.QueryRowIndexCtx(ctx, &resp, shortenerShortUrlMapSurlKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `surl` = ? limit 1", shortUrlMapRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, surl); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) Insert(ctx context.Context, data *ShortUrlMap) (sql.Result, error) {
	shortenerShortUrlMapIdKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapIdPrefix, data.Id)
	shortenerShortUrlMapMd5Key := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapMd5Prefix, data.Md5)
	shortenerShortUrlMapSurlKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapSurlPrefix, data.Surl)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, shortUrlMapRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateBy, data.IsDel, data.Lurl, data.Md5, data.Surl)
	}, shortenerShortUrlMapIdKey, shortenerShortUrlMapMd5Key, shortenerShortUrlMapSurlKey)
	return ret, err
}

func (m *defaultShortUrlMapModel) Update(ctx context.Context, newData *ShortUrlMap) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	shortenerShortUrlMapIdKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapIdPrefix, data.Id)
	shortenerShortUrlMapMd5Key := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapMd5Prefix, data.Md5)
	shortenerShortUrlMapSurlKey := fmt.Sprintf("%s%v", cacheShortenerShortUrlMapSurlPrefix, data.Surl)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shortUrlMapRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.CreateBy, newData.IsDel, newData.Lurl, newData.Md5, newData.Surl, newData.Id)
	}, shortenerShortUrlMapIdKey, shortenerShortUrlMapMd5Key, shortenerShortUrlMapSurlKey)
	return err
}

func (m *defaultShortUrlMapModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheShortenerShortUrlMapIdPrefix, primary)
}

func (m *defaultShortUrlMapModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shortUrlMapRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShortUrlMapModel) tableName() string {
	return m.table
}
