# cache 缓存

> 目前使用 redis 作为缓存驱动，更多详情请见：`pkg/cache/redis_driver.go`  
> 
> 缓存 key 的前缀为：项目名称加 `:cache:` ，比如：`gin-biz-web-api:cache:users`

## 持久化值

```go

// 持久化一个 key
cache.Forever("key", "value")

```

## 保存任意值

```go

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


// 保存所有的用户信息
var users []model.User
database.DB.Find(&users)  // 从数据库中取出所有的用户信息
// 将所有的用户信息保存到缓存中
cache.Set("users", users, 0)
// 打印出缓存里面的所有用户信息
fmt.Println(cache.Get("users"))

// 可以通过以下示例代码反序列化出来
cacheUsers := cache.Get("users")

var cacheAllUsers []model.User
for _, v := range cacheUsers.([]interface{}) {
    var userItem model.User
    
    vv, _ := json.Marshal(v)
    _ = json.Unmarshal(vv, &userItem)
    
    cacheAllUsers = append(cacheAllUsers, userItem)
}
fmt.Printf("从缓存中取出来的所有用户数据为： %#v \n", cacheAllUsers)

```

## 删除

```go

// 删除指定的 key
cache.Forget("users")

// 清空所有缓存
// ⚠️ 注意：这是一个危险操作，会清空 cache 所在的 redis 数据库中所有的数据
// 比如说 cache 连接的是 redis db2 那么则会清空 db2 中所有的数据
cache.Flush()

```