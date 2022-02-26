package config

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewConfig(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 设置配置文件的名称为 config

	for _, config := range configs {
		if config != "" {
			// 设置配置文件路径为相对路径，比如：vp.AddConfigPath("configs/")
			vp.AddConfigPath(config)
		}
	}

	vp.SetConfigType("yaml") // 设置配置文件的类型为 yaml
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp: vp}

	return s, nil
}
