package main

import (
	_ "class/models"
	_ "class/routers"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("PrePage", HandlePrePage)
	beego.AddFuncMap("NextPage", HandleNextPage)
	beego.Run()
}

func HandlePrePage(data int) string {
	pageIndex := strconv.Itoa(data - 1)
	return pageIndex

}
func HandleNextPage(data int) string {
	pageIndex := strconv.Itoa(data + 1)
	return pageIndex

}
