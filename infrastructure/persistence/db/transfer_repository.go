package db

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/clients"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
	"log"
)


type TransferRepositoryInterface interface {
	GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr)
	CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr)
}

type transferRepository struct {
}

func NewTransferRepository() TransferRepositoryInterface {
	return &transferRepository{}
}

func (r *transferRepository)GetTransfer(transferID string) (*entity.Transfer, *errors.RestErr){
	client :=clients.GetRedisClient()
	transfer :=entity.Transfer{}
	ctx := context.TODO()
	data := client.Get(ctx, transferID)
	err := json.Unmarshal([]byte(data.Val()), &transfer)
	if err != nil {
		return nil, errors.NewNotFoundError("No Such Transfer")
	}
	return &transfer,nil
}

func (r *transferRepository)CreateTransfer(transfer entity.Transfer) (*entity.Transfer, *errors.RestErr){
	transfer.ID=uuid.New().String()
	client :=clients.GetRedisClient()
	ctx := context.TODO()
	client.Del(ctx,transfer.ID)
	v,err :=json.Marshal(transfer)
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("Some thing went wrong")
	}
	client.Set(ctx, transfer.ID, v, 0)
	return &transfer,nil
}