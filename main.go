package main

import (
	"github.com/astaxie/beego"
	_ "github.com/sy/bomer/routers"
)

func main() {
	beego.AddFuncMap("add", add)
	beego.Run()
}

func add(n, m int) int {
	return n + m
}
