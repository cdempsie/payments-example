package persist

import "github.com/cdempsie/payments-example/api"

// PaymentStore defines the methods a persistent store must provide.
type PaymentStore interface {
	Create(payment *api.Payment) error
	Update(payment *api.Payment) error
	Delete(paymentUID string) error
	Load(paymentUID string) (payment *api.Payment, err error)
	List() (results *api.ListHolder, err error)
}
