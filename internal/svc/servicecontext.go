package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map

	Sequence sequence.Sequence // sequence

	ShortUrlBlackList map[string]struct{}

	// bloom filter
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	// 把配置文件中配置的黑名单加载到map，方便后续判断
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}

	// 初始化布隆过滤器
	// 初始化 redisBitSet
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	filter := bloom.New(store, "bloom_filter", 20*(1<<20))
	// 加载已有的短链接数据

	return &ServiceContext{
		Config:            c,
		ShortUrlModel:     model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:          sequence.NewMySQL(c.Sequence.DSN), // sequence
		ShortUrlBlackList: m,                                 // 短链接黑名单map
		Filter:            filter,
	}
}
