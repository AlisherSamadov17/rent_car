package handler

import (
	"fmt"
	"net/http"
	"rent-car/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CreateOrder(c *gin.Context)  {
	order:=models.CreateOrder{}

	if err := c.ShouldBindJSON(&order);err != nil {
		handleResponse(c ,"error while reding request",http.StatusBadRequest,err.Error())
		return
	}
	
	id,err :=h.Store.Order().Create(order)
	if err != nil{
		handleResponse(c,"error while creating order",http.StatusInternalServerError,err.Error())
		return
	}
	handleResponse(c,"ok",http.StatusOK,id)
}

func (h Handler) UpdateOrder(c *gin.Context)  {
	order:=models.UpdateOrder{}
	
	if err := c.ShouldBindJSON(&order);err != nil {
		handleResponse(c,"error while reading request body",http.StatusBadRequest,err.Error())
		return
	}

	order.Id = c.Param("id")
	err := uuid.Validate(order.Id)
	if err != nil {
		handleResponse(c,"error while validating",http.StatusBadRequest,err.Error())
	return
	}
	
id,err :=h.Store.Order().Update(order)
if err != nil {
	handleResponse(c,"error while updating customer,err",http.StatusInternalServerError,err.Error())
	return
}
handleResponse(c ,"ok",http.StatusOK,id)
}

func (h Handler) GetAllOrder(c *gin.Context)  {
	var (
		
		request = models.GetAllOrdersRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c,"error while parsing limit",http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	orders, err := h.Store.Order().GetAll(request)
	if err != nil {
		handleResponse(c,"error while getting orders",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,"ok", http.StatusOK, orders)
}

func (h Handler) GetByIDOrder(c *gin.Context)  {
      id := c.Param("id")

	order,err := h.Store.Car().GetByID(id)
	if err != nil {
		handleResponse(c,"error while getting order by id",http.StatusInternalServerError,err.Error())
		return
	}
	handleResponse(c,"ok",http.StatusOK,order)
}


func (h Handler) DeleteOrder(c *gin.Context)  {
	id :=c.Param("id")
	err := uuid.Validate(id)
	if err != nil {
	handleResponse(c,"error while validating id",http.StatusBadRequest,err.Error())
	return
}
err = h.Store.Customer().Delete(id)
if err != nil {
	handleResponse(c,"error while deleting customer",http.StatusInternalServerError,err.Error())
return
}
handleResponse(c,"ok",http.StatusOK,id)
}