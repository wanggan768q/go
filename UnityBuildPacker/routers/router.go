package routers

import (
	"UnityBuildPacker/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/Connect", &controllers.WebSocketController{}, "get:Connect")

	beego.Router("/Upload", &controllers.UploadController{})
	beego.Router("/Upload/File", &controllers.UploadController{}, "post:UploadFile")

	beego.Router("/record/get", &controllers.RecordController{}, "*:Get")
	beego.Router("/record/post", &controllers.RecordController{}, "*:Post")
	beego.AutoRouter(&controllers.RecordController{})
}
