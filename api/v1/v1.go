// Author: d1y<chenhonzhou@gmail.com>
// link: https://www.alexedwards.net/blog/golang-response-snippets
// link: https://stackoverflow.com/a/39711522

package v1

import (
	"time"

	"github.com/d1y/note/conf"
	"github.com/d1y/note/db"
)

// Note note
type Note struct {
	ID      string    `db:"id,omitempty"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Router  string    `db:"router"`
	Create  time.Time `db:"create_at"`
	Update  time.Time `db:"update_at"`
}

// OutputData 返回数据
type OutputData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Note   `json:"data"`
}

// GetData 获取数据
func GetData(router string) (o OutputData) {
	var cacheData = db.Session.SQL().Select().From(conf.DatabaseName).Where("router = ?", router)
	var r Note
	err := cacheData.One(&r)
	o = OutputData{
		Code:    SuccessCode,
		Message: GetDataSuccess,
		Data:    r,
	}
	if err != nil {
		o.Code = ErrorCode
		o.Message = err.Error()
	}
	return o
}

// UpdateData 更新数据
func UpdateData(note Note) (e error) {
	var R = note.Router
	var r = GetData(R)
	if r.Code == SuccessCode {

		var sqlMatch = ""
		var arg []string
		if len(note.Title) >= 1 {
			sqlMatch += "title = ?"
			arg = append(arg, note.Title)
		}
		if len(note.Content) >= 1 {
			sqlMatch += "content = ?"
			arg = append(arg, note.Content)
		}
		if len(note.Title) >= 1 && len(note.Content) >= 1 {
			sqlMatch = "title = ?, content = ?"
			arg = []string{note.Title, note.Content}
		}

		// []string => []interface{}
		//
		// https://stackoverflow.com/a/59674266
		var iargs = make([]interface{}, 0)
		for i, x := range arg {
			if i == 0 {
				iargs = append(iargs, sqlMatch, x)
				continue
			}
			iargs = append(iargs, x)
		}

		var cache = db.Session.SQL().Update(conf.DatabaseName).Set(iargs...).Where("router = ?", R)
		_, err := cache.Exec()
		e = err
	} else {
		_, err := db.NoteDB.Insert(note)
		if err != nil {
			e = err
		}
	}
	return e
}
