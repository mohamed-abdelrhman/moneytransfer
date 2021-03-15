package application

import (
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/service"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
)

type transferApp struct {
	cs service.TransferServiceInterface
}


var _ TransferAppInterface = &transferApp{}

type TransferAppInterface interface {
	GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr)
	CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr)
}
func NewTransferApp(cs service.TransferServiceInterface ) TransferAppInterface {
	return &transferApp{
		cs: cs,
	}
}
func (c *transferApp) GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr) {
	return c.cs.GetTransfer(transferID)
}
func (c *transferApp) CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr) {
	return c.cs.CreateTransfer(transfer)
}