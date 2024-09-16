package main

import (
	"net/http"

	"github.com/hsnice16/golang_learning/deployment/logs"
)


func main() {
	addr := "127.0.0.1:80"
	logs.Logger.Info("Start server at:%v", addr)
	err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	logs.Logger.Critical("Server err:%v", err)
}