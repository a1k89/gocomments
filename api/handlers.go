package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "gocomments/models"
	u "gocomments/utils"
	"net/http"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cType := vars["content_type"]
	comment := &m.Comment{ContentType: cType}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		return
	}
	resp := comment.Create()
	u.Response(w, resp)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cType := vars["content_type"]
	data := m.GetComments(cType)
	resp := u.Message(data)
	u.Response(w, resp)
}
