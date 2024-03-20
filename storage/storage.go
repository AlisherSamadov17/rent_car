package storage

import "rent-car/api/models"

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
	Create(models.Car) (string, error)
	GetByID(id string) (models.Car, error)
	GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(id string) error
}

type ICustomerStorage interface {
	Create(models.Customer) (string, error)
	GetByID(id string) (models.Customer, error)
	GetAllCustomerCars(request models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error)
	GetAllCustomer(req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	UpdateCustomer(models.Customer) (string, error)
	Delete(id string) error
}

type IOrderStorage interface {
	Create(models.CreateOrder) (string, error)
	GetByID(id string) (models.OrderAll, error)
	GetAll(request models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	Update(models.UpdateOrder) (string, error)
	Delete(id string) error
}

