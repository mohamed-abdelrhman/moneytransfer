package api

import (
	"github.com/mohamed-abdelrhman/moneytransfer/application"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/service"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/moneytransfer/interfaces/http"
)

func TransferUrlMapping()  {
	transfers := http.NewTransfers(application.NewTransferApp(service.NewTransferService(db.NewTransferRepository(),db.NewCustomerRepository())))
	Router.GET("/transfers/:transfer_id", transfers.GetTransfers)
	Router.POST("/transfers", transfers.CreateTransfer)
}
