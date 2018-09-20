package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cdempsie/payments-example/api"
	payment_handler "github.com/cdempsie/payments-example/handler"
	"github.com/cdempsie/payments-example/persist/mocks"
	"github.com/cdempsie/payments-example/test"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

const APIBase = "/v1/payment"

func TestCreateRequest(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Create", mock.Anything).Return(nil)
	handler = payment_handler.NewPaymentHandler(mockStore)
	req, err := http.NewRequest(http.MethodPost, APIBase, strings.NewReader(test.CreatePayment))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	createPaymentHandler(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateBadRequestNilBody(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Create", mock.Anything).Return(nil)
	req, err := http.NewRequest(http.MethodPost, APIBase, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	createPaymentHandler(recorder, req)

	//response := recorder.Result()

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestUpdateRequest(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Update", mock.Anything).Return(nil)
	handler = payment_handler.NewPaymentHandler(mockStore)
	req, err := http.NewRequest(http.MethodPut, APIBase, strings.NewReader(test.Payment))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	updatePaymentHandler(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateRequestFails(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Update", mock.Anything).Return(errors.New("failed to update"))
	handler = payment_handler.NewPaymentHandler(mockStore)
	req, err := http.NewRequest(http.MethodPut, APIBase, strings.NewReader(test.Payment))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	updatePaymentHandler(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestGetRequest(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(test.Payment))
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		t.Fatal(err)
	}

	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Load", mock.Anything).Return(payment, nil)
	handler = payment_handler.NewPaymentHandler(mockStore)

	path := fmt.Sprintf("%s/%s", APIBase, uuid.New().String())
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	pathPattern := fmt.Sprintf("%s/{payment-id}", APIBase)
	router.HandleFunc(pathPattern, getPaymentHandler)
	router.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetRequestFails(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Load", mock.Anything).Return(nil, errors.New("failed to load"))
	handler = payment_handler.NewPaymentHandler(mockStore)
	path := fmt.Sprintf("%s/%s", APIBase, uuid.New().String())
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	pathPattern := fmt.Sprintf("%s/{payment-id}", APIBase)
	router.HandleFunc(pathPattern, getPaymentHandler)
	router.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestDeleteRequest(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Delete", mock.Anything).Return(nil)
	handler = payment_handler.NewPaymentHandler(mockStore)

	path := fmt.Sprintf("%s/%s", APIBase, uuid.New().String())
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	pathPattern := fmt.Sprintf("%s/{payment-id}", APIBase)
	router.HandleFunc(pathPattern, deletePaymentHandler)
	router.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteRequestFails(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("Delete", mock.Anything).Return(errors.New("delete failed"))
	handler = payment_handler.NewPaymentHandler(mockStore)

	path := fmt.Sprintf("%s/%s", APIBase, uuid.New().String())
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	pathPattern := fmt.Sprintf("%s/{payment-id}", APIBase)
	router.HandleFunc(pathPattern, deletePaymentHandler)
	router.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestListRequest(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(test.Payment))
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		t.Fatal(err)
	}

	result := &api.ListHolder{}
	for i := 0; i < 3; i++ {
		result.Data = append(result.Data, *payment)
	}

	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("List", mock.Anything).Return(result, nil)
	handler = payment_handler.NewPaymentHandler(mockStore)

	req, err := http.NewRequest(http.MethodGet, "/v1/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	listPaymentsHandler(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestListRequestFails(t *testing.T) {
	// Pass a mock store to the handler
	mockStore := &mocks.PaymentStore{}
	mockStore.On("List", mock.Anything).Return(nil, errors.New("list failed"))
	handler = payment_handler.NewPaymentHandler(mockStore)

	req, err := http.NewRequest(http.MethodGet, "/v1/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	listPaymentsHandler(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
