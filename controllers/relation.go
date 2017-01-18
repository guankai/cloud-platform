package controllers

import (
	"github.com/astaxie/beego"
	"service-cloud/utils/pagination"
	"service-cloud/models"
	"github.com/astaxie/beego/logs"
	"net/http"
	"github.com/astaxie/beego/orm"
	km "service-cloud/utils/kong/models"
	kong "service-cloud/utils/kong"
	"github.com/satori/go.uuid"
	"service-cloud/utils/environment"
)

type RelationController struct {
	beego.Controller
}

type UserServiceRet struct {
	UserServices interface{} `json:"userServices"`
	PageInfo     *pagination.Paginator `json:"pageInfo"`
}

var SCHEMAURL = environment.GetEnv("KONG_URL", "http://13.76.42.81:8000")
// @Description 用户获取所有服务
// @Param userName formData string true "用户名"
// @Param page formData int true "当前页"
// @Param pageSize formData int true "每页行数"
// @router /query [post]
func (this *RelationController) GetServiceList() {
	userName := this.GetString("userName")
	page, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("pageSize")
	var query models.RelJson
	_page := pagination.NewPaginator(page, pageSize)
	query.UserName = userName
	query.Limit = pageSize
	query.Offset = _page.Offset()
	userServices, count, err := models.GetAllRelations(&query)
	if err != nil {
		logs.Error("获取用户可用的服务失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取用户可用的服务失败")
	}
	_page.SetNums(count)
	_page.GetTotalPages()
	var relationRet UserServiceRet
	relationRet.UserServices = userServices
	relationRet.PageInfo = _page
	this.Data["json"] = relationRet
	this.ServeJSON()
}
// @Description 用户开启服务
// @Param userName formData string true "用户名"
// @Param serviceId formData string true "服务ID"
// @router /startup [post]
func (this *RelationController) EnableService() {
	userName := this.GetString("userName")
	if len(userName) == 0 {
		logs.Error("用户名不可为空")
		this.CustomAbort(http.StatusBadRequest, "用户名不可为空")
	}
	serviceId := this.GetString("serviceId")
	if len(serviceId) == 0 {
		logs.Error("服务ID不可为空")
		this.CustomAbort(http.StatusBadRequest, "服务ID不可为空")
	}
	//查询服务状态
	relation, err := models.GetServiceByUser(serviceId, userName)
	if err != nil {
		if err == orm.ErrNoRows {
			//用户之前未开启过服务,需要开启服务
			//为用户创建consumer
			var consumer km.Consumer
			consumer.Username = userName
			consumerRet, errCon := kong.AddConsumer(&consumer)
			if errCon != nil {
				logs.Error("生成服务消费者失败%v", errCon)
				this.CustomAbort(http.StatusInternalServerError, "生成服务消费者失败")
			}
			//为用户设置apikey
			apiKeyRet, errKey := kong.CreateAPIKey(consumerRet.ID, "")
			if errKey != nil {
				logs.Error("生成APIKEY失败", errKey)
				this.CustomAbort(http.StatusInternalServerError, "生成APIKEY失败")
			}
			//更新数据库
			var relationInsert models.ClRelation
			relationInsert.UserName = userName
			relationInsert.ApiKey = apiKeyRet.Key
			relationInsert.ApiKeyId = apiKeyRet.ID
			relationInsert.ConsumerId = consumerRet.ID
			relationInsert.Service = relation.Service
			relationInsert.Status = "1"
			relationInsert.RelationId = uuid.NewV4().String()
			errRel := models.InsertRelation(&relationInsert)
			if errRel != nil {
				logs.Error("服务开启失败%v", errRel)
				this.CustomAbort(http.StatusInternalServerError, "服务开启失败")
			}
			this.Data["json"] = map[string]string{"msg":"服务开启成功", "apikey":apiKeyRet.Key, "requestPath":SCHEMAURL + relation.Service.RequestPath}
		} else {
			logs.Error("开启服务失败%v", err)
			this.CustomAbort(http.StatusInternalServerError, "开启服务失败")
		}
	} else {
		//用户之前开启过服务,只需要重新申请apikey
		if relation.Status == "1" {
			this.Data["json"] = map[string]string{"msg":"服务已经开启", "apiKey":relation.ApiKey, "requestPath":SCHEMAURL + relation.Service.RequestPath}
		} else {
			//为用户设置apikey
			apikeyRet, errKey := kong.CreateAPIKey(relation.ConsumerId, "")
			if errKey != nil {
				logs.Error("生成APIKEY失败%v", errKey)
				this.CustomAbort(http.StatusInternalServerError, "生成APIKEY失败")
			}
			//更新relation
			errUpd := models.SetStatus(relation.RelationId, "1", apikeyRet.Key, apikeyRet.ID)
			if errUpd != nil {
				logs.Error("开启服务失败%v", errUpd)
				this.CustomAbort(http.StatusInternalServerError, "开启服务失败")
			}
			this.Data["json"] = map[string]string{"msg":"服务已经开启", "apiKey":apikeyRet.Key, "requestPath":SCHEMAURL + relation.Service.RequestPath}
		}
	}
	this.ServeJSON()
}

// @Description 关闭服务
// @Param userName formData string true "用户名"
// @Param serviceId formData string true "服务名"
// @router /shutdown [post]
func (this *RelationController) ShutdownService() {
	userName := this.GetString("userName")
	if len(userName) == 0 {
		logs.Error("用户名不可为空")
		this.CustomAbort(http.StatusBadRequest, "用户名不可为空")
	}
	serviceId := this.GetString("serviceId")
	if len(serviceId) == 0 {
		logs.Error("服务ID不可为空")
		this.CustomAbort(http.StatusBadRequest, "服务ID不可为空")
	}
	//查询服务状态
	relation, err := models.GetServiceByUser(serviceId, userName)
	if err != nil {
		logs.Error("服务查询失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "服务查询失败")
	}
	//删除KONG对应的apikey
	errKong := kong.DeleteAPIKey(relation.ConsumerId, relation.ApiKeyId)
	if errKong != nil {
		logs.Error("删除apikey失败%v", errKong)
		this.CustomAbort(http.StatusInternalServerError, "删除apikey失败")
	}
	//更新对应的status
	errUpd := models.SetStatus(relation.RelationId, "0", relation.ApiKey, relation.ApiKeyId)
	if errUpd != nil {
		logs.Error("服务关闭失败%v", errUpd)
		this.CustomAbort(http.StatusInternalServerError, "服务关闭失败")
	}
	this.Data["json"] = map[string]string{"msg":"服务关闭成功", "requestPath":SCHEMAURL + relation.Service.RequestPath}
	this.ServeJSON()
}