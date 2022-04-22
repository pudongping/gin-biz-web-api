// 缓存相关的命令
// eg：go run main.go cache -h
package cmd

import (
	"github.com/spf13/cobra"

	"gin-biz-web-api/pkg/cache"
	"gin-biz-web-api/pkg/console"
)

// CacheCmd 缓存管理
var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

// cacheClearCmd 清空所有的缓存
var cacheClearCmd = &cobra.Command{
	Use:     "clear",
	Short:   "Clear all cache. Notice: this command is most dangerous! eg: cache clear",
	Example: "go run main.go cache clear",
	Long:    "Notice: this command is most dangerous!",
	Run:     runCacheClear,
}

// cacheForgetCmd 删除缓存中的某个 key
var cacheForgetCmd = &cobra.Command{
	Use:     "forget",
	Short:   "Delete cache key, eg: cache forget cache-key",
	Example: "go run main.go cache forget cache-key",
	Run:     runCacheForget,
}

// cacheGetCmd 获取缓存中的某个值
var cacheGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get cache value, eg: cache get cache-key",
	Example: "go run main.go cache get cache-key",
	Run:     runCacheGet,
}

// 需要删除的缓存 key
var cacheKey string

func init() {
	// 注册 cache 命令的子命令
	CacheCmd.AddCommand(cacheClearCmd)

	CacheCmd.AddCommand(cacheForgetCmd)
	cacheForgetCmd.Flags().StringVarP(&cacheKey, "key", "k", "", "Key of the cache")
	// 设置 key 参数为必须
	_ = cacheForgetCmd.MarkFlagRequired("key")

	CacheCmd.AddCommand(cacheGetCmd)
	cacheGetCmd.Flags().StringVarP(&cacheKey, "key", "k", "", "Key of the cache")
	// 设置 key 参数为必须
	_ = cacheGetCmd.MarkFlagRequired("key")

}

// runCacheClear 清空所有的缓存
func runCacheClear(cmd *cobra.Command, args []string) {
	console.Info("Clearing all cache now")
	cache.Flush()
	console.Success("Cache cleared.")
}

// runCacheForget 删除指定缓存 key
func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success("Cache key [%s] deleted.", cacheKey)
}

// runCacheGet 获取某个缓存值
func runCacheGet(cmd *cobra.Command, args []string) {
	data := cache.Get(cacheKey)
	console.Info("%v", data)
	console.Info("type of ==> %T", data)
}
