package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RecordController struct {
	beego.Controller
}

var log = logs.NewLogger()
var IsInit = false

func Init() {
	if IsInit {
		return
	}
	IsInit = true
	log.Async(4080)
	log.SetLogger("file", `{"filename":"logs/AppListInfo.log"}`)
}

func (this *RecordController) Get() {
	Init()
	log.Info(this.GetString("data"))
	this.TplName = "record.tpl"
}

func (this *RecordController) Post() {
	Init()
	log.Info(string(this.Ctx.Input.RequestBody))
	this.TplName = "record.tpl"
}
