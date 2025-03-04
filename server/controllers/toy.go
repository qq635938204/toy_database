package controllers

import (
	"net/http"
	"path/filepath"
	"toy_database/tool"

	beego "github.com/beego/beego/v2/server/web"
)

type ToyController struct {
	beego.Controller
}

// @router / [get]
func (c *ToyController) Get() {
	// 获取 index.html 文件的路径
	indexPath := filepath.Join("web", "vite-toy-app", "dist", "index.html")
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
