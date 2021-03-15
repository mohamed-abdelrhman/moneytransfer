package api

import (
	"github.com/mohamed-abdelrhman/moneytransfer/application"
	"github.com/mohamed-abdelrhman/moneytransfer/domain/service"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/moneytransfer/interfaces/http"
)

func CustomerUrlMapping()  {
	customers := http.NewCustomers(application.NewCustomerApp(service.NewCustomerService(db.NewCustomerRepository())))
	Router.GET("/customers/:customer_id", customers.GetCustomers)
	Router.POST("/customers", customers.CreateCustomer)
}
