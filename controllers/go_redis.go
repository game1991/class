package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

type RedisController struct {
	beego.Controller
}

//1、连接数据库

//2、操作数据库

//3、关闭数据库

func (this *RedisController) Get() {
	conn, _ := redis.Dial("tcp", ":6379")
	defer conn.Close()
	//conn.Send("set","name","game")
	//conn.Send("mset","age",11,"score",120)
	//conn.Flush()
	//reply,err:=conn.Receive()
	//if err!=nil{
	//	this.Ctx.WriteString("设置内容错误")
	//return
	//}
	//beego.Info(reply)
	//reply,err =conn.Do("set","sex","female")
	//beego.Info(reply)
	//this.Ctx.WriteString("执行成功！")

	//conn.Send("MULTI")
	//conn.Send("get","name")
	//conn.Send("set","class","1 class")
	//reply,err:=conn.Do("EXEC")
	//if err!=nil{
	//	this.Ctx.WriteString("事务操作错误")
	//	beego.Info(err)
	//	return
	//}
	//beego.Info(reply)
	reply, err := redis.Values(conn.Do("mget", "name", "age"))
	if err != nil {
		this.Ctx.WriteString("事务操作错误")
		beego.Info(err)
		return
	}
	var i int
	var s string
	redis.Scan(reply, &s, &i)
	beego.Info(s, i)

	this.Ctx.WriteString("执行成功！")

}
