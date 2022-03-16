// Package config 负责配置信息
package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/file"
	"gin-biz-web-api/pkg/helper"
)

// vp viper 库实例
var vp *viper.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 在动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 初始化 Viper 库
	vp = viper.New()

	// 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//              "props", "prop", "env", "dotenv"
	vp.SetConfigType("yaml") // 设置配置文件的类型为 yaml

	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	// vp.SetEnvPrefix("app_env")
	// 读取环境变量（支持 flags）
	// vp.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// NewConfig 初始化配置信息
func NewConfig(env string, configs ...string) {

	fileName := fetchConfigFile(env) // 根据环境变量 env 获取对应的配置文件名称（不带文件后缀）
	vp.SetConfigName(fileName)       // 设置配置文件的名称为 config

	for _, config := range configs {
		if config != "" {

			configFile := config + fileName + ".yaml"
			if _, ok := file.IsExists(configFile); !ok {
				console.Exit("配置文件 [%s] 不存在！", configFile)
			}

			// 设置配置文件路径为相对路径，相对于 main.go 比如：vp.AddConfigPath("config/")
			vp.AddConfigPath(config)
		}
	}

	err := vp.ReadInConfig() // 读取配置文件中的内容
	if err != nil {
		panic(err)
	}

	loadConfig()
}

// fetchConfigFile 根据环境变量加载对应的配置文件
func fetchConfigFile(env string) string {
	// 如果没设置环境变量，那么默认使用【etc/config.yaml 配置文件】
	if env == "" {
		return "config"
	}

	return env + "_config"
}

// loadConfig 注册配置信息，动态增加配置文件
func loadConfig() {
	for name, fn := range ConfigFuncs {
		vp.Set(name, fn())
	}
}

// get 从 vp 中获取配置信息（仅供内部调用）
func get(path string, defaultValue ...interface{}) interface{} {
	// 配置文件中不存在的情况
	if !vp.IsSet(path) || helper.Empty(vp.Get(path)) {
		// 如果有默认参数，则使用默认参数
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		// 没有默认参数，则直接返回空
		return nil
	}

	return vp.Get(path)
}

// Get 供外部获取配置信息使用，支持默认值
func Get(path string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return get(path, defaultValue[0])
	}

	return get(path)
}

// Add 动态新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Instance 返回 viper 实例（但是不建议外部调用）
func Instance() *viper.Viper {
	return vp
}

// GetString 获取 string 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(get(path, defaultValue...))
}

// GetInt 获取 int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(get(path, defaultValue...))
}

// GetInt64 获取 int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(get(path, defaultValue...))
}

// GetUint 获取 uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(get(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(get(path, defaultValue...))
}

// GetBool 获取 bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(get(path, defaultValue...))
}

// GetStringMapString 获取 map 类型的配置信息
func GetStringMapString(path string) map[string]string {
	return vp.GetStringMapString(path)
}

// GetStringSlice 获取切片类型的配置信息
func GetStringSlice(path string) []string {
	return vp.GetStringSlice(path)
}
