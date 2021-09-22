package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/polaris1119/chatroom/logic"
)

// 使用通配符可以解决对应的上级目录的问题
//go:embed  ../template/*
var tempfs embed.FS

func HomeHandleFunc(w http.ResponseWriter, req *http.Request) {
	data, err := tempfs.ReadFile("index.html")
	if err != nil {
		return
	}
	fmt.Fprint(w, data)
	//
	// tpl, err := template.ParseFiles(global.RootDir + "/template/home.html")
	// if err != nil {
	// 	fmt.Fprint(w, "模板解析错误！")
	// 	return
	// }
	//
	// err = tpl.Execute(w, nil)
	// if err != nil {
	// 	fmt.Fprint(w, "模板执行错误！")
	// 	return
	// }
}

func UserListHandleFunc(w http.ResponseWriter, req *http.Request) {
	// 加上头部的一些信息
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}
}
