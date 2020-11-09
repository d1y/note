// create by d1y<chenhonzhou@gmail.com>
// api接口

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	v1 "github.com/d1y/note/api/v1"
	"github.com/d1y/note/conf"
	"github.com/d1y/note/utils/randx"
	"github.com/julienschmidt/httprouter"
)

func easySendJSON(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

// Index302 首页
func Index302(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var r = randx.CreateEasyRandomString()
	r = fmt.Sprintf("/%s/%s", conf.WebPrefix, r)
	http.Redirect(w, req, r, http.StatusFound)
}

// GetRouterData 获取数据
func GetRouterData(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var router = p.ByName("router")
	var data = v1.GetData(router)
	easySendJSON(w, data)
}

// PostRouterData 更新数据
func PostRouterData(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	req.ParseForm()
	var r, c, t = req.FormValue("router"), req.FormValue("content"), req.FormValue("title")
	var o = v1.OutputData{
		Code:    v1.ErrorCode,
		Message: v1.ErrorValidation.Error(),
	}
	if len(r) >= 1 && (len(c) >= 1 || len(t) >= 1) {
		var n = v1.Note{
			Title:   t,
			Content: c,
			Router:  r,
		}
		var flag = v1.UpdateData(n)
		if flag == nil {
			o.Code = v1.SuccessCode
			o.Message = v1.UpdateDataSuccess
		}
		o.Data = n
	}
	easySendJSON(w, o)
}
