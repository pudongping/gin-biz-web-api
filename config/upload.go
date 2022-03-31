// 上传文件时的配置
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.upload", func() map[string]interface{} {
		return map[string]interface{}{

			// 文件服务器的相对路径
			// 访问文件时，`http://localhost:3000/static` 即可访问到 `/public/uploads` 目录中
			// eg：
			// 文件保存路径为："public/uploads/avatar/2022/03/19/PqOem-6512bd43d9caa6e02c990b0a82652dca-20220319235446.jpg"
			// 那么访问连接地址为："http://localhost:3000/static/avatar/2022/03/19/PqOem-6512bd43d9caa6e02c990b0a82652dca-20220319235446.jpg"
			"static_fs_relative_path": config.Get("Upload.StaticFSRelativePath", "/static"),

			// 上传文件的最终保存目录
			"save_path": config.Get("Upload.SavePath", "public/uploads"),

			// 上传文件所允许的最大空间大小（单位：MB）
			"max_size": config.Get("Upload.MaxSize", 5),

			// 上传图片时的配置项
			// 其实上传图片也是文件上传的一种类型，但是也有可能是需要对图片进行裁剪之类的操作的，故另起了一个配置
			"image": map[string]interface{}{
				// 上传图片时，允许上传的最大空间大小，不设置时，会使用 upload.max_size 配置
				"max_size": config.Get("Upload.ImageMaxSize"),

				// 上传图片时，允许的图片后缀
				"allow_suffix": config.Get("Upload.ImageAllowSuffix", []string{".jpg", ".jpeg", ".png"}),
			},

			// 上传文件时的配置项
			"file": map[string]interface{}{
				// 上传文件时，允许上传的最大空间大小，不设置时，会使用 upload.max_size 配置
				"max_size": config.Get("Upload.FileMaxSize", 20),

				// 上传文件时，允许的文件后缀（只要存在 `*` 表示不对文件后缀做限制，即允许所有的文件上传）
				"allow_suffix": config.Get("Upload.FileAllowSuffix", []string{"*", ".ppt"}),
			},
		}
	})
}
