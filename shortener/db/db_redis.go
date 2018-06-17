package db

import (
	"sync/atomic"

	"github.com/go-redis/redis"
	"github.com/wrfly/yasuser/types"
)

type redisDB struct {
	cli    *redis.Client
	length *int64
}

func newRedisDB(redisAddr string) (*redisDB, error) {
	opts, err := redis.ParseURL(redisAddr)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opts)

	if err := cli.Ping().Err(); err != nil {
		return nil, err
	}

	return &redisDB{
		cli:    cli,
		length: &initLength,
	}, nil
}

func (r *redisDB) Close() error {
	return r.cli.Close()
}

func (r *redisDB) Len() int64 {
	return atomic.LoadInt64(r.length)
}

func (r *redisDB) SetShort(md5sum, shortURL string) error {
	if err := r.set(md5sum, shortURL); err != nil {
		return err
	}
	atomic.AddInt64(r.length, 1)
	return nil
}

func (r *redisDB) GetShort(md5sum string) (short string, err error) {
	return r.get(md5sum)
}

func (r *redisDB) SetLong(shortURL, longURL string) error {
	return r.set(shortURL, longURL)
}

func (r *redisDB) GetLong(md5sum string) (long string, err error) {
	return r.get(md5sum)
}

func (r *redisDB) set(key, value string) error {
	return r.cli.Set(key, value, -1).Err()
}

func (r *redisDB) get(key string) (value string, err error) {
	stringCmd := r.cli.Get(key)
	if err := stringCmd.Err(); err != nil {
		if err == redis.Nil {
			return "", types.ErrNotFound
		}
	}
	return stringCmd.String(), nil
}
