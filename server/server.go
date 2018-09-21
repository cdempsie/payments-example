package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cdempsie/payments-example/api"
	payment_handler "github.com/cdempsie/payments-example/handler"
	"github.com/cdempsie/payments-example/persist"
	"github.com/gorilla/mux"
)

var (
	handler *payment_handler.PaymentHandler
	port    int
	store   string
)

func init() {
	flag.StringVar(&store, "store", "in-memory", "The persitance store to use, the default and only option at the moment is in-memory")
	flag.IntVar(&port, "port", 8000, "The port number to start the server on, defaults to 8000")
}

func main() {
	router := mux.NewRouter()
	paymentSubRoute := router.PathPrefix("/v1/payment").Subrouter()
	// CRUD for payment
	paymentSubRoute.HandleFunc("", createPaymentHandler).Methods(http.MethodPost)
	paymentSubRoute.HandleFunc("", updatePaymentHandler).Methods(http.MethodPut)
	paymentSubRoute.HandleFunc("/{payment-id}", getPaymentHandler).Methods(http.MethodGet)
	paymentSubRoute.HandleFunc("/{payment-id}", deletePaymentHandler).Methods(http.MethodDelete)

	// Collection of payments
	router.HandleFunc("/v1/payments", listPaymentsHandler).Methods(http.MethodGet)

	if err := configure(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// start the server, defaults to :8000
	portStr := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(portStr, router))
}

// configure will parse the command line flags and setup the handler with the requested persistent store type.
func configure() error {
	if err := parseFlags(); err != nil {
		return err
	}

	// more store types could be added here for example DB, file, etc
	if store == "in-memory" {
		handler = payment_handler.NewPaymentHandler(persist.NewInMemoryStore())
	} else {
		return fmt.Errorf("unknown store type requested: %s", store)
	}
	return nil
}

// parseFlags parses the command line flags returning any errors.
func parseFlags() error {
	flag.Parse()
	if store != "in-memory" {
		return fmt.Errorf("invalid store value: %s only \"in-memory\" is currently supported", store)
	}

	return nil
}

// createPaymentHandler creates a new payment with the given details.
// If the request is badly formed a 400 bad request is returned.
func createPaymentHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, "Badly formed request: empty body")
		return
	}

	dec := json.NewDecoder(request.Body)
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseWriter, "Badly formed request: %v", err)
		return
	}

	if ok, msg := payment.Valid(); !ok {
		http.Error(responseWriter, msg, http.StatusBadRequest)
		return
	}

	err = handler.Create(payment)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to create payment: %v", err)
		return
	}

	writeResult(responseWriter, payment)
}

// updatePaymentHandler creates a new payment with the given details.
// If the request is badly formed a 400 bad request is returned.
func updatePaymentHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, "Badly formed request: empty body")
		return
	}

	dec := json.NewDecoder(request.Body)
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseWriter, "Badly formed request: %v", err)
		return
	}

	if ok, msg := payment.Valid(); !ok {
		http.Error(responseWriter, msg, http.StatusBadRequest)
		return
	}

	err = handler.Update(payment)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to update payment: %v", err)
		return
	}

	writeResult(responseWriter, payment)
}

// getPaymentHandler fetches the payment with the given ID. If the ID is missing a 400 bad request is returned.
func getPaymentHandler(responseWriter http.ResponseWriter, request *http.Request) {
	paymentID, ok := validPaymentID(responseWriter, request)
	if !ok {
		return
	}

	log.Printf("Got payment ID: %s", paymentID)

	payment, err := handler.Load(paymentID)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to get payment: %v", err)
		return
	}

	writeResult(responseWriter, payment)
}

// deletePaymentHandler deleted the payment with the given ID. If the ID is missing a 400 bad request is returned.
func deletePaymentHandler(responseWriter http.ResponseWriter, request *http.Request) {
	paymentID, ok := validPaymentID(responseWriter, request)
	if !ok {
		return
	}

	log.Printf("Got payment ID: %s", paymentID)

	err := handler.Delete(paymentID)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to delete payment: %v", err)
		return
	}
}

// validPaymentID checks for the presence of the payment ID in the path.
// If the ID is found, true is returned along with the payment ID.
// If the ID is not found, false is returned and a 400 bad request is sent to the caller.
func validPaymentID(responseWriter http.ResponseWriter, request *http.Request) (paymentID string, isValid bool) {
	vars := mux.Vars(request)
	paymentID = vars["payment-id"]
	if paymentID == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return "", false
	}

	return paymentID, true
}

// listPaymentsHandler returns a list of payments.
func listPaymentsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	payments, err := handler.List()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to list payments: %v", err)
		return
	}

	writeResult(responseWriter, payments)
}

// writeResult writes the value as JSON to the response. If the encoding fails 500 is returned with a message.
func writeResult(responseWriter http.ResponseWriter, val interface{}) {
	enc := json.NewEncoder(responseWriter)
	err := enc.Encode(val)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "failed to encode response: %v", err)
	}

	responseWriter.Header().Add("Content-Type", "application/json")
}
