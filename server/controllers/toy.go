package controllers

import (
	"net/http"
	"path/filepath"
	"toy_database/models"
	"toy_database/tool"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type ToyController struct {
	beego.Controller
}

func (c *ToyController) URLMapping() {
	c.Mapping("List", c.List)
}

// @router / [get]
func (c *ToyController) Get() {
	// 获取 index.html 文件的路径
	rootPaht := tool.GetRootPath()
	indexPath := filepath.Join(rootPaht, "static", "index.html")
	if data, err := tool.ReadFile(indexPath); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Internal Server Error")
		return
	} else {
		// 设置响应头并返回 index.html 文件的内容
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html")
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
		c.Ctx.ResponseWriter.Write(data)
	}
}

// 获取玩具列表
// @Title 获取玩具列表
// @Description 获取玩具列表
// @Param   page    query   int64     false   "页码"
// @Param   size    query   int64     false   "每页数量"
// @Success 200 {object} models.JsonToyInfoResp
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Failure 500 Internal Server Error
// @router /list [get]
func (c *ToyController) List() {
	var ret models.JsonToyInfoResp
	// 获取页码和每页数量
	// page, _ := c.GetInt64("page", 1)
	// size, _ := c.GetInt64("size", 10)
	// 获取玩具列表
	ret.Data = append(ret.Data, models.JsonToyInfo{
		ID:          uuid.NewString(),
		Name:        "奥迪·卡特罗",
		Price:       10.5,
		Stock:       1,
		ImagePath:   "image/001.jpg",
		Description: "该车是莱斯尼的遗作之一，合金底盘，有避震，透明独立大灯，后挡风玻璃有蚀刻电热除雾丝。最早于1982年在英格兰生产。引入现代轿车制造风格以后，火柴盒模型更趋时尚和现代化。该车便是迎合当时德国奥迪汽车公司研发的Quattro四轮驱动系列的问世而开发的品种。火柴盒不但在1-75小比例系列中及时推出了该车型，还在其著名的SuperKings“超级王牌”系列中推出了1/38左右大比例模型（参见K95）。1983年起在澳门生产，1987年开始在上海环球制造，1992年模具转移到巴西Inbrima S.A.生产，直到1993年最终退出市场。（数据库基本资料图片提供者：Miketorres，特此表示感谢！）",
		MBNum:       "MB025",
		Brand:       "MatchBox",
		ModelName:   "Audi Quattro",
		YearOnBase:  "1982",
		BuyDate:     "2021-01-01",
	})
	ret.Data = append(ret.Data, models.JsonToyInfo{
		ID:          uuid.NewString(),
		Name:        "宝马I3",
		Price:       2.5,
		Stock:       1,
		ImagePath:   "image/002.jpg",
		Description: "该车由宝马汽车公司授权美泰公司制造，最早于2017年在泰国生产。合金车身，透明全景天窗及全封闭车窗，内饰精致，底盘与侧面底边一体。",
		MBNum:       "MB1039",
		Brand:       "MatchBox",
		ModelName:   "BMW I3",
		YearOnBase:  "2015",
		BuyDate:     "2021-01-01",
	})
	c.Data["json"] = ret
	if err := c.ServeJSON(); err != nil {
		logs.Error("ServeJSON error:", err)
	}
}
