package db

import (
	"context"
	"encoding/json"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/entity"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/clients"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/utils/errors"
	"log"
)


type CustomerRepositoryInterface interface {
	GetCustomer(customerID string) (*entity.Customer, *errors.RestErr)
	AddBalance(customerID string, amount int) (*entity.Customer, *errors.RestErr)
	DeductBalance(customerID string, amount int) (*entity.Customer, *errors.RestErr)
	CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr)
}

type customerRepository struct {
}

func NewCustomerRepository() CustomerRepositoryInterface {
	return &customerRepository{}
}

func (r *customerRepository)GetCustomer(customerID string) (*entity.Customer, *errors.RestErr){
	client :=clients.GetRedisClient()
	customer :=entity.Customer{}
	ctx := context.TODO()
	data := client.Get(ctx, customerID)
	err := json.Unmarshal([]byte(data.Val()), &customer)
	if err != nil {
		return nil, errors.NewNotFoundError("No Such Customer")
	}
	return &customer,nil
}


func (r *customerRepository)CreateCustomer(customer entity.Customer) (*entity.Customer, *errors.RestErr){
	client :=clients.GetRedisClient()
	ctx := context.TODO()
	client.Del(ctx,customer.ID)

	v,err :=json.Marshal(customer)
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("Some thing went wrong")
	}
	client.Set(ctx, customer.ID, v, 0)

	return &customer,nil
}


func (r *customerRepository)AddBalance(customerID string, amount int) (*entity.Customer, *errors.RestErr){
	client :=clients.GetRedisClient()
	customer :=entity.Customer{}
	ctx := context.TODO()
	data := client.Get(ctx, customerID)
	err := json.Unmarshal([]byte(data.Val()), &customer)
	if err != nil {
		return nil, errors.NewNotFoundError("No Such Customer")
	}

	client.Del(ctx,customer.ID)
	customer.Balance=customer.Balance+ amount
	return r.CreateCustomer(customer)
}

func (r *customerRepository)DeductBalance(customerID string, amount int) (*entity.Customer, *errors.RestErr){
	client :=clients.GetRedisClient()
	customer :=entity.Customer{}
	ctx := context.TODO()
	data := client.Get(ctx, customerID)
	err := json.Unmarshal([]byte(data.Val()), &customer)
	if err != nil {
		return nil, errors.NewNotFoundError("No Such Customer")
	}

	client.Del(ctx,customer.ID)
	customer.Balance=customer.Balance- amount
	return r.CreateCustomer(customer)
}

