package controllers

import (
	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GoodsController struct {
	beego.Controller
}

func GetUser(this *beego.Controller) string {
	userName := this.GetSession("userName")
	if userName == nil {
		this.Data["userName"] = ""
	} else {
		this.Data["userName"] = userName.(string)
		return userName.(string)
	}
	return ""
}

func (this *GoodsController) ShowIndex() {
	GetUser(&this.Controller)

	this.TplName = "index.html"
}

//分层页面设计，layout
func ShowLayout(c *beego.Controller){
	//查询类型
	o:=orm.NewOrm()
	var types []models.GoodsType
	o.QueryTable("GoodsType").All(&types)
	c.Data["types"]=types
	//获取用户信息
	GetUser(c)
	//指定Layout
	c.Layout="goodsLayout.html"

}

func (this *GoodsController) ShowDetail() {
	//获取数据
	id, err := this.GetInt("tpyeId")
	//校验数据
	if err != nil {
		beego.Error("浏览器请求路径出错！")
		this.Redirect("/index", 302)
		return
	}
	//处理数据
	o := orm.NewOrm()
	var goodsSKU models.GoodsSKU
	goodsSKU.Id = id
	//o.Read(&goodsSKU)
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("Id", id).One(&goodsSKU)
	//获取同类型的两条商品数据
	var goodsNew []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType", goodsSKU.GoodsType).OrderBy("Time").Limit(2, 0).All(&goodsNew)
	this.Data["goodsNew"] = goodsNew

	//返回视图
	this.Data["goodsSKU"] = goodsSKU

	ShowLayout(&this.Controller)
	this.TplName = "detail.html"
}
//展示商品列表页
func (this *GoodsController)ShowList(){
	//获取数据
	id,err:=this.GetInt("typeId")
	//校验数据
	if err!=nil{
		beego.Error("浏览器请求路径出错！")
		this.Redirect("/", 302)
		return
	}
	//处理数据
	ShowLayout(&this.Controller)
	//获取新品
	o:=orm.NewOrm()
	var goodsNew []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).OrderBy("Time").Limit(2,0).All(&goodsNew)
	this.Data["goodsNew"]=goodsNew
	//获取商品
	var goods []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).All(&goods)
	this.Data["goods"]=goods
	//分页实现
	var pages []int
	if pageCount <= 5 {
		pages={1,2,..,pageCount}
	}

	//返回视图
	this.TplName="list.html"

}