# cache 缓存

## 示例

```go

import (
    "fmt"

	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/cache"
)

// 持久化一个 key
cache.Forever("key", "value")

// 保存一个任意值
// 例如：保存一个用户对象实例
user := model.User{
	ID: 1,
	Name: "alex",
}
cache.Set("obj", user, 0)

// 取出一个对象实例
var user model.User
cache.GetObject("obj", &user)
fmt.Println(user)

// 情况所有缓存
// ⚠️ 注意：这是一个危险操作，会清空 cache 所在的 redis 数据库中所有的数据
cache.Flush()

```