package db

import (
	"github.com/jinzhu/gorm"
	"model"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() ([]model.Product, error) {
	return nil, nil
}

func (db *DBORM) GetPromos() ([]model.Product, error) {
	return nil, nil
}

func (db *DBORM) GetCustomerByName(string, string) (model.Customer, error) {
	return model.Customer{}, nil
}

func (db *DBORM) GetCustomerByID(int) (model.Customer, error) {
	return model.Customer{}, nil
}

func (db *DBORM) GetProduct(uint) (model.Product, error) {
	return model.Product{}, nil
}

func (db *DBORM) SignInUser(model.Customer) error {
	return nil
}

func (db *DBORM) SignOutUserById(int) error {
	return nil
}

func (db *DBORM) GetCustomerOrdersByID(int) ([]model.Order, error) {
	return nil, nil
}
