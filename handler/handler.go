package handler

import (
	"github.com/cdempsie/payments-example/persist"
)

// PaymentHandler holds a persistent store that can be used to store payments.
// Calls are simply delegated to the underlying store implementation.
// The Handler exists to allow plugability of different stores.
type PaymentHandler struct {
	persist.PaymentStore
}

// NewPaymentHandler returns a new handler configured to use the given PaymentStore.
func NewPaymentHandler(store persist.PaymentStore) *PaymentHandler {
	return &PaymentHandler{store}
}
