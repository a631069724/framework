package controllers

import (
	"errors"

	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/monitor/app/logics/dashboard"
	"github.com/totoval/framework/request/websocket"
)

type DashboardWebsocketController struct {
	controller.BaseController
	websocket.BaseHandler
}

func (d *DashboardWebsocketController) OnMessage(hub websocket.Hub, msg *websocket.Msg) {
	mm := &websocket.Msg{}
	// need login~, just for an example of websocket authentication support
	if err := hub.ScanUser(); err != nil {
		mm.SetString(err.Error())
		hub.Send(mm)
		return
	}

	mm.SetJSON(hub.User().Value())
	hub.Send(mm)
	return
}

func (d *DashboardWebsocketController) Loop(hub websocket.Hub) error {
	select {
	case flow, ok := <-dashboard.Flow.Current():
		if !ok {
			return errors.New("connection closed")
		}
		msg := websocket.Msg{}
		msg.SetJSON(flow)
		hub.Broadcast(&msg)
	default:
		return nil
	}
	return nil
}
