package service

import (
	"github.com/google/uuid"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
)

type CustomerServiceInterface interface {
	GetCustomer(customerID string) (*entity.Customer, *errors.RestErr)
	CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr)

}


type customerService struct {
	ur db.CustomerRepositoryInterface
}
func NewCustomerService(ur db.CustomerRepositoryInterface ) CustomerServiceInterface {
	return &customerService{
		ur: ur,
	}
}
func (s *customerService)GetCustomer(customerID string) (*entity.Customer, *errors.RestErr){
	return s.ur.GetCustomer(customerID)
}

func (s *customerService)CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr){
	customer.ID=uuid.New().String()
	return s.ur.CreateCustomer(customer)
}