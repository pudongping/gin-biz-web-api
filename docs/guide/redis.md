# redis 缓存

## 引入

```go
import (
"context"

"gin-biz-web-api/pkg/redis"
)

```

## 简单使用

```go

// 设置一个 key
redis.Set("name", "alex", 0)

// 取一个值
name := redis.Get("name")

// 递增
// 默认步长为 1
redis.Decrement("num")
// 指定步长
redis.Decrement("num", 3)
// 指定数据库组递增数据，比如这里采用 config/redis.go 配置文件中的 `cache` 组连接
redis.Decrement("num", 5, "cache")

// 递减
redis.Increment("num")
redis.Increment("num", 3)
redis.Increment("num", 5, "cache")

// 清空数据库中所有的数据
redis.FlushDB()

```

## 指定客户端实例

```go

// 获取一个 redis 客户端实例对象，默认获取 config/redis.go 配置文件中的 `default` 组连接
redis.Instance()

// 获取一个指定组连接客户端实例对象，比如这里获取的是 `cache` 组对应的实例对象
redis.Instance("cache")

```

## 使用 redis 类库包提供的方法

```go

// 使用 redis 类库包的方法
var ctx = context.Background()
// 设置一个值
redis.Instance().Client.Set(ctx, "hello", "alex", 0)
// 获取一个值
val, err := redis.Instance().Client.Get(ctx, "hello").Result()

```