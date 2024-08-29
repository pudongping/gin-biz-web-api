package example_ctrl

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/paginator"
	"gin-biz-web-api/pkg/responses"
)

type PagerController struct {
}

// Pager 这里演示一条稍微比较复杂的 sql 分页
// curl --location --request GET '0.0.0.0:3000/api/example/pager?aa=11&bb=22&order_by=a.item_id,desc|b.pid,asc&per_page=2&page=3'
// SELECT count(*) FROM item as a left join product as b on a.pid = b.pid left join product_data as c on b.pid = c.pid left join site as d on b.source_site = d.site_id WHERE a.user_id = 681972 AND a.status = 0
// SELECT a.*,b.pid as b_pid,b.source_site as b_source_site,b.title as b_title,c.pid as c_pid,c.json as c_json,d.site_id as d_site_id,d.proxy_name as d_proxy_name FROM item as a left join product as b on a.pid = b.pid left join product_pagedata as c on b.pid = c.pid left join site as d on b.source_site = d.site_id WHERE a.user_id = 681972 AND a.status = 0 ORDER BY a.item_id desc,b.pid asc LIMIT 2 OFFSET 4
func (ctrl *PagerController) Pager(c *gin.Context) {
	response := responses.New(c)

	var results []map[string]interface{}

	fields := []string{
		"a.*",
		"b.pid as b_pid",
		"b.source_site as b_source_site",
		"b.title as b_title",
		"c.pid as c_pid",
		"c.json as c_json",
		"d.site_id as d_site_id",
		"d.proxy_name as d_proxy_name",
	}

	query := database.DB.
		Select(fields).
		Table("item as a").
		Joins("left join product as b on a.pid = b.pid").
		Joins("left join product_data as c on b.pid = c.pid").
		Joins("left join site as d on b.source_site = d.site_id").
		Where("a.user_id = ?", 681972).
		Where("a.status = ?", 0)

	paginate := paginator.Paginate(c, query, &results)

	for i, v := range results {

		js := v["c_json"].(string)
		var jsData map[string]interface{}

		_ = json.Unmarshal([]byte(js), &jsData)

		results[i]["c_decode_json"] = jsData

	}

	response.ToResponse(gin.H{
		"results":    results,
		"pagination": paginate,
	})

}
