package service

import (
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
)

type TransferServiceInterface interface {
	GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr)
	CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr)

}


type transferService struct {
	tr db.TransferRepositoryInterface
	cr db.CustomerRepositoryInterface
}
func NewTransferService(tr db.TransferRepositoryInterface, cr db.CustomerRepositoryInterface ) TransferServiceInterface {
	return &transferService{
		tr: tr,
		cr: cr,
	}
}
func (s *transferService)GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr){
	return s.tr.GetTransfer(transferID)
}

func (s *transferService)CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr){
	//check origin and destination existence
	var err *errors.RestErr
	originCustomer := &entity.Customer{
		ID: transfer.OriginID,
	}
	destinationCustomer := &entity.Customer{
		ID: transfer.DestinationID,
	}
	originCustomer,err=s.cr.GetCustomer(originCustomer.ID)
	if err != nil {
		return nil,err
	}
	destinationCustomer,err=s.cr.GetCustomer(destinationCustomer.ID)
	if err != nil {
		return nil,err
	}
	// check balance in origin
	if originCustomer.Balance<transfer.Amount {
		return nil,errors.NewBadRequestError("No Enough money with origin balanace")
	}
	// deduct balance from origin
	_,err=s.cr.DeductBalance(originCustomer.ID,transfer.Amount)
	if err != nil {
		return nil,err
	}
	// Add balance from origin
	_,err=s.cr.AddBalance(destinationCustomer.ID,transfer.Amount)
	if err != nil {
		return nil,err
	}
	return s.tr.CreateTransfer(transfer)

}