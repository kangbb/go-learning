package service
import (
	"net/http"
	"text/template"
	"log"
)
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["username"][0]
	pass := r.Form["password"][0]
	//请求的是登录数据，那么执行登录的逻辑判断
	if len(name) == 0 || len(pass) == 0{
		log.Fatal("err")
		return
	}
	t := template.Must(template.ParseFiles("templates/chart.html"))
	t.Execute(w, map[string]string{
		"Name":name,
        "Pass":pass,
	})
}

