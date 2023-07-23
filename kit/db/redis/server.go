package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"chatRobot/kit/tools/jsonx"
)

// Proxy 调用方定义的接口(只需要定义使用到的方法)
type Proxy interface {
	Ping() error                                                      // 检查连接是否正常
	Set(key string, value any, expiration time.Duration) error        // 设置
	Get(key string) (any, error)                                      // 获取
	Del(keys ...string) error                                         // 删除
	Incr(key string) (int64, error)                                   // 自增
	Expire(key string, expiration time.Duration) error                // 设置过期时间
	Exists(key string) bool                                           // 判断key是否存在
	SetInt64(key string, value int64, expiration time.Duration) error // 设置int64类型的值
	GetInt64(key string) (int64, error)                               // 获取int64类型的值
	SetOrIncr(key string, expiration time.Duration) (int64, error)    // 设置或者自增
	OutOfLimit(key string, limit int64) bool                          // 判断是否超过限制
}

type ClientProxy struct {
	Client Proxy
}

// New 创建 Redis 服务
func New(addr, password string, db int) *ClientProxy {
	return &ClientProxy{
		Client: newClient(addr, password, db),
	}
}

// Client 具体实现的结构体
type Client struct {
	client *redis.Client
}

func newClient(addr, password string, db int) *Client {
	return &Client{
		client: redis.NewClient(
			&redis.Options{
				Addr:     addr,
				Password: password,
				DB:       db,
			},
		),
	}
}

var (
	ctx = context.Background()
)

// Close 关闭 Redis 客户端连接
func (c *Client) Close() error {
	return c.client.Close()
}

// Ping 检查 Redis 连接是否正常
func (c *Client) Ping() error {
	_, err := c.client.Ping(ctx).Result()
	return err
}

// Set 序列化并存储一个键值对
func (c *Client) Set(key string, value any, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取一个键对应的值，并反序列化为指定的类型
func (c *Client) Get(key string) (any, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

// SetObject 序列化并设置一个对象类型的键值对，并指定过期时间（单位：秒）
func (c *Client) SetObject(key string, value any, expiration int) error {
	return c.client.Set(ctx, key, jsonx.MarshalString(value), time.Duration(expiration)*time.Second).Err()
}

// GetObject 获取一个对象类型的键值对的值，并反序列化为指定类型
func (c *Client) GetObject(key string, value any) error {
	bytes, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	jsonx.UnmarshalString(string(bytes), value)
	return nil
}

// Del 删除一个或多个键值对
func (c *Client) Del(keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// Incr 对一个键的值进行自增操作
func (c *Client) Incr(key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// Decr 对一个键的值进行自减操作
func (c *Client) Decr(key string) (int64, error) {
	return c.client.Decr(ctx, key).Result()
}

// Expire 设置一个键的过期时间（单位：秒）
func (c *Client) Expire(key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

// Exists 判断一个键是否存在
func (c *Client) Exists(key string) bool {
	return c.client.Exists(ctx, key).Val() == 1 // 1表示存在，0表示不存在
}

// SetInt64 序列化并设置一个 int64 类型的键值对，并指定过期时间（单位：秒）
func (c *Client) SetInt64(key string, value int64, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// GetInt64 获取一个 int64 类型的键值对的值
func (c *Client) GetInt64(key string) (int64, error) {
	return c.client.Get(ctx, key).Int64()
}

// SetOrIncr 设置或者自增
func (c *Client) SetOrIncr(key string, expiration time.Duration) (int64, error) {
	// 如果Key存在则自增
	if c.Exists(key) {
		return c.Incr(key)
	}
	// 如果Key不存在则设置Key的值为1
	_ = c.SetInt64(key, 1, expiration)
	return 1, nil
}

// OutOfLimit 判断一个键超出限制, 如果超出限制则返回 true, 否则返回 false
func (c *Client) OutOfLimit(key string, limit int64) bool {
	if c.Exists(key) {
		i, _ := c.GetInt64(key)
		if i >= limit {
			return true
		}
	}
	return false
}
