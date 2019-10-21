package controllers

import (
	"class/models"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	/*Ⅰ、增
	//1、有orm对象
	o := orm.NewOrm()
	//2、有一个要插入数据的结构体对象
	user := models.User{}
	//3、对结构体对象赋值
	user.Name = "小明"
	user.Pwd = "123456"
	//4、插入数据
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		return
	}
	*/
	/*Ⅱ、查
	//1、有orm对象
	o := orm.NewOrm()
	//2、查询的对象
	user := models.User{}
	//3、指定查询对象字段值
	//user.Id = 1
	user.Name = "小明"
	//4、查询
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功", user)
	*/
	/*Ⅲ、改
	//1、有orm对象
	o := orm.NewOrm()
	//2、需要更新的结构体对象
	user := models.User{}
	//3、查到需要更新的数据
	user.Id = 1
	err := o.Read(&user)
	//4、给数据重新赋值
	if err == nil {
		user.Name = "小红"
		user.Pwd = "666666"
	}
	//5、更新
	_, err = o.Update(&user)
	if err != nil {
		beego.Info("更新失败", err)
		return
	}
	*/
	/*Ⅳ删除
	//1、有orm对象
	o := orm.NewOrm()
	//2、删除的对象
	user := models.User{}
	//3、指定删除的那一条
	user.Id = 1
	//4、删除
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除错误", err)
		return
	}
	*/

	c.TplName = "register.html"
}

func (c *UserController) Post() {
	//1、拿到数据(去除两边空格)
	Name := strings.TrimSpace(c.GetString("username"))
	Pwd := strings.TrimSpace(c.GetString("pwd"))
	cpwd:=strings.TrimSpace(c.GetString("cpwd"))
	email:=strings.TrimSpace(c.GetString("email"))

	//beego.Info("账号:",Name,"密码:",Pwd)
	//2、对数据进行校验
	if Name == "" || Pwd == ""||cpwd==""||email==""{
		beego.Error("数据不完整，请重新输入")
		c.Redirect("/register", 302)
		return
	}
	if cpwd!=Pwd{
		c.Data["errmsg"]="两次输入秘密不一致，重新输入！"
		c.Redirect("/register",302)
		return
	}
	reg,_:=regexp.Compile("^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	res:=reg.FindString(email)
	if res==""{
		c.Data["errmsg"]="邮箱输入格式有误，请重新输入"
		c.TplName="register.html"
		return
	}
	//3、插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = Name
	user.Pwd = Pwd
	//user.Email=email
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("注册失败，请更换数据注册", err)
		c.Redirect("/register", 302)
		return
	}
	//注册激活，发送邮箱
	emailConfig:=`{"username":"296813329@qq.com","password":"开启服务时提供的串码","host":"smtp.qq.com","port":587}`
    emailConn:=utils.NewEMail(emailConfig)
    emailConn.From="电商网站注册验证服务"
    emailConn.To=[]string{email}
    emailConn.Subject="电商网站用户注册"
    //正文发送给用户的是激活请求地址
    emailConn.Text="127.0.0.1：8081/active?id="+strconv.Itoa(user.Id)//线上搭好的服务器地址
    emailConn.Send()
	//4、返回登录界面
	//测试是否注册成功
	c.Ctx.WriteString("注册成功,请前往邮箱激活！")
	//c.Redirect("/login", 302)

}

func(c*UserController)ActiveUser(){
	//获取数据
	id,err:=c.GetInt("id")
	//校验数据
	if err!=nil{
		beego.Error("要激活的用户不存在")
		c.TplName="register.html"
		return
	}
	//处理数据
	o:=orm.NewOrm()
	var user models.User
	user.Id=id
	err=o.Read(&user)
	if err!=nil{
		c.Data["errmsg"]="要激活的用户不存在!"
		c.TplName="register.html"
		return
	}
	user.Active=true
	o.Update(&user)
	//返回视图
    c.Redirect("/login",302)
}

//获取login界面
func (c *UserController) ShowLogin() {
	username := c.Ctx.GetCookie("username")
	//解码
	temp, _ := base64.StdEncoding.DecodeString(username)
	if string(temp) == "" {
		c.Data["username"] = ""
		c.Data["checked"] = ""
	}else {
		c.Data["username"] = string(temp)
		c.Data["checked"] = "checked"
	}
	c.TplName = "login.html"

}

//登陆业务处理
func (c *UserController) HandleLogin() {
	//c.Ctx.WriteString("这是登陆的post请求")
	//1、拿到数据(去除两边空格)
	name := strings.TrimSpace(c.GetString("username"))
	pwd := strings.TrimSpace(c.GetString("pwd"))
	remember := c.GetString("remember")
	beego.Info("remember is ", remember)
	beego.Info("账号:", name, "密码:", pwd)
	//2、对数据进行校验
	if name == "" || pwd == "" {
		beego.Info("用户名或密码不能为空")
		c.TplName = "login.html"
		c.Data["errmsg"] = "用户名或密码不能为空!"
		return
	}
	//3、查询帐号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名失败")
		c.TplName = "login.html"
		c.Data["errmsg"] = "用户名或密码登陆失败!!"
		return
	}
	if user.Pwd != pwd {
		beego.Info("密码失败")
		c.TplName = "login.html"
		c.Data["errmsg"] = "用户名或密码登陆失败!!"
		return
	}
	if user.Active !=true{
		c.Data["errmsg"]="用户未激活，请前往邮箱激活！"
		c.TplName = "login.html"
		return
	}
	//实现记住用户名
	if remember == "on" {
		//base64加密
		temp := base64.StdEncoding.EncodeToString([]byte(name))
		c.Ctx.SetCookie("username", temp, time.Second*60*60*24*15)

		beego.Info("123")
	} else {
		c.Ctx.SetCookie("username", name, -1)
		beego.Info("4456")
	}

	c.SetSession("userName", name)

	//4、跳转
	//c.Ctx.WriteString("欢迎你，登陆成功")
	//实际注册成功返回登陆后页面
	c.Redirect("/index", 302)

}
//登出用户
func (c *UserController) Logout() {
	//1、删除登陆状态
	c.DelSession("userName")
	//2、跳转登陆界面
	c.Redirect("/", 302)
}
//用户中心页面
func (c *UserController)ShowUserCenterInfo(){
	userName:=GetUser(&c.Controller)
	c.Data["userName"]=userName
	//思考？不登陆的时候能访问到ShowUserCenterInfo函数吗？答：不能

	//查询地址表的内容
	o:=orm.NewOrm()
	//高级查询，表关联
	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Name",userName).Filter("Isdefault",true).One(&addr)
	if addr.Id==0{
		c.Data["addr"]=""
	}else {
		c.Data["addr"] = addr
	}
	c.Layout="userCenterLayout.html"
	c.TplName="User_Center_Info.html"
}
//展示用户中心订单页
func (c *UserController)ShowUserCenterOrder(){
	GetUser(&c.Controller)

	c.Layout="UserCenterLayout.html"
	c.TplName="User_Center_Order.html"
}
//展示用户收货地址
func (c *UserController)ShowUserCenterSite(){
	userName:=GetUser(&c.Controller)
	//c.Data["userName"]=userName
	//获取地址信息
	o:=orm.NewOrm()
	var addr models.Address
	o.QueryTable("Address").RelatedSel("User").Filter("User__Name",userName).Filter("Isdefault",true).One(&addr)
	c.Data["addr"]=addr
	c.Layout="UserCenterLayout.html"
	c.TplName="User_Center_Site,html"

}
//处理用户中心地址数据
func (c *UserController)HandleUserCenterSite(){
	//获取数据
	receiver:=c.GetString("receiver")
	addr:=c.GetString("addr")
	zipCode:=c.GetString("zipCode")
	phone:=c.GetString("phone")
	//校验数据
	if receiver == ""|| addr == ""|| zipCode==""|| phone==""{
		beego.Info("添加数据不完整")
		c.Redirect("/user/UserCenterSite",302)
	}
	//处理数据
	o:=orm.NewOrm()
	var addrUser models.Address
	addrUser.Isdefault=true
	err:=o.Read(&addrUser,"Isdefault")
	/*
	if err!=nil{
		addrUser.Receiver=receiver
		addrUser.Zipcode=zipCode
		addrUser.Addr=addr
		addrUser.Phone=phone
		o.Insert(&addrUser)
	}else{
		addrUser.Isdefault=false
		o.Update(&addrUser)
	}
	*/
	//添加默认地址前需要把原来的地址更新成非默认地址
	if err == nil{
		addrUser.Isdefault=false
		o.Update(&addrUser)
	}
	//更新默认地址时，给原来的地址对象Id赋值，这时候用原来的地址对象插入，意思是用原来的Id做插入操作，会报错
	//关联
	userName:=c.GetSession("userName")
	var user models.User
	user.Name=userName.(string)
	o.Read(&user,"Name")
	var addrUserNew models.Address
	addrUserNew.Receiver=receiver
	addrUserNew.Zipcode=zipCode
	addrUserNew.Addr=addr
	addrUserNew.Phone=phone
	addrUserNew.Isdefault=true
	addrUserNew.User=&user
	o.Insert(&addrUserNew)

	//返回视图
	c.Redirect("/user/UserCenterSite",302)

}

