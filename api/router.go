package api

import (
	"rent-car/api/handler"
	"rent-car/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetAllCars)
	r.GET("/car", h.GetAllCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	
	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetAllCustomer)
	r.GET("/customer", h.GetAllCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetAllOrder)
	r.GET("/order", h.GetAllOrder)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)

	return r
}
