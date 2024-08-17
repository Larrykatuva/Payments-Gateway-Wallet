package actors

import (
	"example.com/m/gateway/dto"
	"example.com/m/gateway/services"
	"fmt"
	"github.com/anthdm/hollywood/actor"
)

type UtilsManager struct{}

func NewUtilsManager() actor.Receiver {
	return &UtilsManager{}
}

func (u *UtilsManager) Receive(context *actor.Context) {
	switch msg := context.Message().(type) {
	case dto.RequestRrn:
		fmt.Print(msg)
		rrn := services.GenerateRrn()
		context.Respond(dto.RrnResponse{Rrn: rrn})
	}
}
