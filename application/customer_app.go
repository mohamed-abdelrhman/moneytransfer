package application

import (
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/service"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
)

type customerApp struct {
	cs service.CustomerServiceInterface
}


var _ CustomerAppInterface = &customerApp{}

type CustomerAppInterface interface {
	GetCustomer(customerID string) (*entity.Customer, *errors.RestErr)
	CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr)
}
func NewCustomerApp(cs service.CustomerServiceInterface ) CustomerAppInterface {
	return &customerApp{
		cs: cs,
	}
}
func (c *customerApp) GetCustomer(customerID string) (*entity.Customer, *errors.RestErr) {
	return c.cs.GetCustomer(customerID)
}
func (c *customerApp) CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr) {
	return c.cs.CreateCustomer(customer)
}