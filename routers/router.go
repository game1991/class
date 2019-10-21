package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//session过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, FilterFunc)
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{})
	//注意：当实现了自定义的get请求方法，请求将不会访问默认方法
	//激活用户
	beego.Router("/active",&controllers.UserController{},"get:ActiveUser")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
    //首页
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex;post:HandleSelect")
	//添加功能
	beego.Router("/Article/addArticle", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
	//查询功能
	beego.Router("/Article/content", &controllers.MainController{}, "get:ShowContent")
	//更新文章功能
	beego.Router("/Article/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	//删除文章功能
	beego.Router("/Article/delete", &controllers.MainController{}, "get:Delete")
	//分类类型
	beego.Router("/Article/addType", &controllers.MainController{}, "get:ShowAddType;post:HandleAddType")
	//删除类型
	beego.Router("/Article/deleteType", &controllers.MainController{}, "get:DeleteType")
	//推出登陆
	beego.Router("user/logout", &controllers.UserController{}, "get:Logout")
	//用户中心页面
	beego.Router("/user/userCenterInfo",&controllers.UserController{},"get:ShowUserCenterInfo")
	//用户中心订单页面
	beego.Router("/user/userCenterOrder",&controllers.UserController{},"get:ShowUserCenterOrder")
	//用户中心收货地址
	beego.Router("/user/UserCenterSite",&controllers.UserController{},"get:ShowUserCenterSite;post:HandleUserCenterSite")
	//商品详情页
	beego.Router("/goodsDetail",&controllers.GoodsController{},"get:ShowGoodsDetail")
	//商品列表页
	beego.Router("/goodsList",&controllers.GoodsController{},"get:ShowList")
	//redis界面
	beego.Router("/redis", &controllers.RedisController{})

}

var FilterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
		return
	}
}
