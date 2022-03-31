// 分页相关配置
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.paginator", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认每页显示数
			"default_per_page": 20,

			// URL 中用以区分每页显示数的参数名
			"url_query_per_page": config.Get("Paginator.UrlQueryPerPage", "per_page"),

			// URL 中用以区分当前页码的参数名
			"url_query_page": config.Get("Paginator.UrlQueryPage", "page"),

			// URL 中用以分页时排序的参数名
			// eg：`id,desc|name,asc` 表示：先按照 id 字段倒序排列，然后按照 name 字段正序排列
			"url_query_order_by": config.Get("Paginator.UrlQueryOrderBy", "order_by"),
		}
	})
}
