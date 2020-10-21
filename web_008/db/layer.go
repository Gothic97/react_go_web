package db

import (
	"errors"
	"model"
)

type Layer interface {
	GetAllProducts() ([]model.Product, error)
	GetPromos() ([]model.Product, error)
	GetCustomerByName(string, string) (model.Customer, error)
	GetCustomerByID(int) (model.Customer, error)
	GetProduct(int) (model.Product, error)
	AddUser(model.Customer) (model.Customer, error)
	SignInUser(username, password string) (model.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]model.Order, error)
	AddOrder(model.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
