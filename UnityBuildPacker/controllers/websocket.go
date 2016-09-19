package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

func (this *WebSocketController) Get() {
	fmt.Println("Get")
	this.Data["IsWebSocket"] = true
	this.TplName = "websocket.html"
}

func (this *WebSocketController) Connect() {
	fmt.Println("Connect")
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "WS Connect fail -> "+err.Error(), 400)
		return
	} else if err != nil {
		beego.Error("WS Connect fail -> 400")
		return
	}
	fmt.Println("Connect ok")

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(p))
	}

}
