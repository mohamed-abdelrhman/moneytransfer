package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/moneytransfer/application"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/logger"
	"net/http"
)

type CustomerHandlerInterface interface {
	GetCustomers(c *gin.Context)
	CreateCustomer(c *gin.Context)
}

type customerHandler struct {
	ca application.CustomerAppInterface
}

func NewCustomers(ca application.CustomerAppInterface ) CustomerHandlerInterface {
	return &customerHandler{
		ca: ca,
	}
}

func (ch *customerHandler) GetCustomers(c *gin.Context) {

	customerID:= c.Param("customer_id")
	if customerID=="" {
		c.JSON(http.StatusBadRequest,errors.NewBadRequestError("Invalid Id"))
		return
	}
	customers := &entity.Customer{}
	var err *errors.RestErr
	customers, err = ch.ca.GetCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}


func (ch *customerHandler) CreateCustomer(c *gin.Context) 	{

	var customer entity.Customer
	if err:=c.ShouldBindJSON(&customer);err!=nil{
		restErr:=errors.NewBadRequestError("Invalid Form Body")
		logger.Error("error Binding save request",err)
		c.JSON(restErr.Status,restErr)
		return
	}

	savedCustomer := &entity.Customer{}
	var err *errors.RestErr
	savedCustomer, err = ch.ca.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, savedCustomer)
}

