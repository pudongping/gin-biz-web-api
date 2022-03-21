// 分页器
package paginator

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/logger"
)

// Pagination 分页数据
type Pagination struct {
	Page        int    `json:"page"`          // 当前页码
	PerPage     int    `json:"per_page"`      // 每页显示数
	TotalPage   int    `json:"total_page"`    // 总页数
	TotalCount  int    `json:"total_count"`   // 数据总条数
	PrevPageUrl string `json:"prev_page_url"` // 上一页的链接
	NextPageUrl string `json:"next_page_url"` // 下一页的链接
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL    string   // 用以拼接 URL
	Page       int      // 当前页码
	PerPage    int      // 每页显示数
	Offset     int      // 数据库读取数据时的偏移量
	TotalCount int      // 数据总条数
	TotalPage  int      // 总页数 = math.ceil(TotalCount/PerPage)
	SortColumn []string // 排序字段
	SortRules  []string // 排序规则

	query *gorm.DB // gorm 查询句柄
	ctx   *gin.Context
}

// Paginate 分页器
// c 用于获取当前链接和取参
// db gorm 查询句柄，用于查询数据集和获取数据总数
// data 传地址获取数据
// perPage 每页显示数，变长参数，如果不传时，优先使用链接地址中的 per_page 的值，如果都没有则使用默认值
// 使用示例：
// 	var users []user_model.User
//	query := database.DB.Model(user_model.User{}).Where("id >= ?", 3)
//	paginate := paginator.Paginate(c, query, &users, 3)
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, perPage ...int) Pagination {

	// 初始化 Paginator 类实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	// 初始化分页所需要使用到的属性值
	p.initProperties(perPage...)

	// p.query = p.query.Preload(clause.Associations) // 预加载全部关联
	// 排序
	for k := range p.SortColumn {
		p.query = p.query.Order(fmt.Sprintf("%s %s", p.SortColumn[k], p.SortRules[k]))
	}

	if err := p.query.Limit(p.PerPage).Offset(p.Offset).Find(data).Error; err != nil {
		// 数据库出错
		logger.LogErrorIf(err)
		return Pagination{}
	}

	return Pagination{
		Page:        p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		PrevPageUrl: p.getPrevPageURL(),
		NextPageUrl: p.getNextPageURL(),
	}

}

// initProperties 初始化分页必须要用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage ...int) {

	p.BaseURL = p.formatBaseURL()
	p.PerPage = p.getPerPage(perPage...)
	p.SortColumn, p.SortRules = p.getOrderBy()

	// 以下几个字段，需要注意调用顺序，因为上下几个参数有相关依赖
	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getPage()

	// 计算数据偏移量
	p.Offset = (p.Page - 1) * p.PerPage

}

// getPerPage 获取每页显示数
func (p *Paginator) getPerPage(perPage ...int) int {
	var perPageNum int

	// 如果当前的请求地址中已经含有 `per_page` 参数时
	queryPerPage := p.ctx.Query(config.GetString("paginator.url_query_per_page"))
	if len(queryPerPage) > 0 {
		perPageNum = cast.ToInt(queryPerPage)
	}

	if perPageNum > 0 {
		return perPageNum
	}

	// 如果请求地址中没有 `per_page` 参数时，或者 `per_page` 参数为空时
	// 检查是否有从程序中传参
	if len(perPage) > 0 {
		perPageNum = perPage[0]
	} else {
		// 使用配置文件中的默认每页显示数
		perPageNum = config.GetInt("paginator.default_per_page")
	}

	return perPageNum
}

// getPage 获取当前页码
func (p *Paginator) getPage() int {

	// 如果当前的请求地址中已经含有 `page` 参数时
	page := cast.ToInt(p.ctx.Query(config.GetString("paginator.url_query_page")))
	if page <= 0 {
		// 默认为第一页
		page = 1
	}

	// 如果数据总页数等于 0 时，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}

	// 当请求页数大于总页数时，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}

	return page
}

// getOrderBy 获取排序字段
func (p *Paginator) getOrderBy() (sortColumn []string, sortRules []string) {

	// eg：id,desc|name,asc
	queryOrderBy := p.ctx.Query(config.GetString("paginator.url_query_order_by"))

	if "" == queryOrderBy {
		return
	}

	// eg： [id,desc name,asc]
	orderBySlice := strings.Split(queryOrderBy, "|")

	for _, orderByItem := range orderBySlice {

		rank := strings.Split(orderByItem, ",")
		if len(rank) != 2 || (rank[0] == "" || rank[1] == "") {
			continue
		}

		sortColumn = append(sortColumn, rank[0]) // eg：[id name]
		sortRules = append(sortRules, rank[1])   // eg：[desc asc]
	}

	return
}

// getTotalCount 获取数据总条数
func (p *Paginator) getTotalCount() int {
	var count int64
	newQuery := *p

	// 获取数据总条数的时候不应该出现 order 、limit、offset
	if err := newQuery.query.Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return 0
	}

	return int(count)
}

// getTotalPage 获取总页数
func (p *Paginator) getTotalPage() int {
	// 数据总数为 0 时，则数据总页数（最大页码数就不需要计算）
	if p.TotalCount == 0 {
		return 0
	}

	// 最大页码数 = 数据总数除以每页显示数的商并向上取整数
	// 注意这里不能写成 math.Ceil(float64(p.TotalCount / p.PerPage)) 因为整数除以整数 go 依然会返回整数
	// eg： 9 / 2 = 4
	// 不加 int math.Ceil 计算结果可能为 `-0` ==> math.Ceil(float64(-1) / float64(2)) = -0
	maxPage := int(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if maxPage == 0 {
		maxPage = 1
	}

	return maxPage
}

// formatBaseURL 处理原始请求地址
func (p *Paginator) formatBaseURL() string {
	// `0.0.0.0:3000` + `/api/user`
	baseURL := p.ctx.Request.Host + p.ctx.Request.URL.Path

	// eg：`aa=11&bb=22&page=3`
	if query := p.ctx.Request.URL.RawQuery; "" != query {
		// 从原始链接中去除掉 `page` 参数，因为只有 `page` 参数需要通过程序重新计算返回给客户端
		if s := app.RemoveQueryKey(query, []string{config.GetString("paginator.url_query_page")}); "" != s {
			// 新的链接地址
			baseURL = baseURL + "?" + s
		}
	}

	if strings.Contains(baseURL, "?") {
		// 如果此时的 url 中已经包含 query 参数时
		baseURL = baseURL + "&" + config.GetString("paginator.url_query_page") + "="
	} else {
		// 如果此时的 url 中不包含 query 参数时
		baseURL = baseURL + "?" + config.GetString("paginator.url_query_page") + "="
	}

	return baseURL
}

// getPrevPageURL 得到上一页的链接
func (p *Paginator) getPrevPageURL() string {
	// 如果当前页小于 1 或者就为第一页
	// 或者当前页大于总页数时，不显示上一页的链接（其实这里不可能出现当前页大于总页数的情况，因为 getPage() 中有限制）
	// 假设当前页为 8，最大页数为 6 那么按照正常逻辑上一页应该为 7 但是 6 < 7 则依然无法访问数据
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}

	return p.getPageLink(p.Page - 1)
}

// getNextPageURL 获取下一页的链接
func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}

	return ""
}

// getPageLink 拼接分页链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf(
		"%v%v",
		p.BaseURL,
		page, // 当前页码数
	)
}
