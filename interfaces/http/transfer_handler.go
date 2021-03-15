package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/moneytransfer/application"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/logger"
	"net/http"
)

type TransferHandlerInterface interface {
	GetTransfers(c *gin.Context)
	CreateTransfer(c *gin.Context)
}

type transferHandler struct {
	ca application.TransferAppInterface
}

func NewTransfers(ca application.TransferAppInterface ) TransferHandlerInterface {
	return &transferHandler{
		ca: ca,
	}
}

func (ch *transferHandler) GetTransfers(c *gin.Context) {

	transferID:= c.Param("transfer_id")
	if transferID=="" {
		c.JSON(http.StatusBadRequest,errors.NewBadRequestError("Invalid Id"))
		return
	}
	transfers := &entity.Transfer{}
	var err *errors.RestErr
	transfers, err = ch.ca.GetTransfer(transferID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transfers)
}


func (ch *transferHandler) CreateTransfer(c *gin.Context) 	{

	var transfer entity.Transfer
	if err:=c.ShouldBindJSON(&transfer);err!=nil{
		restErr:=errors.NewBadRequestError("Invalid Form Body")
		logger.Error("error Binding save request",err)
		c.JSON(restErr.Status,restErr)
		return
	}

	savedTransfer := &entity.Transfer{}
	var err *errors.RestErr
	savedTransfer, err = ch.ca.CreateTransfer(transfer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, savedTransfer)
}

