package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//表的设计

type User struct {
	Id       int
	Name     string `orm:"unique"`
	Pwd      string
	Active   bool `orm:"default(false)"` //用户注册激活判断
	Power    int `orm:"default(0)"`//管理员权限 0表示未激活 1表示激活
	Address  []*Address `orm:"reverse(many)"`
	OrderInfo []*OrderInfo `orm:"reverse(many)"`
}
//地址表
type Address struct {
	Id int
	Receiver string `orm:"size(20)"`//收件人
	Addr     string `orm:"size(50)"`//收件地址
	Zipcode  string `orm:"size(20)"`//邮编
	Phone    string `orm:"size(20)"`//联系电话
	Isdefault bool  `orm:"default(false)"`//是否默认地址 false为非默认 true为默认
	User *User `orm:"rel(fk)"`//用户ID
	OrderInfo  []*OrderInfo `orm:"reverse(many)"`

}
//商品SPU表
type Goods struct {
	Id int
	Name string `orm:"size(20)"`//商品名称
	Detail string `orm:"size(200)"`//商品详细描述
	GoodsSKU []*GoodsSKU `orm:"reverse(many)"`
}
//商品类型
type GoodsType struct {
	Id int
	Name string//种类名字
	Logo string//图标
	Image string//图片
	GoodsSKU []*GoodsSKU `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}
//商品SKU表
type GoodsSKU struct {
	Id int
	Goods *Goods `orm:"rel(fk)"`
	GoodsType *GoodsType `"orm:"rel(fk)"`
	Name string
	Desc string//商品描述
	Price float32 //商品价格
	Unite string//商品单位
	Image string//商品图片
	Stock int `orm:"default(1)"`//商品库存
	Sales int `orm:"default(0)"`//商品销量
	Status int `orm:"default(1)"`//商品状态
	Time time.Time `orm:"auto_now_add"`
	GoodsImage []*GoodsImage `orm:"reverse(many)"`
	IndexGoodsBanner []*IndexGoodsBanner `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
	OrderGoods []*OrderGoods `orm:"reverse(many)"`
}
//商品图片表
type GoodsImage struct {
	Id int
	Image string
	GoodsSKU *GoodsSKU `orm:"rel(fk)"`
}
//首页轮播展示图
type IndexGoodsBanner struct {
	Id int
	Image string
	Index int `orm:"default(0)"`
	GoodsSKU *GoodsSKU `orm:"rel(fk)"`
}
//首页分类商品展示图
type IndexTypeGoodsBanner struct {
	Id int
	GoodsType *GoodsType `orm:"rel(fk)"`
	GoodsSKU *GoodsSKU `orm:"rel(fk)"`
	Display_Type int `orm:"default(1)"`//展示类型方式 0代表图片 1代表图片
	Index int `orm:"default(0)"`//展示顺序
}
//首页促销商品展示表
type IndexPromotionBanner struct {
	Id int
	Name string `orm:"size(20)"`//活动名称
	Url string `orm:"size(50)"`//活动链接
	Image string
	Index int `orm:"default(0)"`
}
//订单表
type OrderInfo struct {
	Id int
	OrderId string `orm:"unique"`
	User *User `orm:"rel(fk)"`
	Address *Address `orm:"rel(fk)"`
	Pay_Method int
	Total_Count int `orm:"default(1)"`//商品数量
	Total_Price float32//商品总价
	Transit_Price float32//运费
	Order_status int `orm:"default(1)"`//订单状态
	Trade_No string `orm:"default('')"`//支付编号
	Time time.Time `orm:"auto_now_add"`//评论时间
	OrderGoods []*OrderGoods `orm:"reverse(many)"`
}
//订单商品表
type OrderGoods struct {
	Id int
	OrderInfo *OrderInfo `orm:"rel(fk)"`//订单信息
	GoodsSKU *GoodsSKU `orm:"rel(fk)"`//商品
	Count int `orm:"default(1)"`//商品数量
	Price int //商品价格
	Comment string `orm:"default('')"`//评论
}

func init() {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test2?charset=utf8&loc=Asia%2FShanghai")
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	//映射model数据
	orm.RegisterModel(new(User), new(Address), new(Goods),new(GoodsType),new(GoodsSKU),new(GoodsImage),new(IndexGoodsBanner),new(IndexTypeGoodsBanner),new(IndexPromotionBanner),new(OrderInfo),new(OrderGoods))
	//生成表
	orm.RunSyncdb("default", false, true)

}
