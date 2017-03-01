package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	km "cloud-platform/utils/kong/models"
	kong "cloud-platform/utils/kong"
	"cloud-platform/models"
	"github.com/satori/go.uuid"
	"cloud-platform/utils/pagination"
)

type ServiceController struct {
	beego.Controller
}

type ServiceRet struct {
	ServiceList interface{} `json:"serviceList"`
	PageInfo    *pagination.Paginator `json:"pageInfo"`
}
// @Description 注册服务
// @Param serviceName formData string true "服务名称"
// @Param servicePic formData string true "服务logo"
// @Param requestPath formData string true "访问路径"
// @Param provider formData string true "服务提供者"
// @Param version formData string true "版本"
// @Param serviceDesc formData string true "服务概述"
// @Param upstreamUrl formData string true "服务跳转地址"
// @Param typeId formData int true "类型id"
// @router /add [post]
func (this *ServiceController) AddService() {
	serviceName := this.GetString("serviceName")
	if len(serviceName) == 0 {
		logs.Error("服务名称不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务名称不能为空")
	}
	servicePic := this.GetString("servicePic")
	if len(servicePic) == 0 {
		logs.Error("服务logo图片不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务logo图片不能为空")
	}
	requestPath := this.GetString("requestPath")
	if len(requestPath) == 0 {
		logs.Error("服务标识不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务标识不能为空")
	}
	provider := this.GetString("provider")
	if len(provider) == 0 {
		logs.Error("服务提供者不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务提供者不能为空")
	}
	version := this.GetString("version")
	if len(version) == 0 {
		logs.Error("版本不能为空")
		this.CustomAbort(http.StatusBadRequest, "版本不能为空")
	}
	serviceDesc := this.GetString("serviceDesc")
	if len(serviceDesc) == 0 {
		logs.Error("服务概述不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务概述不能为空")
	}
	upstreamUrl := this.GetString("upstreamUrl")
	if len(upstreamUrl) == 0 {
		logs.Error("服务访问路径不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务访问路径不能为空")
	}
	typeId, _ := this.GetInt("typeId")
	var service models.ClService
	service.ServiceName = serviceName
	service.ServicePic = servicePic
	service.RequestPath = `/` + requestPath
	service.Provider = provider
	service.Version = version
	service.ServiceDesc = serviceDesc
	service.UpstreamUrl = upstreamUrl
	service.Id = uuid.NewV4().String()
	appType, errApp := models.GetAppType(typeId)
	if errApp != nil {
		logs.Error("应用类型获取失败%v", errApp)
		this.CustomAbort(http.StatusInternalServerError, "应用类型获取失败")
	}
	service.Type = appType
	//注册KONG的服务
	//生成api
	var api km.API
	api.RequestPath = "/" + requestPath
	api.UpstreamURL = upstreamUrl
	api.Name = requestPath
	api.StripRequestPath = true
	api.RequestHost = requestPath
	apiRet, errApi := kong.AddAPI(&api)
	if errApi != nil {
		logs.Error("服务注册失败 %v", errApi)
		this.CustomAbort(http.StatusInternalServerError, "服务注册失败")
	}
	service.ApiId = apiRet.ID
	logs.Debug("the api_id is %s", apiRet.ID)
	//关联对应的plugin
	plugin, errPlugin := models.GetPluginByName("key-auth")
	if errPlugin != nil {
		logs.Error("获取插件失败 %v", errPlugin)
		this.CustomAbort(http.StatusInternalServerError, "获取插件失败")
	}
	service.Plugin = plugin
	// todo 调用插件API关联apis和plugin
	_, errKong := kong.AddKeyAuthPlugin2Api(apiRet.ID)
	if errKong != nil {
		logs.Error("关联API和plugin失败%v", errKong)
		this.CustomAbort(http.StatusInternalServerError, "关联API和plugin失败")
	}
	errService := models.InsertService(&service)
	if errService != nil {
		logs.Error("服务注册失败 %v", errService)
		this.CustomAbort(http.StatusInternalServerError, "服务注册失败")
	}
	this.Data["json"] = "服务注册成功"
	this.ServeJSON()
}
// @Description "根据条件查询所有条件列表"
// @Param page formData int true "当前页"
// @Param pageSize formData int true "页面显示个数"
// @Param keyword formData string false "serviceName关键字,模糊查询"
// @router /query [post]
func (this *ServiceController) QueryServiceList() {
	page, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("pageSize")
	keyword := this.GetString("keyword")
	var query models.QueryCon
	_page := pagination.NewPaginator(page, pageSize)
	query.Limit = pageSize
	query.Offset = _page.Offset()
	query.Keyword = keyword
	serviceList, count, err := models.GetServiceList(&query)
	if err != nil {
		logs.Error("获取服务列表失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取服务列表失败")
	}
	_page.SetNums(count)
	_page.GetTotalPages()
	var serviceRet ServiceRet
	serviceRet.ServiceList = serviceList
	serviceRet.PageInfo = _page
	this.Data["json"] = serviceRet
	this.ServeJSON()
}

// @Description 修改服务
// @Param serviceId formData string true "服务uuid"
// @Param serviceName formData string true "服务名称"
// @Param servicePic formData string true "服务logo"
// @Param requestPath formData string true "访问路径"
// @Param provider formData string true "服务提供者"
// @Param version formData string true "版本"
// @Param serviceDesc formData string true "服务概述"
// @Param upstreamUrl formData string true "服务跳转地址"
// @router /update [post]
func (this *ServiceController) UpdateService() {
	serviceId := this.GetString("serviceId")
	if len(serviceId) == 0 {
		logs.Error("服务ID不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务ID不能为空")
	}
	serviceName := this.GetString("serviceName")
	if len(serviceName) == 0 {
		logs.Error("服务名称不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务名称不能为空")
	}
	servicePic := this.GetString("servicePic")
	if len(servicePic) == 0 {
		logs.Error("服务logo图片不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务logo图片不能为空")
	}
	requestPath := this.GetString("requestPath")
	if len(requestPath) == 0 {
		logs.Error("服务标识不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务标识不能为空")
	}
	provider := this.GetString("provider")
	if len(provider) == 0 {
		logs.Error("服务提供者不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务提供者不能为空")
	}
	version := this.GetString("version")
	if len(version) == 0 {
		logs.Error("版本不能为空")
		this.CustomAbort(http.StatusBadRequest, "版本不能为空")
	}
	serviceDesc := this.GetString("serviceDesc")
	if len(serviceDesc) == 0 {
		logs.Error("服务概述不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务概述不能为空")
	}
	upstreamUrl := this.GetString("upstreamUrl")
	if len(upstreamUrl) == 0 {
		logs.Error("服务访问路径不能为空")
		this.CustomAbort(http.StatusBadRequest, "服务访问路径不能为空")
	}
	var service models.ClService
	currService, err := models.GetService(serviceId)
	if err != nil {
		logs.Error("获取服务失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取服务失败")
	}
	if currService.RequestPath != `/` + requestPath || currService.UpstreamUrl != upstreamUrl {
		//todo 调用KONG API对相应的apiId进行更新
		var updApi km.API
		updApi.RequestPath = `/` + requestPath
		updApi.UpstreamURL = upstreamUrl
		updApi.ID = currService.ApiId
		_, errUpd := kong.UpdateAPI(currService.ApiId, &updApi)
		if errUpd != nil {
			logs.Error("更新API失败%v", errUpd)
			this.CustomAbort(http.StatusInternalServerError, "更新API失败")
		}
	}
	//更新数据库
	service.Id = serviceId
	service.ServiceName = serviceName
	service.ServicePic = servicePic
	service.RequestPath = `/` + requestPath
	service.Provider = provider
	service.Version = version
	service.ServiceDesc = serviceDesc
	service.UpstreamUrl = upstreamUrl
	errUpd := models.UpdService(&service)
	if errUpd != nil {
		logs.Error("更新服务失败%v", errUpd)
		this.CustomAbort(http.StatusInternalServerError, "更新服务失败")
	}
	this.Data["json"] = map[string]string{"msg":"更新服务成功"}
	this.ServeJSON()
}
// @Description 删除服务
// @Param serviceId path string true "服务uuid"
// @router /del/:serviceId [post]
func (this *ServiceController) DeleteService() {
	serviceId := this.GetString(":serviceId")
	//调用KONG api删除对应的service
	_service, err := models.GetService(serviceId)
	if err != nil {
		logs.Error("获取服务失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取服务失败")
	}
	errDel := kong.DeleteAPI(_service.ApiId)
	if errDel != nil {
		logs.Error("删除服务失败%v", errDel)
		this.CustomAbort(http.StatusInternalServerError, "删除服务失败")
	}
	//删除数据库中的API
	errSer := models.DelService(serviceId)
	if errSer != nil {
		logs.Error("删除服务失败%v", errSer)
		this.CustomAbort(http.StatusInternalServerError, "删除服务失败")
	}
	this.Data["json"] = "删除服务成功"
	this.ServeJSON()
}