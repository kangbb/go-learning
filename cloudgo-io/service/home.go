package service

import (
	"net/http"
    "time"
	"github.com/unrolled/render"
	"text/template"
)

func apitest(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.Text(w, http.StatusOK, time.Now().Format("2006-01-02"))
	}
}
func homeHandle(w http.ResponseWriter, r *http.Request){
	/*使用template.ParseFiles()实现模板的渲染输出
	 *文件路径的根目录以可执行文件为基准*/
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)

	/*使用render的封装的模板实现
	 *文件路径的根目录以可执行文件为基准*/
	/*formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
	})
	formatter.HTML(w, http.StatusOK, "index", "")*/
}