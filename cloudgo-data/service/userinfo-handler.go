package service

import (
	"net/http"
	"strconv"

	"cloudgo-data/entities"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u.DepartName = req.Form["departname"][0]
		entities.UserInfoService.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Form["userid"] == nil || len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			u := entities.UserInfoService.FindByID(int(i))
            if isEmpty(u){
            	formatter.Text(w, http.StatusBadRequest, "There is no record for the user id.")
			}
			formatter.JSON(w, http.StatusBadRequest, u)
			return
		}
		ulist := entities.UserInfoService.FindAll()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}

func isEmpty(u *entities.UserInfo) bool{
	if len(u.UserName) == 0{
		return true
	}
	return false
}
