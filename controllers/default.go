package controllers

import (
	"class/models"
	"math"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}
//举例内容Get、Post

//显示首页内容
func (c *MainController) ShowIndex() {
	userName:=c.GetSession("userName")
	beego.Info(userName)
	if userName==nil{
		c.Data["userName"]=""
		//c.Redirect("/login",302)//如果没有session，回到登录页面
	}else {
		c.Data["userName"]=userName.(string)
	}

	o := orm.NewOrm()
	//fmt.Printf("o的类型:%T\n",o)
	//var articles []models.Aticle
	qs := o.QueryTable("Aticle")
	//fmt.Printf("type:%T\n", qs)
	//fmt.Printf("type:%T\n", o)
	//qs.All(&articles) 相当于select * from Aticle
	//var articlewithtype []models.Aticle
	//根据类型获取数据
	//1、接收数据
	typeName := c.GetString("select")
	//2、处理数据
	if typeName == "" {
		beego.Info("下拉框传递数据失败或者默认显示主页")
		//count, err := qs.RelatedSel("AType").Count() //返回数据条目数  加个过滤器
		//beego.Info(count)
		//if err != nil {
		//	beego.Info("查询页面为空或者错误!")
		//	return
		//}
		//pageIndex, err := strconv.Atoi(c.GetString("pageIndex")) //每页的页码，起始值为1
		//if err != nil {
		//	pageIndex = 1
		//}
		////获取当前页码
		//pageSize := 2 //设定每页显示多少个记录
		//start := pageSize * (pageIndex - 1)
		//qs.Limit(pageSize, start).RelatedSel("AType").All(&articles) //第一个参数表示一页显示多少个内容；第二个参数表示从哪里开始start
		//
		//pageCount := math.Ceil(float64(count) / float64(pageSize))   //总页数=总记录数/每页记录数；ceil是向上取整//向下取整是floor
		////避免上一页超过范围，首页
		//FirstPage := false //标识是否是末页
		//LastPage := false
		//if pageIndex == 1 {
		//	FirstPage = true
		//}
		//if pageIndex == int(pageCount) {
		//	LastPage = true
		//}
		////获取类型数据
		//var types []models.AticleType
		//o.QueryTable("AticleType").All(&types)
		//c.Data["types"] = types
		//c.Data["FirstPage"] = FirstPage
		//c.Data["LastPage"] = LastPage
		//c.Data["count"] = count
		//c.Data["pageCount"] = pageCount
		//c.Data["pageIndex"] = pageIndex
		//qs.Limit(pageSize, start).RelatedSel("AType").All(&articlewithtype)
		filter(c, qs, typeName, false)
	} else {
		//count, err := qs.RelatedSel("AType").Filter("AType__TypeName", typeName).Count()
		//if err != nil {
		//	beego.Info("查询页面为空或者错误!")
		//	return
		//}
		//pageIndex, err := strconv.Atoi(c.GetString("pageIndex")) //每页的页码，起始值为1
		//if err != nil {
		//	pageIndex = 1
		//}
		////获取当前页码
		//pageSize := 2 //设定每页显示多少个记录
		//start := pageSize * (pageIndex - 1)
		//qs.Limit(pageSize, start).RelatedSel("AType").All(&articles) //第一个参数表示一页显示多少个内容；第二个参数表示从哪里开始start
		//pageCount := math.Ceil(float64(count) / float64(pageSize))   //总页数=总记录数/每页记录数；ceil是向上取整//向下取整是floor
		////避免上一页超过范围，首页
		//FirstPage := false //标识是否是末页
		//LastPage := false
		//if pageIndex == 1 {
		//	FirstPage = true
		//}
		//if pageIndex == int(pageCount) {
		//	LastPage = true
		//}
		////获取类型数据
		//var types []models.AticleType
		//o.QueryTable("AticleType").All(&types)
		//c.Data["types"] = types
		//
		//c.Data["FirstPage"] = FirstPage
		//c.Data["LastPage"] = LastPage
		//c.Data["count"] = count
		//c.Data["pageCount"] = pageCount
		//c.Data["pageIndex"] = pageIndex
		//
		//qs.Limit(pageSize, start).RelatedSel("AType").Filter("AType__TypeName", typeName).All(&articlewithtype)
		filter(c, qs, typeName, true)
	}
	//获取类型数据
	var types []models.AticleType
	o.QueryTable("AticleType").All(&types)
	////把类型数据写到redis
	////1、连接数据库
	//conn,_:=redis.Dial("tcp",":6379")
	////3、关闭数据库
	//defer conn.Close()
	////2、执行存储数据
	////conn.Do("set","types",types)此方法行不通，因为无法取出自定义的对象
	////1、编码操作
	//var buffer bytes.Buffer//容器
	//enc:=gob.NewEncoder(&buffer)//获取编码器
	//enc.Encode(types)//编码
	//beego.Info(types)
	//conn.Do("set","types",buffer.Bytes())//redis存储
	//
	//
	//conn,_:=redis.Dial("tcp",":6379")
	//defer conn.Close()
	////查询
	////rep,err:=conn.Do("get","types")
	////reulst,_:=redis.Bytes(rep,err)
	////dec:=gob.NewDecoder(bytes.NewReader(reulst))
	//buffer,err :=redis.Bytes(conn.Do("get","types"))
	//if err!=nil{
	//	beego.Info("获取失败")
	//	beego.Info(err)
	//}
    ////解码
    //dec:=gob.NewDecoder(bytes.NewReader(buffer))
    //dec.Decode(&types)
    //beego.Info(types)


	c.Data["types"] = types
	c.Data["typeName"] = typeName
	//把数据传递给视图
	c.Layout="Layout.html"
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["HtmlHead"]="indexHead.html"
	c.LayoutSections["Scripts"]="index_scripts.html"
	c.TplName = "index.html"

}
//过滤器
func filter(c *MainController, qs orm.QuerySeter, typeName string, isNeedFilter bool ) {
	var articles []models.Aticle
	var articlewithtype []models.Aticle
	var count int64
	if !isNeedFilter==false {
		co, _ := qs.RelatedSel("AType").Filter("AType__TypeName", typeName).Count()
		count=co
	} else {
		co, _ := qs.RelatedSel("AType").Count()
		count=co
	}
	beego.Info(count)
	//刚才复制的重复代码
	pageIndex, err := strconv.Atoi(c.GetString("pageIndex")) //每页的页码，起始值为1
	if err != nil {
		pageIndex = 1
	}
	//获取当前页码
	pageSize := 2 //设定每页显示多少个记录
	start := pageSize * (pageIndex - 1)
	qs.Limit(pageSize, start).RelatedSel("AType").All(&articles) //第一个参数表示一页显示多少个内容；第二个参数表示从哪里开始start
	pageCount := math.Ceil(float64(count) / float64(pageSize))   //总页数=总记录数/每页记录数；ceil是向上取整//向下取整是floor

	//避免上一页超过范围，首页
	FirstPage := false //标识是否是末页
	LastPage := false
	if pageIndex == 1 {
		FirstPage = true
	}
	if pageIndex == int(pageCount) {
		LastPage = true
	}

	c.Data["FirstPage"] = FirstPage
	c.Data["LastPage"] = LastPage
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageIndex"] = pageIndex

	if isNeedFilter ==true{
		qs.Limit(pageSize, start).RelatedSel("AType").Filter("AType__TypeName", typeName).All(&articlewithtype)
	} else {
		qs.Limit(pageSize, start).RelatedSel("AType").All(&articlewithtype)
	}
	c.Data["articles"] = articlewithtype
}

//处理下拉框改变首页显示信息
func (c *MainController) HandleSelect() {
	//1、接收数据
	typeName := c.GetString("select")
	//beego.Info(typeName)
	//2、处理数据
	if typeName == "" {
		beego.Info("下拉框传递数据失败")
		return
	}
	//3、查询数据
	o := orm.NewOrm()
	var artis []models.Aticle
	//fifter()过滤器，相当于sql语句中的where指定查询条件（select * from user where），第一个参数是查询字段，第二个参数是要匹配的值
	//RelatedSel是关系数据表关联，由于filter是惰性查询，需要详细到哪张表
	o.QueryTable("Aticle").RelatedSel("AType").Filter("AType__TypeName", typeName).All(&artis)
	beego.Info(artis)
}

//显示添加文章界面
func (c *MainController) ShowAdd() {
	//查询类型数据，传递给视图中
	o := orm.NewOrm()
	var types []models.AticleType
	o.QueryTable("AticleType").All(&types)

	c.Data["types"] = types
	c.Layout="Layout.html"
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["HtmlHead"]="addHead.html"
	c.TplName = "add.html"
}

//处理添加文章界面
func (c *MainController) HandleAdd() {
	//1、拿到数据
	var fileName string
	artName := c.GetString("articleName")
	artContent := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		beego.Info("没有上传文件或者文件上传失败!")

	} else {
		defer f.Close()
		//1)、限制格式
		fileExt := path.Ext(h.Filename)
		beego.Info(fileExt)
		if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".JPG" && fileExt != ".PNG" {
			beego.Info("上传文件格式不符合")
			return
		}
		//2)、限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}
		//3)、需要对获取文件进行重命名
		fileName = time.Now().Format("2006-01-02 15：04：05") + fileExt

		beego.Info("正在保存")
		err := c.SaveToFile("uploadname", "./static/img/"+fileName)
		if err != nil {
			beego.Info(err)
			return
		}

	}
	//beego.Info(fileName)
	//2、判断数据是否合法
	if artName == "" || artContent == "" {
		beego.Info("添加文章为空不允许")
		return
	}
	//3、插入数据
	o := orm.NewOrm()
	arti := models.Aticle{}
	arti.Aname = artName
	arti.Acontent = artContent
	arti.Aimg = "/static/img/" + fileName

	//给aticle对象赋值
	//获取下拉框传递过来的类型数据
	typeName := c.GetString("select")
	//类型判断
	if typeName == "" {
		beego.Info("下拉框数据为空!")
		return
	}
	//获取type类型对象
	var atiType models.AticleType
	atiType.TypeName = typeName
	err = o.Read(&atiType, "TypeName")
	if err != nil {
		beego.Info("获取类型错误!")
		return
	}
	arti.AType = &atiType

	//3、插入数据
	_, err = o.Insert(&arti)
	if err != nil {
		beego.Info("插入数据库出错!")
		return
	}

	//4、返回文章界面
	c.Redirect("/Article/index", 302)
}

//显示文章内容详情
func (c *MainController) ShowContent() {
	//1、获取文章Id
	id, err := c.GetInt("id")
	beego.Info("id is ", id)
	if err != nil {
		beego.Info("获取文章id出错!!")
		return
	}
	//2、查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Aticle{Id2: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询数据错误!!")
		return
	}
	arti.Acount+=1
	//多对多插入读者
	//1、获取操作对象

	//2、获取多对多操作对象
	m2m:=o.QueryM2M(&arti,"AUsers")
	//3、获取插入对象
	userName:=c.GetSession("userName")
	user:=models.User{}
	user.Name=userName.(string)
	o.Read(&user,"Name")
	//4、多对多插入
	_,err=m2m.Add(&user)
	if err!=nil {
		beego.Info("插入失败")
		return

	}
	o.Update(&arti)//没有指定哪一列时，他会自动搜索更新


	o.LoadRelated(&arti,"AUsers")

	//o.QueryTable("aticle").RelatedSel("User").Filter("AUsers__User__Name",userName.(string)).Distinct().Filter("Id2",id).One(&arti)




	//3、传递给views视图
	c.Data["aticle"] = arti
    c.Layout="Layout.html"
    c.LayoutSections=make(map[string]string)
    c.LayoutSections["HtmlHead"]="contentHead.html"
	c.TplName = "content.html"

}

//显示编辑页面
func (c *MainController) ShowUpdate() {
	//1、获取文章Id
	id, err := c.GetInt("id")
	beego.Info("id is ", id)
	if err != nil {
		beego.Info("获取文章id出错!!")
		return
	}
	//2、查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Aticle{Id2: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询数据错误!!")
		return
	}
	//3、传递给views视图
	c.Data["aticle"] = arti
	c.TplName = "update.html"
}

func (c *MainController) HandleUpdate() {
	//1、拿到数据
	var fileName string
	id, _ := c.GetInt("id")
	beego.Info("id is ", id)

	artName := c.GetString("articleName")
	artContent := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		beego.Info("没有上传文件或者文件上传失败!")
	} else {
		defer f.Close()
		//1)、限制格式
		fileExt := path.Ext(h.Filename)
		//beego.Info(fileExt)
		if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".JPG" && fileExt != ".PNG" {
			beego.Info("上传文件格式不符合")
			return
		}
		//2)、限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}
		//3)、需要对获取文件进行重命名
		fileName = time.Now().Format("2006-01-02 15：04：05") + fileExt
		beego.Info(artName, artContent, fileName)
		beego.Info("正在保存")
		err = c.SaveToFile("uploadname", "./static/img/"+fileName)
		if err != nil {
			beego.Info(err)
			return
		}
	}
	//2、判断数据是否合法
	if artName == "" || artContent == "" {
		beego.Info("添加文章为空不允许")
		return
	}
	//3、插入数据
	o := orm.NewOrm()
	arti := models.Aticle{Id2: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info(err)
		beego.Info("读取数据失败!!")
		return
	}
	arti.Aname = artName
	arti.Acontent = artContent
	arti.Aimg = "/static/img/" + fileName

	_, err = o.Update(&arti, "Aname", "Acontent", "Aimg")
	if err != nil {
		beego.Info("更新数据出错!")
		return
	}

	//4、返回文章界面
	c.Redirect("/Article/index", 302)
}

func (c *MainController) Delete() {
	//1、拿到数据
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	arti := models.Aticle{Id2: id}
	err := o.Read(&arti)
	if err != nil {
		beego.Info("读取数据出错!!")
		return
	}

	//2、删除
	o.Delete(&arti)
	//3、返回页面
	c.Redirect("/Article/index", 302)
}

func (c *MainController) ShowAddType() {
	o := orm.NewOrm()
	var artiTypes []models.AticleType
	_, err := o.QueryTable("AticleType").All(&artiTypes)
	if err != nil {
		beego.Info("查询类型错误!")
	}
	c.Data["types"] = artiTypes
	c.Layout="Layout.html"
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["HtmlHead"]="addTypeHead.html"
	c.LayoutSections["Scripts"]="addType_scripts.html"

	c.TplName = "addType.html"
}

//处理添加类型业务
func (c *MainController) HandleAddType() {
	//1、获取数据
	typename := c.GetString("typeName")
	//2、判断数据
	if typename == "" {
		beego.Info("添加类型不能为空!!")
		return
	}
	//3、执行插入操作
	o := orm.NewOrm()
	var artiTypes models.AticleType
	artiTypes.TypeName = typename
	_, err := o.Insert(&artiTypes)
	if err != nil {
		beego.Info("插入数据错误!!")
		return
	}
	//4、展示视图
	c.Redirect("/Article/addType", 302)
}

func (c *MainController) DeleteType() {
	//1、拿到数据
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	arti := models.AticleType{Id: id}
	err := o.Read(&arti)
	if err != nil {
		beego.Info("读取数据出错!!")
		return
	}
	//2、删除
	o.Delete(&arti)
	//3、返回页面
	c.Redirect("/Article/addType", 302)

}
//退出登陆

