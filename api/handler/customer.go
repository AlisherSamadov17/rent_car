package handler

import (

	"fmt"
	"net/http"
	"rent-car/api/models"
	"rent-car/pkg/check"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)





func (h Handler) CreateCustomer(c *gin.Context)  {
	customer:=models.Customer{}

	if err := c.ShouldBindJSON(&customer);err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
         return
	}

	if err := check.ValidateGmailCustomer(customer.Gmail);!true {
	   handleResponse(c,"error while validating Email"+customer.Gmail,http.StatusBadRequest,err)
	   return
	}

	if err := check.ValidatePhoneNumberOfCustomer(customer.Phone);!true {
	   handleResponse(c,"error while validating PhoneNumber"+customer.Phone,http.StatusBadRequest,err)
	   return
	}

	id,err :=h.Store.Customer().Create(customer)
	if err != nil{
		handleResponse(c,"error while creating customer",http.StatusInternalServerError,err.Error())
		return
	}
	handleResponse(c,"",http.StatusOK,id)
}
func (h Handler) UpdateCustomer(c *gin.Context) {

	customer:=models.Customer{}

	if err := c.ShouldBindJSON(&customer);err != nil {
		handleResponse(c,"error while reading request body",http.StatusBadRequest,err.Error())
		return
	}
	
	if err := check.ValidateGmailCustomer(customer.Gmail); !true {
		handleResponse(c,"error while validating Email"+customer.Gmail,http.StatusBadRequest,err)
		return
	 }
 
	 if err := check.ValidatePhoneNumberOfCustomer(customer.Phone);!true {
		handleResponse(c,"error while validating PhoneNumber"+customer.Phone,http.StatusBadRequest,err)
		return
	 }
	customer.Id = c.Query("id")

	err := uuid.Validate(customer.Id)
	if err != nil {
		handleResponse(c,"error while validating",http.StatusBadRequest,err.Error())
	return
	}

 id,err :=h.Store.Customer().UpdateCustomer(customer)
 
if err != nil {
	handleResponse(c,"error while updating customer",http.StatusInternalServerError,err)
	return
}
handleResponse(c,"ok",http.StatusOK,id)
}
func (h Handler) GetAllCustomer(c *gin.Context)  {
	var (
		request = models.GetAllCustomersRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c,"error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	customers, err := h.Store.Customer().GetAllCustomer(request)
	if err != nil {
		handleResponse(c,"error while getting customers",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,"ok",http.StatusOK, customers)
}
func (h Handler) GetByIDCustomer(c *gin.Context)  {
      
    id := c.Param("id")
	fmt.Println("id:",id)

	customer,err := h.Store.Customer().GetByID(id)
	if err != nil {
		handleResponse(c,"error while getting customer by id",http.StatusInternalServerError,err.Error())
		return
	}
	handleResponse(c,"ok",http.StatusOK,customer)
}


func (h Handler) GetAllCustomerCars(c *gin.Context) {
	var (
		request = models.GetAllCustomerCarsRequest{}
	)

	request.Search = c.Query("search")
	request.Id=c.Param("id")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Orders, err := h.Store.Customer().GetAllCustomerCars(request)
	if err != nil {
		handleResponse(c, "error while gettign CustomerCars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, Orders)
}

func (h Handler) DeleteCustomer(c *gin.Context)  {
          
    id := c.Param("id")
	fmt.Println("id:",id)

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