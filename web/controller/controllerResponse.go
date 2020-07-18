package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/chainHero/heroes-service/blockchain"
)

type Application struct {
	Fabric *blockchain.FabricSetup
}

//ShowView ...
func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	// 指定视图所在路径
	pagePath := filepath.Join("web", "tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("创建模板实例错误: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("在模板中融合数据时发生错误: %v", err)
		//fmt.Fprintf(w, "显示在客户端浏览器中的错误信息")
		return
	}

}
