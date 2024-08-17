package actors

import "github.com/anthdm/hollywood/actor"

type TransactionManager struct{}

func NewTransactionManager() actor.Receiver {
	return &TransactionManager{}
}

func (t TransactionManager) Receive(context *actor.Context) {

}
