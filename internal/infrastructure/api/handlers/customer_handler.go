package handlers

import (
	"neoway_test/internal/domain/customer/dto"
	usecaseCreate "neoway_test/internal/usecase/customer/create"
	usecaseDelete "neoway_test/internal/usecase/customer/delete"
	usecaseFind "neoway_test/internal/usecase/customer/find"
	usecaseList "neoway_test/internal/usecase/customer/list"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CustomerHandler handles HTTP requests for customer operations.
type CustomerHandler struct {
	getCustomersListUsecase    *usecaseList.GetCustomersListUseCase
	createCustomerUsecase      *usecaseCreate.CreateCustomerUseCase
	createCustomersBulkUsecase *usecaseCreate.CreateCustomerBulkUseCase
	getCustomerByCpfUsecase    *usecaseFind.GetCustomerByCpfUseCase
	getCustomerByIdUsecase     *usecaseFind.GetCustomerByIdUseCase
	deleteCustomersUsecase     *usecaseDelete.DeleteCustomerUseCase
}

// NewCustomerHandler creates a new CustomerHandler.
func NewCustomerHandler(
	getCustomersListUsecase *usecaseList.GetCustomersListUseCase,
	createCustomerUsecase *usecaseCreate.CreateCustomerUseCase,
	createCustomersBulkUsecase *usecaseCreate.CreateCustomerBulkUseCase,
	getCustomerByCpfUsecase *usecaseFind.GetCustomerByCpfUseCase,
	getCustomerByIdUsecase *usecaseFind.GetCustomerByIdUseCase,
	deleteCustomersUsecase *usecaseDelete.DeleteCustomerUseCase,
) *CustomerHandler {
	return &CustomerHandler{
		getCustomersListUsecase:    getCustomersListUsecase,
		createCustomerUsecase:      createCustomerUsecase,
		createCustomersBulkUsecase: createCustomersBulkUsecase,
		getCustomerByCpfUsecase:    getCustomerByCpfUsecase,
		getCustomerByIdUsecase:     getCustomerByIdUsecase,
		deleteCustomersUsecase:     deleteCustomersUsecase,
	}
}

// CustomerPost handles the request to create a new customer.
// @Summary Create a new customer
// @Description Create a new customer with the provided details
// @Tags Customers
// @Accept json
// @Produce json
// @Param input body dto.InputCreateCustomerDto true "Customer data"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer [post]
func (h *CustomerHandler) CustomerPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request dto.InputCreateCustomerDto

	if err := render.DecodeJSON(r.Body, &request); err != nil {
		return nil, http.StatusBadRequest, err // handle JSON decode error
	}

	output, err := h.createCustomerUsecase.Execute(request)

	if err != nil {
		return map[string]string{"id": output.ID}, http.StatusInternalServerError, err
	}

	return map[string]string{"id": output.ID}, http.StatusCreated, err
}

// CustomerPostBulk handles the request to create customers in bulk.
// @Summary Create multiple customers in bulk
// @Description Create multiple customers from a provided file
// @Tags Customers
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file with customer data"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer/bulkCreation [post]
func (h *CustomerHandler) CustomerPostBulk(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	file, _, err := r.FormFile("file")

	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	defer file.Close()

	message, err := h.createCustomersBulkUsecase.Execute(file)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return map[string]string{"message": message}, http.StatusCreated, err
}

// CustomerGet handles the request to list customers.
// @Summary List all customers
// @Description Get a paginated list of customers
// @Tags Customers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Success 200 {array} dto.OutputGetCustomersListDto
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer [get]
func (h *CustomerHandler) CustomerGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	input := dto.InputGetCustomersListDto{Page: page}

	customers, err := h.getCustomersListUsecase.Execute(input)

	if err == nil && customers == nil {
		return nil, http.StatusNotFound, err
	}
	return customers, http.StatusOK, err
}

// CustomerGetById handles the request to get a customer by ID.
// @Summary Get customer details by ID
// @Description Get details of a customer by ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} dto.OutputGetCustomerDto
// @Failure 404 {object} string "Customer not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer/getById/{id} [get]
func (h *CustomerHandler) CustomerGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	input := dto.InputGetCustomerByIdDto{ID: id}

	customer, err := h.getCustomerByIdUsecase.Execute(input)
	if err == nil && customer == nil {
		return nil, http.StatusNotFound, err
	}
	return customer, http.StatusOK, err
}

// CustomerGetByCpf handles the request to get a customer by CPF.
// @Summary Get customer details by CPF
// @Description Get details of a customer by CPF
// @Tags Customers
// @Accept json
// @Produce json
// @Param cpf path string true "Customer CPF"
// @Success 200 {object} dto.OutputGetCustomerDto
// @Failure 404 {object} string "Customer not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer/getByCpf/{cpf} [get]
func (h *CustomerHandler) CustomerGetByCpf(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	cpf := chi.URLParam(r, "cpf")

	input := dto.InputGetCustomerByCpfDto{Cpf: cpf}

	customer, err := h.getCustomerByCpfUsecase.Execute(input)
	if err == nil && customer == nil {
		return nil, http.StatusNotFound, err
	}
	return customer, http.StatusOK, err
}

// CustomerDelete handles the request to delete a customer by ID.
// @Summary Delete a customer
// @Description Delete a customer by ID
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} string "Customer successfully deleted"
// @Failure 404 {object} string "Customer not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/customer/{id} [delete]
func (h *CustomerHandler) CustomerDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	input := dto.InputDeleteCustomerDto{ID: id}

	err := h.deleteCustomersUsecase.Execute(input)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return nil, http.StatusOK, err
}
